// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/tomxiong/protocol/livekit"
)

type FakeParticipant struct {
	AddSubscriberStub        func(types.LocalParticipant, types.AddSubscriberParams) (int, error)
	addSubscriberMutex       sync.RWMutex
	addSubscriberArgsForCall []struct {
		arg1 types.LocalParticipant
		arg2 types.AddSubscriberParams
	}
	addSubscriberReturns struct {
		result1 int
		result2 error
	}
	addSubscriberReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	CloseStub        func(bool, types.ParticipantCloseReason) error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
		arg1 bool
		arg2 types.ParticipantCloseReason
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	DebugInfoStub        func() map[string]interface{}
	debugInfoMutex       sync.RWMutex
	debugInfoArgsForCall []struct {
	}
	debugInfoReturns struct {
		result1 map[string]interface{}
	}
	debugInfoReturnsOnCall map[int]struct {
		result1 map[string]interface{}
	}
	GetPublishedTrackStub        func(livekit.TrackID) types.MediaTrack
	getPublishedTrackMutex       sync.RWMutex
	getPublishedTrackArgsForCall []struct {
		arg1 livekit.TrackID
	}
	getPublishedTrackReturns struct {
		result1 types.MediaTrack
	}
	getPublishedTrackReturnsOnCall map[int]struct {
		result1 types.MediaTrack
	}
	GetPublishedTracksStub        func() []types.MediaTrack
	getPublishedTracksMutex       sync.RWMutex
	getPublishedTracksArgsForCall []struct {
	}
	getPublishedTracksReturns struct {
		result1 []types.MediaTrack
	}
	getPublishedTracksReturnsOnCall map[int]struct {
		result1 []types.MediaTrack
	}
	HiddenStub        func() bool
	hiddenMutex       sync.RWMutex
	hiddenArgsForCall []struct {
	}
	hiddenReturns struct {
		result1 bool
	}
	hiddenReturnsOnCall map[int]struct {
		result1 bool
	}
	IDStub        func() livekit.ParticipantID
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 livekit.ParticipantID
	}
	iDReturnsOnCall map[int]struct {
		result1 livekit.ParticipantID
	}
	IdentityStub        func() livekit.ParticipantIdentity
	identityMutex       sync.RWMutex
	identityArgsForCall []struct {
	}
	identityReturns struct {
		result1 livekit.ParticipantIdentity
	}
	identityReturnsOnCall map[int]struct {
		result1 livekit.ParticipantIdentity
	}
	IsRecorderStub        func() bool
	isRecorderMutex       sync.RWMutex
	isRecorderArgsForCall []struct {
	}
	isRecorderReturns struct {
		result1 bool
	}
	isRecorderReturnsOnCall map[int]struct {
		result1 bool
	}
	RemoveSubscriberStub        func(types.LocalParticipant, livekit.TrackID, bool)
	removeSubscriberMutex       sync.RWMutex
	removeSubscriberArgsForCall []struct {
		arg1 types.LocalParticipant
		arg2 livekit.TrackID
		arg3 bool
	}
	SetMetadataStub        func(string)
	setMetadataMutex       sync.RWMutex
	setMetadataArgsForCall []struct {
		arg1 string
	}
	StartStub        func()
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	SubscriptionPermissionStub        func() *livekit.SubscriptionPermission
	subscriptionPermissionMutex       sync.RWMutex
	subscriptionPermissionArgsForCall []struct {
	}
	subscriptionPermissionReturns struct {
		result1 *livekit.SubscriptionPermission
	}
	subscriptionPermissionReturnsOnCall map[int]struct {
		result1 *livekit.SubscriptionPermission
	}
	ToProtoStub        func() *livekit.ParticipantInfo
	toProtoMutex       sync.RWMutex
	toProtoArgsForCall []struct {
	}
	toProtoReturns struct {
		result1 *livekit.ParticipantInfo
	}
	toProtoReturnsOnCall map[int]struct {
		result1 *livekit.ParticipantInfo
	}
	UpdateMediaLossStub        func(livekit.NodeID, livekit.TrackID, uint32) error
	updateMediaLossMutex       sync.RWMutex
	updateMediaLossArgsForCall []struct {
		arg1 livekit.NodeID
		arg2 livekit.TrackID
		arg3 uint32
	}
	updateMediaLossReturns struct {
		result1 error
	}
	updateMediaLossReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateSubscribedQualityStub        func(livekit.NodeID, livekit.TrackID, []types.SubscribedCodecQuality) error
	updateSubscribedQualityMutex       sync.RWMutex
	updateSubscribedQualityArgsForCall []struct {
		arg1 livekit.NodeID
		arg2 livekit.TrackID
		arg3 []types.SubscribedCodecQuality
	}
	updateSubscribedQualityReturns struct {
		result1 error
	}
	updateSubscribedQualityReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateSubscriptionPermissionStub        func(*livekit.SubscriptionPermission, func(participantIdentity livekit.ParticipantIdentity) types.LocalParticipant, func(participantID livekit.ParticipantID) types.LocalParticipant) error
	updateSubscriptionPermissionMutex       sync.RWMutex
	updateSubscriptionPermissionArgsForCall []struct {
		arg1 *livekit.SubscriptionPermission
		arg2 func(participantIdentity livekit.ParticipantIdentity) types.LocalParticipant
		arg3 func(participantID livekit.ParticipantID) types.LocalParticipant
	}
	updateSubscriptionPermissionReturns struct {
		result1 error
	}
	updateSubscriptionPermissionReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateVideoLayersStub        func(*livekit.UpdateVideoLayers) error
	updateVideoLayersMutex       sync.RWMutex
	updateVideoLayersArgsForCall []struct {
		arg1 *livekit.UpdateVideoLayers
	}
	updateVideoLayersReturns struct {
		result1 error
	}
	updateVideoLayersReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeParticipant) AddSubscriber(arg1 types.LocalParticipant, arg2 types.AddSubscriberParams) (int, error) {
	fake.addSubscriberMutex.Lock()
	ret, specificReturn := fake.addSubscriberReturnsOnCall[len(fake.addSubscriberArgsForCall)]
	fake.addSubscriberArgsForCall = append(fake.addSubscriberArgsForCall, struct {
		arg1 types.LocalParticipant
		arg2 types.AddSubscriberParams
	}{arg1, arg2})
	stub := fake.AddSubscriberStub
	fakeReturns := fake.addSubscriberReturns
	fake.recordInvocation("AddSubscriber", []interface{}{arg1, arg2})
	fake.addSubscriberMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeParticipant) AddSubscriberCallCount() int {
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	return len(fake.addSubscriberArgsForCall)
}

