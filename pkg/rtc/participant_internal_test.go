package rtc

import (
	"strings"
	"testing"
	"time"

	"github.com/pion/webrtc/v3"
	"github.com/stretchr/testify/require"

	"github.com/tomxiong/protocol/auth"
	"github.com/tomxiong/protocol/livekit"
	"github.com/tomxiong/protocol/utils"

	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/routing"
	"github.com/livekit/livekit-server/pkg/routing/routingfakes"
	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/livekit/livekit-server/pkg/rtc/types/typesfakes"
	"github.com/livekit/livekit-server/pkg/sfu/connectionquality"
)

func TestIsReady(t *testing.T) {
	tests := []struct {
		state livekit.ParticipantInfo_State
		ready bool
	}{
		{
			state: livekit.ParticipantInfo_JOINING,
			ready: false,
		},
		{
			state: livekit.ParticipantInfo_JOINED,
			ready: true,
		},
		{
			state: livekit.ParticipantInfo_ACTIVE,
			ready: true,
		},
		{
			state: livekit.ParticipantInfo_DISCONNECTED,
			ready: false,
		},
	}

	for _, test := range tests {
		t.Run(test.state.String(), func(t *testing.T) {
			p := &ParticipantImpl{}
			p.state.Store(test.state)
			require.Equal(t, test.ready, p.IsReady())
		})
	}
}

func TestTrackPublishing(t *testing.T) {
	t.Run("should send the correct events", func(t *testing.T) {
		p := newParticipantForTest("test")
		p.state.Store(livekit.ParticipantInfo_ACTIVE)
		track := &typesfakes.FakeMediaTrack{}
		track.IDReturns("id")
		published := false
		updated := false
		p.OnTrackUpdated(func(p types.LocalParticipant, track types.MediaTrack) {
			updated = true
		})
		p.OnTrackPublished(func(p types.LocalParticipant, track types.MediaTrack) {
			published = true
		})
		p.UpTrackManager.AddPublishedTrack(track)
		p.handleTrackPublished(track)

		require.True(t, published)
		require.False(t, updated)
		require.Len(t, p.UpTrackManager.publishedTracks, 1)

		track.AddOnCloseArgsForCall(0)()
		require.Len(t, p.UpTrackManager.publishedTracks, 0)
		require.True(t, updated)
	})

	t.Run("sends back trackPublished event", func(t *testing.T) {
		p := newParticipantForTest("test")
		sink := p.params.Sink.(*routingfakes.FakeMessageSink)
		p.AddTrack(&livekit.AddTrackRequest{
			Cid:    "cid",
			Name:   "webcam",
			Type:   livekit.TrackType_VIDEO,
			Width:  1024,
			Height: 768,
		})
		require.Equal(t, 1, sink.WriteMessageCallCount())
		res := sink.WriteMessageArgsForCall(0).(*livekit.SignalResponse)
		require.IsType(t, &livekit.SignalResponse_TrackPublished{}, res.Message)
		published := res.Message.(*livekit.SignalResponse_TrackPublished).TrackPublished
		require.Equal(t, "cid", published.Cid)
		require.Equal(t, "webcam", published.Track.Name)
		require.Equal(t, livekit.TrackType_VIDEO, published.Track.Type)
		require.Equal(t, uint32(1024), published.Track.Width)
		require.Equal(t, uint32(768), published.Track.Height)
	})

	t.Run("should not allow adding of duplicate tracks", func(t *testing.T) {
		p := newParticipantForTest("test")
		sink := p.params.Sink.(*routingfakes.FakeMessageSink)
		p.AddTrack(&livekit.AddTrackRequest{
			Cid:  "cid",
			Name: "webcam",
			Type: livekit.TrackType_VIDEO,
		})
		p.AddTrack(&livekit.AddTrackRequest{
			Cid:  "cid",
			Name: "duplicate",
			Type: livekit.TrackType_AUDIO,
		})

		require.Equal(t, 1, sink.WriteMessageCallCount())
	})

	t.Run("should not allow adding of duplicate tracks if already published by client id in signalling", func(t *testing.T) {
		p := newParticipantForTest("test")
		sink := p.params.Sink.(*routingfakes.FakeMessageSink)

		track := &typesfakes.FakeLocalMediaTrack{}
		track.SignalCidReturns("cid")
		// directly add to publishedTracks without lock - for testing purpose only
		p.UpTrackManager.publishedTracks["cid"] = track

		p.AddTrack(&livekit.AddTrackRequest{
			Cid:  "cid",
			Name: "webcam",
			Type: livekit.TrackType_VIDEO,
		})
		require.Equal(t, 0, sink.WriteMessageCallCount())
	})

	t.Run("should not allow adding of duplicate tracks if already published by client id in sdp", func(t *testing.T) {
		p := newParticipantForTest("test")
		sink := p.params.Sink.(*routingfakes.FakeMessageSink)

		track := &typesfakes.FakeLocalMediaTrack{}
		track.HasSdpCidCalls(func(s string) bool { return s == "cid" })
		// directly add to publishedTracks without lock - for testing purpose only
		p.UpTrackManager.publishedTracks["cid"] = track

		p.AddTrack(&livekit.AddTrackRequest{
			Cid:  "cid",
			Name: "webcam",
			Type: livekit.TrackType_VIDEO,
		})
		require.Equal(t, 0, sink.WriteMessageCallCount())
	})
}

