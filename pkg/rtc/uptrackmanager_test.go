package rtc

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tomxiong/protocol/livekit"

	"github.com/tomxiong/livekit/pkg/rtc/types"
	"github.com/tomxiong/livekit/pkg/rtc/types/typesfakes"
)

func TestUpdateSubscriptionPermission(t *testing.T) {
	t.Run("updates subscription permission", func(t *testing.T) {
		um := NewUpTrackManager(UpTrackManagerParams{})

		tra := &typesfakes.FakeMediaTrack{}
		tra.IDReturns("audio")
		um.publishedTracks["audio"] = tra

		trv := &typesfakes.FakeMediaTrack{}
		trv.IDReturns("video")
		um.publishedTracks["video"] = trv

		// no restrictive subscription permission
		subscriptionPermission := &livekit.SubscriptionPermission{
			AllParticipants: true,
		}
		um.UpdateSubscriptionPermission(subscriptionPermission, nil, nil)
		require.Nil(t, um.subscriberPermissions)

		// nobody is allowed to subscribe
		subscriptionPermission = &livekit.SubscriptionPermission{
			TrackPermissions: []*livekit.TrackPermission{},
		}
		um.UpdateSubscriptionPermission(subscriptionPermission, nil, nil)
		require.NotNil(t, um.subscriberPermissions)
		require.Equal(t, 0, len(um.subscriberPermissions))

		lp1 := &typesfakes.FakeLocalParticipant{}
		lp1.IdentityReturns("p1")
		lp2 := &typesfakes.FakeLocalParticipant{}
		lp2.IdentityReturns("p2")

		sidResolver := func(sid livekit.ParticipantID) types.LocalParticipant {
			if sid == "p1" {
				return lp1
			}

			if sid == "p2" {
				return lp2
			}

			return nil
		}

		// allow all tracks for participants
		perms1 := &livekit.TrackPermission{
			ParticipantSid: "p1",
			AllTracks:      true,
		}
		perms2 := &livekit.TrackPermission{
			ParticipantSid: "p2",
			AllTracks:      true,
		}
		subscriptionPermission = &livekit.SubscriptionPermission{
			TrackPermissions: []*livekit.TrackPermission{
				perms1,
				perms2,
			},
		}
		um.UpdateSubscriptionPermission(subscriptionPermission, nil, sidResolver)
		require.Equal(t, 2, len(um.subscriberPermissions))
		require.EqualValues(t, perms1, um.subscriberPermissions["p1"])
		require.EqualValues(t, perms2, um.subscriberPermissions["p2"])

		// allow all tracks for some and restrictive for others
		perms1 = &livekit.TrackPermission{
			ParticipantIdentity: "p1",
			AllTracks:           true,
		}
		perms2 = &livekit.TrackPermission{
			ParticipantIdentity: "p2",
			TrackSids:           []string{"audio"},
		}
		perms3 := &livekit.TrackPermission{
			ParticipantIdentity: "p3",
			TrackSids:           []string{"video"},
		}
		subscriptionPermission = &livekit.SubscriptionPermission{
			TrackPermissions: []*livekit.TrackPermission{
				perms1,
				perms2,
				perms3,
			},
		}
		um.UpdateSubscriptionPermission(subscriptionPermission, nil, nil)
		require.Equal(t, 3, len(um.subscriberPermissions))
		require.EqualValues(t, perms1, um.subscriberPermissions["p1"])
		require.EqualValues(t, perms2, um.subscriberPermissions["p2"])
		require.EqualValues(t, perms3, um.subscriberPermissions["p3"])
	})

	t.Run("updates subscription permission using both", func(t *testing.T) {
		um := NewUpTrackManager(UpTrackManagerParams{})

		tra := &typesfakes.FakeMediaTrack{}
		tra.IDReturns("audio")
		um.publishedTracks["audio"] = tra

		trv := &typesfakes.FakeMediaTrack{}
		trv.IDReturns("video")
		um.publishedTracks["video"] = trv

		lp1 := &typesfakes.FakeLocalParticipant{}
		lp1.IdentityReturns("p1")
		lp2 := &typesfakes.FakeLocalParticipant{}
		lp2.IdentityReturns("p2")

		sidResolver := func(sid livekit.ParticipantID) types.LocalParticipant {
			if sid == "p1" {
				return lp1
			}

			if sid == "p2" {
				return lp2
			}

			return nil
		}

		// allow all tracks for participants
		perms1 := &livekit.TrackPermission{
			ParticipantSid:      "p1",
			ParticipantIdentity: "p1",
			AllTracks:           true,
		}
		perms2 := &livekit.TrackPermission{
			ParticipantSid:      "p2",
			ParticipantIdentity: "p2",
			AllTracks:           true,
		}
		subscriptionPermission := &livekit.SubscriptionPermission{
			TrackPermissions: []*livekit.TrackPermission{
				perms1,
				perms2,
			},
		}
		err := um.UpdateSubscriptionPermission(subscriptionPermission, nil, sidResolver)
		require.NoError(t, err)
		require.Equal(t, 2, len(um.subscriberPermissions))
		require.EqualValues(t, perms1, um.subscriberPermissions["p1"])
		require.EqualValues(t, perms2, um.subscriberPermissions["p2"])

		// mismatched identities should fail a permission update
		badSidResolver := func(sid livekit.ParticipantID) types.LocalParticipant {
			if sid == "p1" {
				return lp2
			}

			if sid == "p2" {
				return lp1
			}

			return nil
		}

		err = um.UpdateSubscriptionPermission(subscriptionPermission, nil, badSidResolver)
		require.NoError(t, err)
		require.Equal(t, 2, len(um.subscriberPermissions))
		require.EqualValues(t, perms1, um.subscriberPermissions["p1"])
		require.EqualValues(t, perms2, um.subscriberPermissions["p2"])
	})
}