func (fake *FakeParticipant) AddSubscriberCalls(stub func(types.LocalParticipant, types.AddSubscriberParams) (int, error)) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = stub
}

func (fake *FakeParticipant) AddSubscriberArgsForCall(i int) (types.LocalParticipant, types.AddSubscriberParams) {
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	argsForCall := fake.addSubscriberArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeParticipant) AddSubscriberReturns(result1 int, result2 error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = nil
	fake.addSubscriberReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeParticipant) AddSubscriberReturnsOnCall(i int, result1 int, result2 error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = nil
	if fake.addSubscriberReturnsOnCall == nil {
		fake.addSubscriberReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.addSubscriberReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeParticipant) Close(arg1 bool, arg2 types.ParticipantCloseReason) error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
		arg1 bool
		arg2 types.ParticipantCloseReason
	}{arg1, arg2})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{arg1, arg2})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeParticipant) CloseCalls(stub func(bool, types.ParticipantCloseReason) error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeParticipant) CloseArgsForCall(i int) (bool, types.ParticipantCloseReason) {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	argsForCall := fake.closeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeParticipant) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) DebugInfo() map[string]interface{} {
	fake.debugInfoMutex.Lock()
	ret, specificReturn := fake.debugInfoReturnsOnCall[len(fake.debugInfoArgsForCall)]
	fake.debugInfoArgsForCall = append(fake.debugInfoArgsForCall, struct {
	}{})
	stub := fake.DebugInfoStub
	fakeReturns := fake.debugInfoReturns
	fake.recordInvocation("DebugInfo", []interface{}{})
	fake.debugInfoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) DebugInfoCallCount() int {
	fake.debugInfoMutex.RLock()
	defer fake.debugInfoMutex.RUnlock()
	return len(fake.debugInfoArgsForCall)
}

func (fake *FakeParticipant) DebugInfoCalls(stub func() map[string]interface{}) {
	fake.debugInfoMutex.Lock()
	defer fake.debugInfoMutex.Unlock()
	fake.DebugInfoStub = stub
}

func (fake *FakeParticipant) DebugInfoReturns(result1 map[string]interface{}) {
	fake.debugInfoMutex.Lock()
	defer fake.debugInfoMutex.Unlock()
	fake.DebugInfoStub = nil
	fake.debugInfoReturns = struct {
		result1 map[string]interface{}
	}{result1}
}