func TestOutOfOrderUpdates(t *testing.T) {
	p := newParticipantForTest("test")
	p.SetMetadata("initial metadata")
	sink := p.GetResponseSink().(*routingfakes.FakeMessageSink)
	pi1 := p.ToProto()
	p.SetMetadata("second update")
	pi2 := p.ToProto()

	require.Greater(t, pi2.Version, pi1.Version)

	// send the second update first
	require.NoError(t, p.SendParticipantUpdate([]*livekit.ParticipantInfo{pi2}))
	require.NoError(t, p.SendParticipantUpdate([]*livekit.ParticipantInfo{pi1}))

	// only sent once, and it's the earlier message
	require.Equal(t, 1, sink.WriteMessageCallCount())
	sent := sink.WriteMessageArgsForCall(0).(*livekit.SignalResponse)
	require.Equal(t, "second update", sent.GetUpdate().Participants[0].Metadata)
}

// after disconnection, things should continue to function and not panic
func TestDisconnectTiming(t *testing.T) {
	t.Run("Negotiate doesn't panic after channel closed", func(t *testing.T) {
		p := newParticipantForTest("test")
		msg := routing.NewMessageChannel(routing.DefaultMessageChannelSize)
		p.params.Sink = msg
		go func() {
			for msg := range msg.ReadChan() {
				t.Log("received message from chan", msg)
			}
		}()
		track := &typesfakes.FakeMediaTrack{}
		p.UpTrackManager.AddPublishedTrack(track)
		p.handleTrackPublished(track)

		// close channel and then try to Negotiate
		msg.Close()
	})
}

func TestCorrectJoinedAt(t *testing.T) {
	p := newParticipantForTest("test")
	info := p.ToProto()
	require.NotZero(t, info.JoinedAt)
	require.True(t, time.Now().Unix()-info.JoinedAt <= 1)
}

func TestMuteSetting(t *testing.T) {
	t.Run("can set mute when track is pending", func(t *testing.T) {
		p := newParticipantForTest("test")
		ti := &livekit.TrackInfo{Sid: "testTrack"}
		p.pendingTracks["cid"] = &pendingTrackInfo{TrackInfo: ti}

		p.SetTrackMuted(livekit.TrackID(ti.Sid), true, false)
		require.True(t, ti.Muted)
	})

	t.Run("can publish a muted track", func(t *testing.T) {
		p := newParticipantForTest("test")
		p.AddTrack(&livekit.AddTrackRequest{
			Cid:   "cid",
			Type:  livekit.TrackType_AUDIO,
			Muted: true,
		})

		_, ti := p.getPendingTrack("cid", livekit.TrackType_AUDIO)
		require.NotNil(t, ti)
		require.True(t, ti.Muted)
	})
}

