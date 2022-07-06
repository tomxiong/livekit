package types

import (
	"fmt"
	"time"

	"github.com/pion/webrtc/v3"

	"github.com/tomxiong/protocol/auth"
	"github.com/tomxiong/protocol/livekit"
	"github.com/tomxiong/protocol/logger"

	"github.com/tomxiong/livekit/pkg/routing"
	"github.com/tomxiong/livekit/pkg/sfu"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . WebsocketClient
type WebsocketClient interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	WriteControl(messageType int, data []byte, deadline time.Time) error
}

type AddSubscriberParams struct {
	AllTracks bool
	TrackIDs  []livekit.TrackID
}

// ---------------------------------------------

type MigrateState int32

const (
	MigrateStateInit MigrateState = iota
	MigrateStateSync
	MigrateStateComplete
)

func (m MigrateState) String() string {
	switch m {
	case MigrateStateInit:
		return "MIGRATE_STATE_INIT"
	case MigrateStateSync:
		return "MIGRATE_STATE_SYNC"
	case MigrateStateComplete:
		return "MIGRATE_STATE_COMPLETE"
	default:
		return fmt.Sprintf("%d", int(m))
	}
}

// ---------------------------------------------

type SubscribedCodecQuality struct {
	CodecMime string
	Quality   livekit.VideoQuality
}

// ---------------------------------------------

type ParticipantCloseReason int

const (
	ParticipantCloseReasonClientRequestLeave ParticipantCloseReason = iota
	ParticipantCloseReasonRoomManagerStop
	ParticipantCloseReasonRoomClose
	ParticipantCloseReasonVerifyFailed
	ParticipantCloseReasonJoinFailed
	ParticipantCloseReasonJoinTimeout
	ParticipantCloseReasonRTCSessionFinish
	ParticipantCloseReasonStateDisconnected
	ParticipantCloseReasonPeerConnectionDisconnected
	ParticipantCloseReasonDuplicateIdentity
	ParticipantCloseReasonMigrationComplete
	ParticipantCloseReasonStale
	ParticipantCloseReasonServiceRequestRemoveParticipant
	ParticipantCloseReasonServiceRequestDeleteRoom
	ParticipantCloseReasonSimulateMigration
	ParticipantCloseReasonSimulateNodeFailure
	ParticipantCloseReasonSimulateServerLeave
	ParticipantCloseReasonNegotiateFailed
)

func (p ParticipantCloseReason) String() string {
	switch p {
	case ParticipantCloseReasonClientRequestLeave:
		return "CLIENT_REQUEST_LEAVE"
	case ParticipantCloseReasonRoomManagerStop:
		return "ROOM_MANAGER_STOP"
	case ParticipantCloseReasonRoomClose:
		return "ROOM_CLOSE"
	case ParticipantCloseReasonVerifyFailed:
		return "VERIFY_FAILED"
	case ParticipantCloseReasonJoinFailed:
		return "JOIN_FAILED"
	case ParticipantCloseReasonJoinTimeout:
		return "JOIN_TIMEOUT"
	case ParticipantCloseReasonRTCSessionFinish:
		return "RTC_SESSION_FINISH"
	case ParticipantCloseReasonStateDisconnected:
		return "STATE_DISCONNECTED"
	case ParticipantCloseReasonPeerConnectionDisconnected:
		return "PEER_CONNECTION_DISCONNECTED"
	case ParticipantCloseReasonDuplicateIdentity:
		return "DUPLICATE_IDENTITY"
	case ParticipantCloseReasonMigrationComplete:
		return "MIGRATION_COMPLETE"
	case ParticipantCloseReasonStale:
		return "STALE"
	case ParticipantCloseReasonServiceRequestRemoveParticipant:
		return "SERVICE_REQUEST_REMOVE_PARTICIPANT"
	case ParticipantCloseReasonServiceRequestDeleteRoom:
		return "SERVICE_REQUEST_DELETE_ROOM"
	case ParticipantCloseReasonSimulateMigration:
		return "SIMULATE_MIGRATION"
	case ParticipantCloseReasonSimulateNodeFailure:
		return "SIMULATE_NODE_FAILURE"
	case ParticipantCloseReasonSimulateServerLeave:
		return "SIMULATE_SERVER_LEAVE"
	case ParticipantCloseReasonNegotiateFailed:
		return "NEGOTIATE_FAILED"
	default:
		return fmt.Sprintf("%d", int(p))
	}
}