func (fake *FakeParticipant) DebugInfoReturnsOnCall(i int, result1 map[string]interface{}) {
	fake.debugInfoMutex.Lock()
	defer fake.debugInfoMutex.Unlock()
	fake.DebugInfoStub = nil
	if fake.debugInfoReturnsOnCall == nil {
		fake.debugInfoReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
		})
	}
	fake.debugInfoReturnsOnCall[i] = struct {
		result1 map[string]interface{}
	}{result1}
}

func (fake *FakeParticipant) GetPublishedTrack(arg1 livekit.TrackID) types.MediaTrack {
	fake.getPublishedTrackMutex.Lock()
	ret, specificReturn := fake.getPublishedTrackReturnsOnCall[len(fake.getPublishedTrackArgsForCall)]
	fake.getPublishedTrackArgsForCall = append(fake.getPublishedTrackArgsForCall, struct {
		arg1 livekit.TrackID
	}{arg1})
	stub := fake.GetPublishedTrackStub
	fakeReturns := fake.getPublishedTrackReturns
	fake.recordInvocation("GetPublishedTrack", []interface{}{arg1})
	fake.getPublishedTrackMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) GetPublishedTrackCallCount() int {
	fake.getPublishedTrackMutex.RLock()
	defer fake.getPublishedTrackMutex.RUnlock()
	return len(fake.getPublishedTrackArgsForCall)
}

func (fake *FakeParticipant) GetPublishedTrackCalls(stub func(livekit.TrackID) types.MediaTrack) {
	fake.getPublishedTrackMutex.Lock()
	defer fake.getPublishedTrackMutex.Unlock()
	fake.GetPublishedTrackStub = stub
}

func (fake *FakeParticipant) GetPublishedTrackArgsForCall(i int) livekit.TrackID {
	fake.getPublishedTrackMutex.RLock()
	defer fake.getPublishedTrackMutex.RUnlock()
	argsForCall := fake.getPublishedTrackArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) GetPublishedTrackReturns(result1 types.MediaTrack) {
	fake.getPublishedTrackMutex.Lock()
	defer fake.getPublishedTrackMutex.Unlock()
	fake.GetPublishedTrackStub = nil
	fake.getPublishedTrackReturns = struct {
		result1 types.MediaTrack
	}{result1}
}

func (fake *FakeParticipant) GetPublishedTrackReturnsOnCall(i int, result1 types.MediaTrack) {
	fake.getPublishedTrackMutex.Lock()
	defer fake.getPublishedTrackMutex.Unlock()
	fake.GetPublishedTrackStub = nil
	if fake.getPublishedTrackReturnsOnCall == nil {
		fake.getPublishedTrackReturnsOnCall = make(map[int]struct {
			result1 types.MediaTrack
		})
	}
	fake.getPublishedTrackReturnsOnCall[i] = struct {
		result1 types.MediaTrack
	}{result1}
}

func (fake *FakeParticipant) GetPublishedTracks() []types.MediaTrack {
	fake.getPublishedTracksMutex.Lock()
	ret, specificReturn := fake.getPublishedTracksReturnsOnCall[len(fake.getPublishedTracksArgsForCall)]
	fake.getPublishedTracksArgsForCall = append(fake.getPublishedTracksArgsForCall, struct {
	}{})
	stub := fake.GetPublishedTracksStub
	fakeReturns := fake.getPublishedTracksReturns
	fake.recordInvocation("GetPublishedTracks", []interface{}{})
	fake.getPublishedTracksMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) GetPublishedTracksCallCount() int {
	fake.getPublishedTracksMutex.RLock()
	defer fake.getPublishedTracksMutex.RUnlock()
	return len(fake.getPublishedTracksArgsForCall)
}

func (fake *FakeParticipant) GetPublishedTracksCalls(stub func() []types.MediaTrack) {
	fake.getPublishedTracksMutex.Lock()
	defer fake.getPublishedTracksMutex.Unlock()
	fake.GetPublishedTracksStub = stub
}

func (fake *FakeParticipant) GetPublishedTracksReturns(result1 []types.MediaTrack) {
	fake.getPublishedTracksMutex.Lock()
	defer fake.getPublishedTracksMutex.Unlock()
	fake.GetPublishedTracksStub = nil
	fake.getPublishedTracksReturns = struct {
		result1 []types.MediaTrack
	}{result1}
}

func (fake *FakeParticipant) GetPublishedTracksReturnsOnCall(i int, result1 []types.MediaTrack) {
	fake.getPublishedTracksMutex.Lock()
	defer fake.getPublishedTracksMutex.Unlock()
	fake.GetPublishedTracksStub = nil
	if fake.getPublishedTracksReturnsOnCall == nil {
		fake.getPublishedTracksReturnsOnCall = make(map[int]struct {
			result1 []types.MediaTrack
		})
	}
	fake.getPublishedTracksReturnsOnCall[i] = struct {
		result1 []types.MediaTrack
	}{result1}
}