func TestConnectionQuality(t *testing.T) {
	testPublishedVideoTrack := func(params connectionquality.TrackScoreParams) *typesfakes.FakeLocalMediaTrack {
		tr := &typesfakes.FakeLocalMediaTrack{}
		score := connectionquality.VideoTrackScore(params)
		t.Log("video score: ", score)
		tr.GetConnectionScoreReturns(score)
		return tr
	}

	testPublishedAudioTrack := func(params connectionquality.TrackScoreParams) *typesfakes.FakeLocalMediaTrack {
		tr := &typesfakes.FakeLocalMediaTrack{}
		score := connectionquality.AudioTrackScore(params)
		t.Log("audio score: ", score)
		tr.GetConnectionScoreReturns(score)
		return tr
	}

	testPublishedScreenshareTrack := func(params connectionquality.TrackScoreParams) *typesfakes.FakeLocalMediaTrack {
		tr := &typesfakes.FakeLocalMediaTrack{}
		score := connectionquality.ScreenshareTrackScore(params)
		t.Log("screen share score: ", score)
		tr.GetConnectionScoreReturns(score)
		return tr
	}

	// TODO: this test is rather limited since we cannot mock DownTrack's Target & Max spatial layers
	// to improve this after split

	t.Run("smooth sailing", func(t *testing.T) {
		p := newParticipantForTest("test")

		// >2Mbps, 30fps,  expected/actual video size = 1280x720
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			PacketsExpected: 100,
			PacketsLost:     0,
			Bytes:           290000,
			Frames:          30,
			Jitter:          0.0,
			Rtt:             0,
			ExpectedWidth:   1280,
			ExpectedHeight:  720,
			ActualWidth:     1280,
			ActualHeight:    720,
		}
		p.UpTrackManager.publishedTracks["video"] = testPublishedVideoTrack(params)

		// no packet loss
		params = connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			Codec:           "opus",
			PacketsExpected: 100,
			PacketsLost:     0,
			Bytes:           1000,
			Jitter:          0.0,
			Rtt:             0,
			DtxDisabled:     false,
		}
		p.UpTrackManager.publishedTracks["audio"] = testPublishedAudioTrack(params)

		require.Equal(t, livekit.ConnectionQuality_EXCELLENT, p.GetConnectionQuality().GetQuality())
	})

	t.Run("reduced publishing", func(t *testing.T) {
		p := newParticipantForTest("test")

		// 1Mbps, 15fps,  expected = 1280x720, actual = 640 x 480
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			PacketsExpected: 100,
			PacketsLost:     0,
			Bytes:           125000,
			Frames:          15,
			Jitter:          0.0,
			Rtt:             0,
			ExpectedWidth:   1280,
			ExpectedHeight:  720,
			ActualWidth:     640,
			ActualHeight:    480,
		}
		p.UpTrackManager.publishedTracks["video"] = testPublishedVideoTrack(params)

		// packet loss of 5%
		params = connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			Codec:           "opus",
			PacketsExpected: 100,
			PacketsLost:     5,
			Bytes:           1000,
			Jitter:          0.0,
			Rtt:             0,
			DtxDisabled:     false,
		}
		p.UpTrackManager.publishedTracks["audio"] = testPublishedAudioTrack(params)

		require.Equal(t, livekit.ConnectionQuality_GOOD, p.GetConnectionQuality().GetQuality())
	})

	t.Run("audio smooth publishing", func(t *testing.T) {
		p := newParticipantForTest("test")
		// no packet loss
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			Codec:           "opus",
			PacketsExpected: 100,
			PacketsLost:     0,
			Bytes:           1000,
			Jitter:          0.0,
			Rtt:             0,
			DtxDisabled:     false,
		}
		p.UpTrackManager.publishedTracks["audio"] = testPublishedAudioTrack(params)

		require.Equal(t, livekit.ConnectionQuality_EXCELLENT, p.GetConnectionQuality().GetQuality())
	})

	t.Run("audio reduced publishing", func(t *testing.T) {
		p := newParticipantForTest("test")
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			Codec:           "opus",
			PacketsExpected: 100,
			PacketsLost:     5,
			Bytes:           1000,
			Jitter:          0.0,
			Rtt:             0,
			DtxDisabled:     false,
		}
		p.UpTrackManager.publishedTracks["audio"] = testPublishedAudioTrack(params)

		require.Equal(t, livekit.ConnectionQuality_GOOD, p.GetConnectionQuality().GetQuality())
	})

	t.Run("audio bad publishing", func(t *testing.T) {
		p := newParticipantForTest("test")
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			Codec:           "opus",
			PacketsExpected: 100,
			PacketsLost:     20,
			Bytes:           1000,
			Jitter:          0.0,
			Rtt:             0,
			DtxDisabled:     false,
		}
		p.UpTrackManager.publishedTracks["audio"] = testPublishedAudioTrack(params)

		require.Equal(t, livekit.ConnectionQuality_POOR, p.GetConnectionQuality().GetQuality())
	})

	t.Run("video smooth publishing", func(t *testing.T) {
		p := newParticipantForTest("test")

		// >2Mbps, 30fps,  expected/actual video size = 1280x720
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			PacketsExpected: 100,
			PacketsLost:     0,
			Bytes:           290000,
			Frames:          30,
			Jitter:          0.0,
			Rtt:             0,
			ExpectedWidth:   1280,
			ExpectedHeight:  720,
			ActualWidth:     1280,
			ActualHeight:    720,
		}
		p.UpTrackManager.publishedTracks["video"] = testPublishedVideoTrack(params)

		require.Equal(t, livekit.ConnectionQuality_EXCELLENT, p.GetConnectionQuality().GetQuality())
	})

	t.Run("video reduced publishing", func(t *testing.T) {
		p := newParticipantForTest("test")

		// 1Mbps, 15fps,  expected = 1280x720, actual = 640 x 480
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			PacketsExpected: 100,
			PacketsLost:     0,
			Bytes:           125000,
			Frames:          15,
			Jitter:          0.0,
			Rtt:             0,
			ExpectedWidth:   1280,
			ExpectedHeight:  720,
			ActualWidth:     640,
			ActualHeight:    480,
		}
		p.UpTrackManager.publishedTracks["video"] = testPublishedVideoTrack(params)

		require.Equal(t, livekit.ConnectionQuality_GOOD, p.GetConnectionQuality().GetQuality())
	})

	t.Run("video poor publishing", func(t *testing.T) {
		p := newParticipantForTest("test")

		// 20kbps, 8fps,  expected = 1280x720, actual = 240x426
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			PacketsExpected: 100,
			PacketsLost:     0,
			Bytes:           2500,
			Frames:          8,
			Jitter:          0.0,
			Rtt:             0,
			ExpectedWidth:   1280,
			ExpectedHeight:  720,
			ActualWidth:     240,
			ActualHeight:    426,
		}
		p.UpTrackManager.publishedTracks["video"] = testPublishedVideoTrack(params)

		require.Equal(t, livekit.ConnectionQuality_POOR, p.GetConnectionQuality().GetQuality())
	})

	t.Run("screen share no loss, not reduced quality, should be excellent", func(t *testing.T) {
		p := newParticipantForTest("test")

		// 20kbps, 2fps
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			PacketsExpected: 100,
			PacketsLost:     0,
			Bytes:           2500,
			Frames:          2,
		}
		p.UpTrackManager.publishedTracks["ss"] = testPublishedScreenshareTrack(params)

		require.Equal(t, livekit.ConnectionQuality_EXCELLENT, p.GetConnectionQuality().GetQuality())
	})

	t.Run("screen share low loss, not reduced quality, should be excellent", func(t *testing.T) {
		p := newParticipantForTest("test")

		// 20kbps, 2fps
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			PacketsExpected: 100,
			PacketsLost:     1,
			Bytes:           2500,
			Frames:          2,
		}
		p.UpTrackManager.publishedTracks["ss"] = testPublishedScreenshareTrack(params)

		require.Equal(t, livekit.ConnectionQuality_EXCELLENT, p.GetConnectionQuality().GetQuality())
	})

	t.Run("screen share high loss, not reduced quality, should be poor", func(t *testing.T) {
		p := newParticipantForTest("test")

		// 20kbps, 2fps
		params := connectionquality.TrackScoreParams{
			Duration:        1 * time.Second,
			PacketsExpected: 100,
			PacketsLost:     5,
			Bytes:           2500,
			Frames:          2,
		}
		p.UpTrackManager.publishedTracks["ss"] = testPublishedScreenshareTrack(params)

		require.Equal(t, livekit.ConnectionQuality_POOR, p.GetConnectionQuality().GetQuality())
	})

	t.Run("screen share no loss, but reduced quality, should be good", func(t *testing.T) {
		p := newParticipantForTest("test")

		// 20kbps, 2fps
		params := connectionquality.TrackScoreParams{
			Duration:         1 * time.Second,
			PacketsExpected:  100,
			PacketsLost:      0,
			Bytes:            2500,
			Frames:           2,
			IsReducedQuality: true,
		}
		p.UpTrackManager.publishedTracks["ss"] = testPublishedScreenshareTrack(params)

		require.Equal(t, livekit.ConnectionQuality_GOOD, p.GetConnectionQuality().GetQuality())
	})
}