func (p ParticipantCloseReason) ToDisconnectReason() livekit.DisconnectReason {
	switch p {
	case ParticipantCloseReasonClientRequestLeave:
		return livekit.DisconnectReason_CLIENT_INITIATED
	case ParticipantCloseReasonRoomManagerStop:
		return livekit.DisconnectReason_SERVER_SHUTDOWN
	case ParticipantCloseReasonVerifyFailed, ParticipantCloseReasonJoinFailed, ParticipantCloseReasonJoinTimeout:
		// expected to be connected but is not
		return livekit.DisconnectReason_JOIN_FAILURE
	case ParticipantCloseReasonPeerConnectionDisconnected:
		return livekit.DisconnectReason_STATE_MISMATCH
	case ParticipantCloseReasonDuplicateIdentity, ParticipantCloseReasonMigrationComplete, ParticipantCloseReasonStale:
		return livekit.DisconnectReason_DUPLICATE_IDENTITY
	case ParticipantCloseReasonServiceRequestRemoveParticipant:
		return livekit.DisconnectReason_PARTICIPANT_REMOVED
	case ParticipantCloseReasonServiceRequestDeleteRoom:
		return livekit.DisconnectReason_ROOM_DELETED
	case ParticipantCloseReasonSimulateMigration:
		return livekit.DisconnectReason_DUPLICATE_IDENTITY
	case ParticipantCloseReasonSimulateNodeFailure:
		return livekit.DisconnectReason_SERVER_SHUTDOWN
	case ParticipantCloseReasonSimulateServerLeave:
		return livekit.DisconnectReason_SERVER_SHUTDOWN
	case ParticipantCloseReasonNegotiateFailed:
		return livekit.DisconnectReason_STATE_MISMATCH
	default:
		// the other types will map to unknown reason
		return livekit.DisconnectReason_UNKNOWN_REASON
	}
}

// ---------------------------------------------

//counterfeiter:generate . Participant
type Participant interface {
	ID() livekit.ParticipantID
	Identity() livekit.ParticipantIdentity

	ToProto() *livekit.ParticipantInfo

	SetMetadata(metadata string)

	GetPublishedTrack(sid livekit.TrackID) MediaTrack
	GetPublishedTracks() []MediaTrack

	AddSubscriber(op LocalParticipant, params AddSubscriberParams) (int, error)
	RemoveSubscriber(op LocalParticipant, trackID livekit.TrackID, resume bool)

	// permissions
	Hidden() bool
	IsRecorder() bool

	Start()
	Close(sendLeave bool, reason ParticipantCloseReason) error

	SubscriptionPermission() *livekit.SubscriptionPermission

	// updates from remotes
	UpdateSubscriptionPermission(
		subscriptionPermission *livekit.SubscriptionPermission,
		resolverByIdentity func(participantIdentity livekit.ParticipantIdentity) LocalParticipant,
		resolverBySid func(participantID livekit.ParticipantID) LocalParticipant,
	) error
	UpdateVideoLayers(updateVideoLayers *livekit.UpdateVideoLayers) error
	UpdateSubscribedQuality(nodeID livekit.NodeID, trackID livekit.TrackID, maxQualities []SubscribedCodecQuality) error
	UpdateMediaLoss(nodeID livekit.NodeID, trackID livekit.TrackID, fractionalLoss uint32) error

	DebugInfo() map[string]interface{}
}