func (fake *FakeParticipant) Hidden() bool {
	fake.hiddenMutex.Lock()
	ret, specificReturn := fake.hiddenReturnsOnCall[len(fake.hiddenArgsForCall)]
	fake.hiddenArgsForCall = append(fake.hiddenArgsForCall, struct {
	}{})
	stub := fake.HiddenStub
	fakeReturns := fake.hiddenReturns
	fake.recordInvocation("Hidden", []interface{}{})
	fake.hiddenMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) HiddenCallCount() int {
	fake.hiddenMutex.RLock()
	defer fake.hiddenMutex.RUnlock()
	return len(fake.hiddenArgsForCall)
}

func (fake *FakeParticipant) HiddenCalls(stub func() bool) {
	fake.hiddenMutex.Lock()
	defer fake.hiddenMutex.Unlock()
	fake.HiddenStub = stub
}

func (fake *FakeParticipant) HiddenReturns(result1 bool) {
	fake.hiddenMutex.Lock()
	defer fake.hiddenMutex.Unlock()
	fake.HiddenStub = nil
	fake.hiddenReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParticipant) HiddenReturnsOnCall(i int, result1 bool) {
	fake.hiddenMutex.Lock()
	defer fake.hiddenMutex.Unlock()
	fake.HiddenStub = nil
	if fake.hiddenReturnsOnCall == nil {
		fake.hiddenReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hiddenReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParticipant) ID() livekit.ParticipantID {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	stub := fake.IDStub
	fakeReturns := fake.iDReturns
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeParticipant) IDCalls(stub func() livekit.ParticipantID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeParticipant) IDReturns(result1 livekit.ParticipantID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 livekit.ParticipantID
	}{result1}
}

func (fake *FakeParticipant) IDReturnsOnCall(i int, result1 livekit.ParticipantID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantID
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 livekit.ParticipantID
	}{result1}
}

func (fake *FakeParticipant) Identity() livekit.ParticipantIdentity {
	fake.identityMutex.Lock()
	ret, specificReturn := fake.identityReturnsOnCall[len(fake.identityArgsForCall)]
	fake.identityArgsForCall = append(fake.identityArgsForCall, struct {
	}{})
	stub := fake.IdentityStub
	fakeReturns := fake.identityReturns
	fake.recordInvocation("Identity", []interface{}{})
	fake.identityMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) IdentityCallCount() int {
	fake.identityMutex.RLock()
	defer fake.identityMutex.RUnlock()
	return len(fake.identityArgsForCall)
}

func (fake *FakeParticipant) IdentityCalls(stub func() livekit.ParticipantIdentity) {
	fake.identityMutex.Lock()
	defer fake.identityMutex.Unlock()
	fake.IdentityStub = stub
}

func (fake *FakeParticipant) IdentityReturns(result1 livekit.ParticipantIdentity) {
	fake.identityMutex.Lock()
	defer fake.identityMutex.Unlock()
	fake.IdentityStub = nil
	fake.identityReturns = struct {
		result1 livekit.ParticipantIdentity
	}{result1}
}

func (fake *FakeParticipant) IdentityReturnsOnCall(i int, result1 livekit.ParticipantIdentity) {
	fake.identityMutex.Lock()
	defer fake.identityMutex.Unlock()
	fake.IdentityStub = nil
	if fake.identityReturnsOnCall == nil {
		fake.identityReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantIdentity
		})
	}
	fake.identityReturnsOnCall[i] = struct {
		result1 livekit.ParticipantIdentity
	}{result1}
}

func (fake *FakeParticipant) IsRecorder() bool {
	fake.isRecorderMutex.Lock()
	ret, specificReturn := fake.isRecorderReturnsOnCall[len(fake.isRecorderArgsForCall)]
	fake.isRecorderArgsForCall = append(fake.isRecorderArgsForCall, struct {
	}{})
	stub := fake.IsRecorderStub
	fakeReturns := fake.isRecorderReturns
	fake.recordInvocation("IsRecorder", []interface{}{})
	fake.isRecorderMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) IsRecorderCallCount() int {
	fake.isRecorderMutex.RLock()
	defer fake.isRecorderMutex.RUnlock()
	return len(fake.isRecorderArgsForCall)
}

func (fake *FakeParticipant) IsRecorderCalls(stub func() bool) {
	fake.isRecorderMutex.Lock()
	defer fake.isRecorderMutex.Unlock()
	fake.IsRecorderStub = stub
}