func TestSubscriptionPermission(t *testing.T) {
	t.Run("checks subscription permission", func(t *testing.T) {
		um := NewUpTrackManager(UpTrackManagerParams{})

		tra := &typesfakes.FakeMediaTrack{}
		tra.IDReturns("audio")
		um.publishedTracks["audio"] = tra

		trv := &typesfakes.FakeMediaTrack{}
		trv.IDReturns("video")
		um.publishedTracks["video"] = trv

		// no restrictive permission
		subscriptionPermission := &livekit.SubscriptionPermission{
			AllParticipants: true,
		}
		um.UpdateSubscriptionPermission(subscriptionPermission, nil, nil)
		require.True(t, um.hasPermission("audio", "p1"))
		require.True(t, um.hasPermission("audio", "p2"))

		// nobody is allowed to subscribe
		subscriptionPermission = &livekit.SubscriptionPermission{
			TrackPermissions: []*livekit.TrackPermission{},
		}
		um.UpdateSubscriptionPermission(subscriptionPermission, nil, nil)
		require.False(t, um.hasPermission("audio", "p1"))
		require.False(t, um.hasPermission("audio", "p2"))

		// allow all tracks for participants
		subscriptionPermission = &livekit.SubscriptionPermission{
			TrackPermissions: []*livekit.TrackPermission{
				{
					ParticipantIdentity: "p1",
					AllTracks:           true,
				},
				{
					ParticipantIdentity: "p2",
					AllTracks:           true,
				},
			},
		}
		um.UpdateSubscriptionPermission(subscriptionPermission, nil, nil)
		require.True(t, um.hasPermission("audio", "p1"))
		require.True(t, um.hasPermission("video", "p1"))
		require.True(t, um.hasPermission("audio", "p2"))
		require.True(t, um.hasPermission("video", "p2"))

		// add a new track after permissions are set
		trs := &typesfakes.FakeMediaTrack{}
		trs.IDReturns("screen")
		um.publishedTracks["screen"] = trs

		require.True(t, um.hasPermission("audio", "p1"))
		require.True(t, um.hasPermission("video", "p1"))
		require.True(t, um.hasPermission("screen", "p1"))
		require.True(t, um.hasPermission("audio", "p2"))
		require.True(t, um.hasPermission("video", "p2"))
		require.True(t, um.hasPermission("screen", "p2"))

		// allow all tracks for some and restrictive for others
		subscriptionPermission = &livekit.SubscriptionPermission{
			TrackPermissions: []*livekit.TrackPermission{
				{
					ParticipantIdentity: "p1",
					AllTracks:           true,
				},
				{
					ParticipantIdentity: "p2",
					TrackSids:           []string{"audio"},
				},
				{
					ParticipantIdentity: "p3",
					TrackSids:           []string{"video"},
				},
			},
		}
		um.UpdateSubscriptionPermission(subscriptionPermission, nil, nil)
		require.True(t, um.hasPermission("audio", "p1"))
		require.True(t, um.hasPermission("video", "p1"))
		require.True(t, um.hasPermission("screen", "p1"))

		require.True(t, um.hasPermission("audio", "p2"))
		require.False(t, um.hasPermission("video", "p2"))
		require.False(t, um.hasPermission("screen", "p2"))

		require.False(t, um.hasPermission("audio", "p3"))
		require.True(t, um.hasPermission("video", "p3"))
		require.False(t, um.hasPermission("screen", "p3"))

		// add a new track after restrictive permissions are set
		trw := &typesfakes.FakeMediaTrack{}
		trw.IDReturns("watch")
		um.publishedTracks["watch"] = trw

		require.True(t, um.hasPermission("audio", "p1"))
		require.True(t, um.hasPermission("video", "p1"))
		require.True(t, um.hasPermission("screen", "p1"))
		require.True(t, um.hasPermission("watch", "p1"))

		require.True(t, um.hasPermission("audio", "p2"))
		require.False(t, um.hasPermission("video", "p2"))
		require.False(t, um.hasPermission("screen", "p2"))
		require.False(t, um.hasPermission("watch", "p2"))

		require.False(t, um.hasPermission("audio", "p3"))
		require.True(t, um.hasPermission("video", "p3"))
		require.False(t, um.hasPermission("screen", "p3"))
		require.False(t, um.hasPermission("watch", "p3"))
	})
}