type IceConfig struct {
	PreferSubTcp bool
	PreferPubTcp bool
}

//counterfeiter:generate . LocalParticipant
type LocalParticipant interface {
	Participant

	GetLogger() logger.Logger
	GetAdaptiveStream() bool

	ProtocolVersion() ProtocolVersion

	ConnectedAt() time.Time

	State() livekit.ParticipantInfo_State
	IsReady() bool
	SubscriberAsPrimary() bool

	GetResponseSink() routing.MessageSink
	SetResponseSink(sink routing.MessageSink)

	// permissions
	ClaimGrants() *auth.ClaimGrants
	SetPermission(permission *livekit.ParticipantPermission) bool
	CanPublish() bool
	CanSubscribe() bool
	CanPublishData() bool

	AddICECandidate(candidate webrtc.ICECandidateInit, target livekit.SignalTarget) error

	HandleOffer(sdp webrtc.SessionDescription) (answer webrtc.SessionDescription, err error)

	AddTrack(req *livekit.AddTrackRequest)
	SetTrackMuted(trackID livekit.TrackID, muted bool, fromAdmin bool)

	SubscriberMediaEngine() *webrtc.MediaEngine
	SubscriberPC() *webrtc.PeerConnection
	HandleAnswer(sdp webrtc.SessionDescription) error
	Negotiate(force bool)
	ICERestart(iceConfig *IceConfig) error
	AddSubscribedTrack(st SubscribedTrack)
	RemoveSubscribedTrack(st SubscribedTrack)
	UpdateSubscribedTrackSettings(trackID livekit.TrackID, settings *livekit.UpdateTrackSettings) error
	GetSubscribedTracks() []SubscribedTrack

	// returns list of participant identities that the current participant is subscribed to
	GetSubscribedParticipants() []livekit.ParticipantID
	IsSubscribedTo(sid livekit.ParticipantID) bool
	IsPublisher() bool

	GetAudioLevel() (smoothedLevel float64, active bool)
	GetConnectionQuality() *livekit.ConnectionQualityInfo

	// server sent messages
	SendJoinResponse(info *livekit.Room, otherParticipants []*livekit.ParticipantInfo, iceServers []*livekit.ICEServer, region string) error
	SendParticipantUpdate(participants []*livekit.ParticipantInfo) error
	SendSpeakerUpdate(speakers []*livekit.SpeakerInfo) error
	SendDataPacket(packet *livekit.DataPacket) error
	SendRoomUpdate(room *livekit.Room) error
	SendConnectionQualityUpdate(update *livekit.ConnectionQualityUpdate) error
	SubscriptionPermissionUpdate(publisherID livekit.ParticipantID, trackID livekit.TrackID, allowed bool)
	SendRefreshToken(token string) error

	// callbacks
	OnStateChange(func(p LocalParticipant, oldState livekit.ParticipantInfo_State))
	// OnTrackPublished - remote added a track
	OnTrackPublished(func(LocalParticipant, MediaTrack))
	// OnTrackUpdated - one of its publishedTracks changed in status
	OnTrackUpdated(callback func(LocalParticipant, MediaTrack))
	// OnParticipantUpdate - metadata or permission is updated
	OnParticipantUpdate(callback func(LocalParticipant))
	OnDataPacket(callback func(LocalParticipant, *livekit.DataPacket))
	OnSubscribedTo(callback func(LocalParticipant, livekit.ParticipantID))
	OnClose(_callback func(LocalParticipant, map[livekit.TrackID]livekit.ParticipantID))
	OnClaimsChanged(_callback func(LocalParticipant))

	// session migration
	SetMigrateState(s MigrateState)
	MigrateState() MigrateState
	SetMigrateInfo(previousAnswer *webrtc.SessionDescription, mediaTracks []*livekit.TrackPublishedResponse, dataChannels []*livekit.DataChannelInfo)

	UpdateRTT(rtt uint32)

	CacheDownTrack(trackID livekit.TrackID, rtpTransceiver *webrtc.RTPTransceiver, forwarderState sfu.ForwarderState)
	UncacheDownTrack(rtpTransceiver *webrtc.RTPTransceiver)
	GetCachedDownTrack(trackID livekit.TrackID) (*webrtc.RTPTransceiver, sfu.ForwarderState)
}

