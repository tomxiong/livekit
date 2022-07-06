// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"context"
	"sync"

	"github.com/tomxiong/livekit/pkg/service"
	"github.com/tomxiong/protocol/livekit"
)

type FakeRoomAllocator struct {
	CreateRoomStub        func(context.Context, *livekit.CreateRoomRequest) (*livekit.Room, error)
	createRoomMutex       sync.RWMutex
	createRoomArgsForCall []struct {
		arg1 context.Context
		arg2 *livekit.CreateRoomRequest
	}
	createRoomReturns struct {
		result1 *livekit.Room
		result2 error
	}
	createRoomReturnsOnCall map[int]struct {
		result1 *livekit.Room
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRoomAllocator) CreateRoom(arg1 context.Context, arg2 *livekit.CreateRoomRequest) (*livekit.Room, error) {
	fake.createRoomMutex.Lock()
	ret, specificReturn := fake.createRoomReturnsOnCall[len(fake.createRoomArgsForCall)]
	fake.createRoomArgsForCall = append(fake.createRoomArgsForCall, struct {
		arg1 context.Context
		arg2 *livekit.CreateRoomRequest
	}{arg1, arg2})
	stub := fake.CreateRoomStub
	fakeReturns := fake.createRoomReturns
	fake.recordInvocation("CreateRoom", []interface{}{arg1, arg2})
	fake.createRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRoomAllocator) CreateRoomCallCount() int {
	fake.createRoomMutex.RLock()
	defer fake.createRoomMutex.RUnlock()
	return len(fake.createRoomArgsForCall)
}

func (fake *FakeRoomAllocator) CreateRoomCalls(stub func(context.Context, *livekit.CreateRoomRequest) (*livekit.Room, error)) {
	fake.createRoomMutex.Lock()
	defer fake.createRoomMutex.Unlock()
	fake.CreateRoomStub = stub
}

func (fake *FakeRoomAllocator) CreateRoomArgsForCall(i int) (context.Context, *livekit.CreateRoomRequest) {
	fake.createRoomMutex.RLock()
	defer fake.createRoomMutex.RUnlock()
	argsForCall := fake.createRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRoomAllocator) CreateRoomReturns(result1 *livekit.Room, result2 error) {
	fake.createRoomMutex.Lock()
	defer fake.createRoomMutex.Unlock()
	fake.CreateRoomStub = nil
	fake.createRoomReturns = struct {
		result1 *livekit.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeRoomAllocator) CreateRoomReturnsOnCall(i int, result1 *livekit.Room, result2 error) {
	fake.createRoomMutex.Lock()
	defer fake.createRoomMutex.Unlock()
	fake.CreateRoomStub = nil
	if fake.createRoomReturnsOnCall == nil {
		fake.createRoomReturnsOnCall = make(map[int]struct {
			result1 *livekit.Room
			result2 error
		})
	}
	fake.createRoomReturnsOnCall[i] = struct {
		result1 *livekit.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeRoomAllocator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRoomMutex.RLock()
	defer fake.createRoomMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRoomAllocator) recordInvocation(key string, args []interface{}) {
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

var _ service.RoomAllocator = new(FakeRoomAllocator)