func TestSubscriberAsPrimary(t *testing.T) {
	t.Run("protocol 4 uses subs as primary", func(t *testing.T) {
		p := newParticipantForTestWithOpts("test", &participantOpts{
			permissions: &livekit.ParticipantPermission{
				CanSubscribe: true,
				CanPublish:   true,
			},
		})
		require.True(t, p.SubscriberAsPrimary())
	})

	t.Run("protocol 2 uses pub as primary", func(t *testing.T) {
		p := newParticipantForTestWithOpts("test", &participantOpts{
			protocolVersion: 2,
			permissions: &livekit.ParticipantPermission{
				CanSubscribe: true,
				CanPublish:   true,
			},
		})
		require.False(t, p.SubscriberAsPrimary())
	})

	t.Run("publisher only uses pub as primary", func(t *testing.T) {
		p := newParticipantForTestWithOpts("test", &participantOpts{
			permissions: &livekit.ParticipantPermission{
				CanSubscribe: false,
				CanPublish:   true,
			},
		})
		require.False(t, p.SubscriberAsPrimary())

		// ensure that it doesn't change after perms
		p.SetPermission(&livekit.ParticipantPermission{
			CanSubscribe: true,
			CanPublish:   true,
		})
		require.False(t, p.SubscriberAsPrimary())
	})
}