func (fake *FakeParticipant) IsRecorderReturns(result1 bool) {
	fake.isRecorderMutex.Lock()
	defer fake.isRecorderMutex.Unlock()
	fake.IsRecorderStub = nil
	fake.isRecorderReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParticipant) IsRecorderReturnsOnCall(i int, result1 bool) {
	fake.isRecorderMutex.Lock()
	defer fake.isRecorderMutex.Unlock()
	fake.IsRecorderStub = nil
	if fake.isRecorderReturnsOnCall == nil {
		fake.isRecorderReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isRecorderReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParticipant) RemoveSubscriber(arg1 types.LocalParticipant, arg2 livekit.TrackID, arg3 bool) {
	fake.removeSubscriberMutex.Lock()
	fake.removeSubscriberArgsForCall = append(fake.removeSubscriberArgsForCall, struct {
		arg1 types.LocalParticipant
		arg2 livekit.TrackID
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.RemoveSubscriberStub
	fake.recordInvocation("RemoveSubscriber", []interface{}{arg1, arg2, arg3})
	fake.removeSubscriberMutex.Unlock()
	if stub != nil {
		fake.RemoveSubscriberStub(arg1, arg2, arg3)
	}
}

func (fake *FakeParticipant) RemoveSubscriberCallCount() int {
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	return len(fake.removeSubscriberArgsForCall)
}

func (fake *FakeParticipant) RemoveSubscriberCalls(stub func(types.LocalParticipant, livekit.TrackID, bool)) {
	fake.removeSubscriberMutex.Lock()
	defer fake.removeSubscriberMutex.Unlock()
	fake.RemoveSubscriberStub = stub
}

func (fake *FakeParticipant) RemoveSubscriberArgsForCall(i int) (types.LocalParticipant, livekit.TrackID, bool) {
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	argsForCall := fake.removeSubscriberArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeParticipant) SetMetadata(arg1 string) {
	fake.setMetadataMutex.Lock()
	fake.setMetadataArgsForCall = append(fake.setMetadataArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetMetadataStub
	fake.recordInvocation("SetMetadata", []interface{}{arg1})
	fake.setMetadataMutex.Unlock()
	if stub != nil {
		fake.SetMetadataStub(arg1)
	}
}

func (fake *FakeParticipant) SetMetadataCallCount() int {
	fake.setMetadataMutex.RLock()
	defer fake.setMetadataMutex.RUnlock()
	return len(fake.setMetadataArgsForCall)
}

func (fake *FakeParticipant) SetMetadataCalls(stub func(string)) {
	fake.setMetadataMutex.Lock()
	defer fake.setMetadataMutex.Unlock()
	fake.SetMetadataStub = stub
}

func (fake *FakeParticipant) SetMetadataArgsForCall(i int) string {
	fake.setMetadataMutex.RLock()
	defer fake.setMetadataMutex.RUnlock()
	argsForCall := fake.setMetadataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) Start() {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	stub := fake.StartStub
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if stub != nil {
		fake.StartStub()
	}
}

func (fake *FakeParticipant) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeParticipant) StartCalls(stub func()) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeParticipant) SubscriptionPermission() *livekit.SubscriptionPermission {
	fake.subscriptionPermissionMutex.Lock()
	ret, specificReturn := fake.subscriptionPermissionReturnsOnCall[len(fake.subscriptionPermissionArgsForCall)]
	fake.subscriptionPermissionArgsForCall = append(fake.subscriptionPermissionArgsForCall, struct {
	}{})
	stub := fake.SubscriptionPermissionStub
	fakeReturns := fake.subscriptionPermissionReturns
	fake.recordInvocation("SubscriptionPermission", []interface{}{})
	fake.subscriptionPermissionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) SubscriptionPermissionCallCount() int {
	fake.subscriptionPermissionMutex.RLock()
	defer fake.subscriptionPermissionMutex.RUnlock()
	return len(fake.subscriptionPermissionArgsForCall)
}

func (fake *FakeParticipant) SubscriptionPermissionCalls(stub func() *livekit.SubscriptionPermission) {
	fake.subscriptionPermissionMutex.Lock()
	defer fake.subscriptionPermissionMutex.Unlock()
	fake.SubscriptionPermissionStub = stub
}

func (fake *FakeParticipant) SubscriptionPermissionReturns(result1 *livekit.SubscriptionPermission) {
	fake.subscriptionPermissionMutex.Lock()
	defer fake.subscriptionPermissionMutex.Unlock()
	fake.SubscriptionPermissionStub = nil
	fake.subscriptionPermissionReturns = struct {
		result1 *livekit.SubscriptionPermission
	}{result1}
}

func (fake *FakeParticipant) SubscriptionPermissionReturnsOnCall(i int, result1 *livekit.SubscriptionPermission) {
	fake.subscriptionPermissionMutex.Lock()
	defer fake.subscriptionPermissionMutex.Unlock()
	fake.SubscriptionPermissionStub = nil
	if fake.subscriptionPermissionReturnsOnCall == nil {
		fake.subscriptionPermissionReturnsOnCall = make(map[int]struct {
			result1 *livekit.SubscriptionPermission
		})
	}
	fake.subscriptionPermissionReturnsOnCall[i] = struct {
		result1 *livekit.SubscriptionPermission
	}{result1}
}

func (fake *FakeParticipant) ToProto() *livekit.ParticipantInfo {
	fake.toProtoMutex.Lock()
	ret, specificReturn := fake.toProtoReturnsOnCall[len(fake.toProtoArgsForCall)]
	fake.toProtoArgsForCall = append(fake.toProtoArgsForCall, struct {
	}{})
	stub := fake.ToProtoStub
	fakeReturns := fake.toProtoReturns
	fake.recordInvocation("ToProto", []interface{}{})
	fake.toProtoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) ToProtoCallCount() int {
	fake.toProtoMutex.RLock()
	defer fake.toProtoMutex.RUnlock()
	return len(fake.toProtoArgsForCall)
}

func (fake *FakeParticipant) ToProtoCalls(stub func() *livekit.ParticipantInfo) {
	fake.toProtoMutex.Lock()
	defer fake.toProtoMutex.Unlock()
	fake.ToProtoStub = stub
}

func (fake *FakeParticipant) ToProtoReturns(result1 *livekit.ParticipantInfo) {
	fake.toProtoMutex.Lock()
	defer fake.toProtoMutex.Unlock()
	fake.ToProtoStub = nil
	fake.toProtoReturns = struct {
		result1 *livekit.ParticipantInfo
	}{result1}
}

func (fake *FakeParticipant) ToProtoReturnsOnCall(i int, result1 *livekit.ParticipantInfo) {
	fake.toProtoMutex.Lock()
	defer fake.toProtoMutex.Unlock()
	fake.ToProtoStub = nil
	if fake.toProtoReturnsOnCall == nil {
		fake.toProtoReturnsOnCall = make(map[int]struct {
			result1 *livekit.ParticipantInfo
		})
	}
	fake.toProtoReturnsOnCall[i] = struct {
		result1 *livekit.ParticipantInfo
	}{result1}
}

func (fake *FakeParticipant) UpdateMediaLoss(arg1 livekit.NodeID, arg2 livekit.TrackID, arg3 uint32) error {
	fake.updateMediaLossMutex.Lock()
	ret, specificReturn := fake.updateMediaLossReturnsOnCall[len(fake.updateMediaLossArgsForCall)]
	fake.updateMediaLossArgsForCall = append(fake.updateMediaLossArgsForCall, struct {
		arg1 livekit.NodeID
		arg2 livekit.TrackID
		arg3 uint32
	}{arg1, arg2, arg3})
	stub := fake.UpdateMediaLossStub
	fakeReturns := fake.updateMediaLossReturns
	fake.recordInvocation("UpdateMediaLoss", []interface{}{arg1, arg2, arg3})
	fake.updateMediaLossMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) UpdateMediaLossCallCount() int {
	fake.updateMediaLossMutex.RLock()
	defer fake.updateMediaLossMutex.RUnlock()
	return len(fake.updateMediaLossArgsForCall)
}

func (fake *FakeParticipant) UpdateMediaLossCalls(stub func(livekit.NodeID, livekit.TrackID, uint32) error) {
	fake.updateMediaLossMutex.Lock()
	defer fake.updateMediaLossMutex.Unlock()
	fake.UpdateMediaLossStub = stub
}

func (fake *FakeParticipant) UpdateMediaLossArgsForCall(i int) (livekit.NodeID, livekit.TrackID, uint32) {
	fake.updateMediaLossMutex.RLock()
	defer fake.updateMediaLossMutex.RUnlock()
	argsForCall := fake.updateMediaLossArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeParticipant) UpdateMediaLossReturns(result1 error) {
	fake.updateMediaLossMutex.Lock()
	defer fake.updateMediaLossMutex.Unlock()
	fake.UpdateMediaLossStub = nil
	fake.updateMediaLossReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) UpdateMediaLossReturnsOnCall(i int, result1 error) {
	fake.updateMediaLossMutex.Lock()
	defer fake.updateMediaLossMutex.Unlock()
	fake.UpdateMediaLossStub = nil
	if fake.updateMediaLossReturnsOnCall == nil {
		fake.updateMediaLossReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateMediaLossReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) UpdateSubscribedQuality(arg1 livekit.NodeID, arg2 livekit.TrackID, arg3 []types.SubscribedCodecQuality) error {
	var arg3Copy []types.SubscribedCodecQuality
	if arg3 != nil {
		arg3Copy = make([]types.SubscribedCodecQuality, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.updateSubscribedQualityMutex.Lock()
	ret, specificReturn := fake.updateSubscribedQualityReturnsOnCall[len(fake.updateSubscribedQualityArgsForCall)]
	fake.updateSubscribedQualityArgsForCall = append(fake.updateSubscribedQualityArgsForCall, struct {
		arg1 livekit.NodeID
		arg2 livekit.TrackID
		arg3 []types.SubscribedCodecQuality
	}{arg1, arg2, arg3Copy})
	stub := fake.UpdateSubscribedQualityStub
	fakeReturns := fake.updateSubscribedQualityReturns
	fake.recordInvocation("UpdateSubscribedQuality", []interface{}{arg1, arg2, arg3Copy})
	fake.updateSubscribedQualityMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) UpdateSubscribedQualityCallCount() int {
	fake.updateSubscribedQualityMutex.RLock()
	defer fake.updateSubscribedQualityMutex.RUnlock()
	return len(fake.updateSubscribedQualityArgsForCall)
}

func (fake *FakeParticipant) UpdateSubscribedQualityCalls(stub func(livekit.NodeID, livekit.TrackID, []types.SubscribedCodecQuality) error) {
	fake.updateSubscribedQualityMutex.Lock()
	defer fake.updateSubscribedQualityMutex.Unlock()
	fake.UpdateSubscribedQualityStub = stub
}

func (fake *FakeParticipant) UpdateSubscribedQualityArgsForCall(i int) (livekit.NodeID, livekit.TrackID, []types.SubscribedCodecQuality) {
	fake.updateSubscribedQualityMutex.RLock()
	defer fake.updateSubscribedQualityMutex.RUnlock()
	argsForCall := fake.updateSubscribedQualityArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeParticipant) UpdateSubscribedQualityReturns(result1 error) {
	fake.updateSubscribedQualityMutex.Lock()
	defer fake.updateSubscribedQualityMutex.Unlock()
	fake.UpdateSubscribedQualityStub = nil
	fake.updateSubscribedQualityReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) UpdateSubscribedQualityReturnsOnCall(i int, result1 error) {
	fake.updateSubscribedQualityMutex.Lock()
	defer fake.updateSubscribedQualityMutex.Unlock()
	fake.UpdateSubscribedQualityStub = nil
	if fake.updateSubscribedQualityReturnsOnCall == nil {
		fake.updateSubscribedQualityReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSubscribedQualityReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) UpdateSubscriptionPermission(arg1 *livekit.SubscriptionPermission, arg2 func(participantIdentity livekit.ParticipantIdentity) types.LocalParticipant, arg3 func(participantID livekit.ParticipantID) types.LocalParticipant) error {
	fake.updateSubscriptionPermissionMutex.Lock()
	ret, specificReturn := fake.updateSubscriptionPermissionReturnsOnCall[len(fake.updateSubscriptionPermissionArgsForCall)]
	fake.updateSubscriptionPermissionArgsForCall = append(fake.updateSubscriptionPermissionArgsForCall, struct {
		arg1 *livekit.SubscriptionPermission
		arg2 func(participantIdentity livekit.ParticipantIdentity) types.LocalParticipant
		arg3 func(participantID livekit.ParticipantID) types.LocalParticipant
	}{arg1, arg2, arg3})
	stub := fake.UpdateSubscriptionPermissionStub
	fakeReturns := fake.updateSubscriptionPermissionReturns
	fake.recordInvocation("UpdateSubscriptionPermission", []interface{}{arg1, arg2, arg3})
	fake.updateSubscriptionPermissionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) UpdateSubscriptionPermissionCallCount() int {
	fake.updateSubscriptionPermissionMutex.RLock()
	defer fake.updateSubscriptionPermissionMutex.RUnlock()
	return len(fake.updateSubscriptionPermissionArgsForCall)
}

func (fake *FakeParticipant) UpdateSubscriptionPermissionCalls(stub func(*livekit.SubscriptionPermission, func(participantIdentity livekit.ParticipantIdentity) types.LocalParticipant, func(participantID livekit.ParticipantID) types.LocalParticipant) error) {
	fake.updateSubscriptionPermissionMutex.Lock()
	defer fake.updateSubscriptionPermissionMutex.Unlock()
	fake.UpdateSubscriptionPermissionStub = stub
}

func (fake *FakeParticipant) UpdateSubscriptionPermissionArgsForCall(i int) (*livekit.SubscriptionPermission, func(participantIdentity livekit.ParticipantIdentity) types.LocalParticipant, func(participantID livekit.ParticipantID) types.LocalParticipant) {
	fake.updateSubscriptionPermissionMutex.RLock()
	defer fake.updateSubscriptionPermissionMutex.RUnlock()
	argsForCall := fake.updateSubscriptionPermissionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeParticipant) UpdateSubscriptionPermissionReturns(result1 error) {
	fake.updateSubscriptionPermissionMutex.Lock()
	defer fake.updateSubscriptionPermissionMutex.Unlock()
	fake.UpdateSubscriptionPermissionStub = nil
	fake.updateSubscriptionPermissionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) UpdateSubscriptionPermissionReturnsOnCall(i int, result1 error) {
	fake.updateSubscriptionPermissionMutex.Lock()
	defer fake.updateSubscriptionPermissionMutex.Unlock()
	fake.UpdateSubscriptionPermissionStub = nil
	if fake.updateSubscriptionPermissionReturnsOnCall == nil {
		fake.updateSubscriptionPermissionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSubscriptionPermissionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) UpdateVideoLayers(arg1 *livekit.UpdateVideoLayers) error {
	fake.updateVideoLayersMutex.Lock()
	ret, specificReturn := fake.updateVideoLayersReturnsOnCall[len(fake.updateVideoLayersArgsForCall)]
	fake.updateVideoLayersArgsForCall = append(fake.updateVideoLayersArgsForCall, struct {
		arg1 *livekit.UpdateVideoLayers
	}{arg1})
	stub := fake.UpdateVideoLayersStub
	fakeReturns := fake.updateVideoLayersReturns
	fake.recordInvocation("UpdateVideoLayers", []interface{}{arg1})
	fake.updateVideoLayersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParticipant) UpdateVideoLayersCallCount() int {
	fake.updateVideoLayersMutex.RLock()
	defer fake.updateVideoLayersMutex.RUnlock()
	return len(fake.updateVideoLayersArgsForCall)
}

func (fake *FakeParticipant) UpdateVideoLayersCalls(stub func(*livekit.UpdateVideoLayers) error) {
	fake.updateVideoLayersMutex.Lock()
	defer fake.updateVideoLayersMutex.Unlock()
	fake.UpdateVideoLayersStub = stub
}

func (fake *FakeParticipant) UpdateVideoLayersArgsForCall(i int) *livekit.UpdateVideoLayers {
	fake.updateVideoLayersMutex.RLock()
	defer fake.updateVideoLayersMutex.RUnlock()
	argsForCall := fake.updateVideoLayersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParticipant) UpdateVideoLayersReturns(result1 error) {
	fake.updateVideoLayersMutex.Lock()
	defer fake.updateVideoLayersMutex.Unlock()
	fake.UpdateVideoLayersStub = nil
	fake.updateVideoLayersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) UpdateVideoLayersReturnsOnCall(i int, result1 error) {
	fake.updateVideoLayersMutex.Lock()
	defer fake.updateVideoLayersMutex.Unlock()
	fake.UpdateVideoLayersStub = nil
	if fake.updateVideoLayersReturnsOnCall == nil {
		fake.updateVideoLayersReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateVideoLayersReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeParticipant) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.debugInfoMutex.RLock()
	defer fake.debugInfoMutex.RUnlock()
	fake.getPublishedTrackMutex.RLock()
	defer fake.getPublishedTrackMutex.RUnlock()
	fake.getPublishedTracksMutex.RLock()
	defer fake.getPublishedTracksMutex.RUnlock()
	fake.hiddenMutex.RLock()
	defer fake.hiddenMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.identityMutex.RLock()
	defer fake.identityMutex.RUnlock()
	fake.isRecorderMutex.RLock()
	defer fake.isRecorderMutex.RUnlock()
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	fake.setMetadataMutex.RLock()
	defer fake.setMetadataMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.subscriptionPermissionMutex.RLock()
	defer fake.subscriptionPermissionMutex.RUnlock()
	fake.toProtoMutex.RLock()
	defer fake.toProtoMutex.RUnlock()
	fake.updateMediaLossMutex.RLock()
	defer fake.updateMediaLossMutex.RUnlock()
	fake.updateSubscribedQualityMutex.RLock()
	defer fake.updateSubscribedQualityMutex.RUnlock()
	fake.updateSubscriptionPermissionMutex.RLock()
	defer fake.updateSubscriptionPermissionMutex.RUnlock()
	fake.updateVideoLayersMutex.RLock()
	defer fake.updateVideoLayersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeParticipant) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ types.Participant = new(FakeParticipant)