// Room is a container of participants, and can provide room-level actions
//counterfeiter:generate . Room
type Room interface {
	Name() livekit.RoomName
	ID() livekit.RoomID
	RemoveParticipant(identity livekit.ParticipantIdentity, reason ParticipantCloseReason)
	UpdateSubscriptions(participant LocalParticipant, trackIDs []livekit.TrackID, participantTracks []*livekit.ParticipantTracks, subscribe bool) error
	UpdateSubscriptionPermission(participant LocalParticipant, permissions *livekit.SubscriptionPermission) error
	SyncState(participant LocalParticipant, state *livekit.SyncState) error
	SimulateScenario(participant LocalParticipant, scenario *livekit.SimulateScenario) error
	SetParticipantPermission(participant LocalParticipant, permission *livekit.ParticipantPermission) error
	UpdateVideoLayers(participant Participant, updateVideoLayers *livekit.UpdateVideoLayers) error
}

// MediaTrack represents a media track
//counterfeiter:generate . MediaTrack
type MediaTrack interface {
	ID() livekit.TrackID
	Kind() livekit.TrackType
	Name() string
	Source() livekit.TrackSource

	ToProto() *livekit.TrackInfo

	PublisherID() livekit.ParticipantID
	PublisherIdentity() livekit.ParticipantIdentity

	IsMuted() bool
	SetMuted(muted bool)

	UpdateVideoLayers(layers []*livekit.VideoLayer)
	IsSimulcast() bool

	Restart()

	// callbacks
	AddOnClose(func())

	// subscribers
	AddSubscriber(participant LocalParticipant) error
	RemoveSubscriber(participantID livekit.ParticipantID, willBeResumed bool)
	IsSubscriber(subID livekit.ParticipantID) bool
	RemoveAllSubscribers(willBeResumed bool)
	RevokeDisallowedSubscribers(allowedSubscriberIdentities []livekit.ParticipantIdentity) []livekit.ParticipantIdentity
	GetAllSubscribers() []livekit.ParticipantID
	GetNumSubscribers() int

	// returns quality information that's appropriate for width & height
	GetQualityForDimension(width, height uint32) livekit.VideoQuality

	NotifySubscriberNodeMaxQuality(nodeID livekit.NodeID, qualites []SubscribedCodecQuality)
	NotifySubscriberNodeMediaLoss(nodeID livekit.NodeID, fractionalLoss uint8)

	Receivers() []sfu.TrackReceiver
}

//counterfeiter:generate . LocalMediaTrack
type LocalMediaTrack interface {
	MediaTrack

	SignalCid() string
	HasSdpCid(cid string) bool

	GetAudioLevel() (level float64, active bool)
	GetConnectionScore() float32

	SetRTT(rtt uint32)
}

// MediaTrack is the main interface representing a track published to the room
//counterfeiter:generate . SubscribedTrack
type SubscribedTrack interface {
	OnBind(f func())
	ID() livekit.TrackID
	PublisherID() livekit.ParticipantID
	PublisherIdentity() livekit.ParticipantIdentity
	SubscriberID() livekit.ParticipantID
	SubscriberIdentity() livekit.ParticipantIdentity
	Subscriber() LocalParticipant
	DownTrack() *sfu.DownTrack
	MediaTrack() MediaTrack
	IsMuted() bool
	SetPublisherMuted(muted bool)
	UpdateSubscriberSettings(settings *livekit.UpdateTrackSettings)
	// selects appropriate video layer according to subscriber preferences
	UpdateVideoLayer()
}