func TestSetStableTrackID(t *testing.T) {
	testCases := []struct {
		name                 string
		trackInfo            *livekit.TrackInfo
		unpublished          []*livekit.TrackInfo
		prefix               string
		remainingUnpublished int
	}{
		{
			name: "first track, generates new ID",
			trackInfo: &livekit.TrackInfo{
				Type:   livekit.TrackType_VIDEO,
				Source: livekit.TrackSource_CAMERA,
			},
			prefix: "TR_VC",
		},
		{
			name: "re-using existing ID",
			trackInfo: &livekit.TrackInfo{
				Type:   livekit.TrackType_VIDEO,
				Source: livekit.TrackSource_CAMERA,
			},
			unpublished: []*livekit.TrackInfo{
				{
					Type:   livekit.TrackType_VIDEO,
					Source: livekit.TrackSource_SCREEN_SHARE,
					Sid:    "TR_VC1234",
				},
				{
					Type:   livekit.TrackType_VIDEO,
					Source: livekit.TrackSource_CAMERA,
					Sid:    "TR_VC1235",
				},
			},
			prefix:               "TR_VC1235",
			remainingUnpublished: 1,
		},
		{
			name: "mismatch name for reuse",
			trackInfo: &livekit.TrackInfo{
				Type:   livekit.TrackType_VIDEO,
				Source: livekit.TrackSource_CAMERA,
				Name:   "new_name",
			},
			unpublished: []*livekit.TrackInfo{
				{
					Type:   livekit.TrackType_VIDEO,
					Source: livekit.TrackSource_CAMERA,
					Sid:    "TR_NotUsed",
				},
			},
			prefix:               "TR_VC",
			remainingUnpublished: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := newParticipantForTest("test")
			p.unpublishedTracks = tc.unpublished

			ti := tc.trackInfo
			p.setStableTrackID(ti)
			require.Contains(t, ti.Sid, tc.prefix)
			require.Len(t, p.unpublishedTracks, tc.remainingUnpublished)
		})
	}
}

func TestDisableCodecs(t *testing.T) {
	participant := newParticipantForTestWithOpts(livekit.ParticipantIdentity("123"), &participantOpts{
		publisher: false,
		clientConf: &livekit.ClientConfiguration{
			DisabledCodecs: &livekit.DisabledCodecs{
				Codecs: []*livekit.Codec{
					{Mime: "video/h264"},
				},
			},
		},
	})

	participant.SetMigrateState(types.MigrateStateComplete)

	pc, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	require.NoError(t, err)
	transceiver, err := pc.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo, webrtc.RTPTransceiverInit{Direction: webrtc.RTPTransceiverDirectionSendrecv})
	require.NoError(t, err)
	sdp, err := pc.CreateOffer(nil)
	require.NoError(t, err)
	pc.SetLocalDescription(sdp)
	codecs := transceiver.Receiver().GetParameters().Codecs
	var found264 bool
	for _, c := range codecs {
		if strings.EqualFold(c.MimeType, "video/h264") {
			found264 = true
		}
	}
	require.True(t, found264)

	// negotiated codec should not contain h264
	anwser, err := participant.HandleOffer(sdp)
	require.NoError(t, err)
	require.NoError(t, pc.SetRemoteDescription(anwser), anwser.SDP, sdp.SDP)
	codecs = transceiver.Receiver().GetParameters().Codecs
	found264 = false
	for _, c := range codecs {
		if strings.EqualFold(c.MimeType, "video/h264") {
			found264 = true
		}
	}
	require.False(t, found264)
}

type participantOpts struct {
	permissions     *livekit.ParticipantPermission
	protocolVersion types.ProtocolVersion
	publisher       bool
	clientConf      *livekit.ClientConfiguration
}

func newParticipantForTestWithOpts(identity livekit.ParticipantIdentity, opts *participantOpts) *ParticipantImpl {
	if opts == nil {
		opts = &participantOpts{}
	}
	if opts.protocolVersion == 0 {
		opts.protocolVersion = 6
	}
	conf, _ := config.NewConfig("", nil)
	// disable mux, it doesn't play too well with unit test
	conf.RTC.UDPPort = 0
	conf.RTC.TCPPort = 0
	rtcConf, err := NewWebRTCConfig(conf, "")
	if err != nil {
		panic(err)
	}
	grants := &auth.ClaimGrants{
		Video: &auth.VideoGrant{},
	}
	if opts.permissions != nil {
		grants.Video.SetCanPublish(opts.permissions.CanPublish)
		grants.Video.SetCanPublishData(opts.permissions.CanPublishData)
		grants.Video.SetCanSubscribe(opts.permissions.CanSubscribe)
	}

	enabledCodecs := make([]*livekit.Codec, 0, len(conf.Room.EnabledCodecs))
	for _, c := range conf.Room.EnabledCodecs {
		enabledCodecs = append(enabledCodecs, &livekit.Codec{
			Mime:     c.Mime,
			FmtpLine: c.FmtpLine,
		})
	}
	p, _ := NewParticipant(ParticipantParams{
		SID:               livekit.ParticipantID(utils.NewGuid(utils.ParticipantPrefix)),
		Identity:          identity,
		Config:            rtcConf,
		Sink:              &routingfakes.FakeMessageSink{},
		ProtocolVersion:   opts.protocolVersion,
		PLIThrottleConfig: conf.RTC.PLIThrottle,
		Grants:            grants,
		EnabledCodecs:     enabledCodecs,
		ClientConf:        opts.clientConf,
	})
	p.isPublisher.Store(opts.publisher)

	return p
}

func newParticipantForTest(identity livekit.ParticipantIdentity) *ParticipantImpl {
	return newParticipantForTestWithOpts(identity, nil)
}
