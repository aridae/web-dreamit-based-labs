// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package eventrepo

import (
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			AddRoomEventFunc: func(event domain.PostEvent) (int64, error) {
// 				panic("mock out the AddRoomEvent method")
// 			},
// 			DeleteRoomEventFunc: func(eventId int64) error {
// 				panic("mock out the DeleteRoomEvent method")
// 			},
// 			GetEventFunc: func(eventId int64) (*domain.Event, error) {
// 				panic("mock out the GetEvent method")
// 			},
// 			GetEventsFunc: func() ([]domain.Event, error) {
// 				panic("mock out the GetEvents method")
// 			},
// 			GetRoomEventsByRoomIdFunc: func(roomId int64) ([]domain.Event, error) {
// 				panic("mock out the GetRoomEventsByRoomId method")
// 			},
// 			GetRoomEventsByUserIdFunc: func(userId uint64) ([]domain.Event, error) {
// 				panic("mock out the GetRoomEventsByUserId method")
// 			},
// 			RescheduleRoomEventFunc: func(eventId int64, event domain.PatchEvent) error {
// 				panic("mock out the RescheduleRoomEvent method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// AddRoomEventFunc mocks the AddRoomEvent method.
	AddRoomEventFunc func(event domain.PostEvent) (int64, error)

	// DeleteRoomEventFunc mocks the DeleteRoomEvent method.
	DeleteRoomEventFunc func(eventId int64) error

	// GetEventFunc mocks the GetEvent method.
	GetEventFunc func(eventId int64) (*domain.Event, error)

	// GetEventsFunc mocks the GetEvents method.
	GetEventsFunc func() ([]domain.Event, error)

	// GetRoomEventsByRoomIdFunc mocks the GetRoomEventsByRoomId method.
	GetRoomEventsByRoomIdFunc func(roomId int64) ([]domain.Event, error)

	// GetRoomEventsByUserIdFunc mocks the GetRoomEventsByUserId method.
	GetRoomEventsByUserIdFunc func(userId uint64) ([]domain.Event, error)

	// RescheduleRoomEventFunc mocks the RescheduleRoomEvent method.
	RescheduleRoomEventFunc func(eventId int64, event domain.PatchEvent) error

	// calls tracks calls to the methods.
	calls struct {
		// AddRoomEvent holds details about calls to the AddRoomEvent method.
		AddRoomEvent []struct {
			// Event is the event argument value.
			Event domain.PostEvent
		}
		// DeleteRoomEvent holds details about calls to the DeleteRoomEvent method.
		DeleteRoomEvent []struct {
			// EventId is the eventId argument value.
			EventId int64
		}
		// GetEvent holds details about calls to the GetEvent method.
		GetEvent []struct {
			// EventId is the eventId argument value.
			EventId int64
		}
		// GetEvents holds details about calls to the GetEvents method.
		GetEvents []struct {
		}
		// GetRoomEventsByRoomId holds details about calls to the GetRoomEventsByRoomId method.
		GetRoomEventsByRoomId []struct {
			// RoomId is the roomId argument value.
			RoomId int64
		}
		// GetRoomEventsByUserId holds details about calls to the GetRoomEventsByUserId method.
		GetRoomEventsByUserId []struct {
			// UserId is the userId argument value.
			UserId uint64
		}
		// RescheduleRoomEvent holds details about calls to the RescheduleRoomEvent method.
		RescheduleRoomEvent []struct {
			// EventId is the eventId argument value.
			EventId int64
			// Event is the event argument value.
			Event domain.PatchEvent
		}
	}
	lockAddRoomEvent          sync.RWMutex
	lockDeleteRoomEvent       sync.RWMutex
	lockGetEvent              sync.RWMutex
	lockGetEvents             sync.RWMutex
	lockGetRoomEventsByRoomId sync.RWMutex
	lockGetRoomEventsByUserId sync.RWMutex
	lockRescheduleRoomEvent   sync.RWMutex
}

// AddRoomEvent calls AddRoomEventFunc.
func (mock *RepositoryMock) AddRoomEvent(event domain.PostEvent) (int64, error) {
	if mock.AddRoomEventFunc == nil {
		panic("RepositoryMock.AddRoomEventFunc: method is nil but Repository.AddRoomEvent was just called")
	}
	callInfo := struct {
		Event domain.PostEvent
	}{
		Event: event,
	}
	mock.lockAddRoomEvent.Lock()
	mock.calls.AddRoomEvent = append(mock.calls.AddRoomEvent, callInfo)
	mock.lockAddRoomEvent.Unlock()
	return mock.AddRoomEventFunc(event)
}

// AddRoomEventCalls gets all the calls that were made to AddRoomEvent.
// Check the length with:
//     len(mockedRepository.AddRoomEventCalls())
func (mock *RepositoryMock) AddRoomEventCalls() []struct {
	Event domain.PostEvent
} {
	var calls []struct {
		Event domain.PostEvent
	}
	mock.lockAddRoomEvent.RLock()
	calls = mock.calls.AddRoomEvent
	mock.lockAddRoomEvent.RUnlock()
	return calls
}

// DeleteRoomEvent calls DeleteRoomEventFunc.
func (mock *RepositoryMock) DeleteRoomEvent(eventId int64) error {
	if mock.DeleteRoomEventFunc == nil {
		panic("RepositoryMock.DeleteRoomEventFunc: method is nil but Repository.DeleteRoomEvent was just called")
	}
	callInfo := struct {
		EventId int64
	}{
		EventId: eventId,
	}
	mock.lockDeleteRoomEvent.Lock()
	mock.calls.DeleteRoomEvent = append(mock.calls.DeleteRoomEvent, callInfo)
	mock.lockDeleteRoomEvent.Unlock()
	return mock.DeleteRoomEventFunc(eventId)
}

// DeleteRoomEventCalls gets all the calls that were made to DeleteRoomEvent.
// Check the length with:
//     len(mockedRepository.DeleteRoomEventCalls())
func (mock *RepositoryMock) DeleteRoomEventCalls() []struct {
	EventId int64
} {
	var calls []struct {
		EventId int64
	}
	mock.lockDeleteRoomEvent.RLock()
	calls = mock.calls.DeleteRoomEvent
	mock.lockDeleteRoomEvent.RUnlock()
	return calls
}

// GetEvent calls GetEventFunc.
func (mock *RepositoryMock) GetEvent(eventId int64) (*domain.Event, error) {
	if mock.GetEventFunc == nil {
		panic("RepositoryMock.GetEventFunc: method is nil but Repository.GetEvent was just called")
	}
	callInfo := struct {
		EventId int64
	}{
		EventId: eventId,
	}
	mock.lockGetEvent.Lock()
	mock.calls.GetEvent = append(mock.calls.GetEvent, callInfo)
	mock.lockGetEvent.Unlock()
	return mock.GetEventFunc(eventId)
}

// GetEventCalls gets all the calls that were made to GetEvent.
// Check the length with:
//     len(mockedRepository.GetEventCalls())
func (mock *RepositoryMock) GetEventCalls() []struct {
	EventId int64
} {
	var calls []struct {
		EventId int64
	}
	mock.lockGetEvent.RLock()
	calls = mock.calls.GetEvent
	mock.lockGetEvent.RUnlock()
	return calls
}

// GetEvents calls GetEventsFunc.
func (mock *RepositoryMock) GetEvents() ([]domain.Event, error) {
	if mock.GetEventsFunc == nil {
		panic("RepositoryMock.GetEventsFunc: method is nil but Repository.GetEvents was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetEvents.Lock()
	mock.calls.GetEvents = append(mock.calls.GetEvents, callInfo)
	mock.lockGetEvents.Unlock()
	return mock.GetEventsFunc()
}

// GetEventsCalls gets all the calls that were made to GetEvents.
// Check the length with:
//     len(mockedRepository.GetEventsCalls())
func (mock *RepositoryMock) GetEventsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetEvents.RLock()
	calls = mock.calls.GetEvents
	mock.lockGetEvents.RUnlock()
	return calls
}

// GetRoomEventsByRoomId calls GetRoomEventsByRoomIdFunc.
func (mock *RepositoryMock) GetRoomEventsByRoomId(roomId int64) ([]domain.Event, error) {
	if mock.GetRoomEventsByRoomIdFunc == nil {
		panic("RepositoryMock.GetRoomEventsByRoomIdFunc: method is nil but Repository.GetRoomEventsByRoomId was just called")
	}
	callInfo := struct {
		RoomId int64
	}{
		RoomId: roomId,
	}
	mock.lockGetRoomEventsByRoomId.Lock()
	mock.calls.GetRoomEventsByRoomId = append(mock.calls.GetRoomEventsByRoomId, callInfo)
	mock.lockGetRoomEventsByRoomId.Unlock()
	return mock.GetRoomEventsByRoomIdFunc(roomId)
}

// GetRoomEventsByRoomIdCalls gets all the calls that were made to GetRoomEventsByRoomId.
// Check the length with:
//     len(mockedRepository.GetRoomEventsByRoomIdCalls())
func (mock *RepositoryMock) GetRoomEventsByRoomIdCalls() []struct {
	RoomId int64
} {
	var calls []struct {
		RoomId int64
	}
	mock.lockGetRoomEventsByRoomId.RLock()
	calls = mock.calls.GetRoomEventsByRoomId
	mock.lockGetRoomEventsByRoomId.RUnlock()
	return calls
}

// GetRoomEventsByUserId calls GetRoomEventsByUserIdFunc.
func (mock *RepositoryMock) GetRoomEventsByUserId(userId uint64) ([]domain.Event, error) {
	if mock.GetRoomEventsByUserIdFunc == nil {
		panic("RepositoryMock.GetRoomEventsByUserIdFunc: method is nil but Repository.GetRoomEventsByUserId was just called")
	}
	callInfo := struct {
		UserId uint64
	}{
		UserId: userId,
	}
	mock.lockGetRoomEventsByUserId.Lock()
	mock.calls.GetRoomEventsByUserId = append(mock.calls.GetRoomEventsByUserId, callInfo)
	mock.lockGetRoomEventsByUserId.Unlock()
	return mock.GetRoomEventsByUserIdFunc(userId)
}

// GetRoomEventsByUserIdCalls gets all the calls that were made to GetRoomEventsByUserId.
// Check the length with:
//     len(mockedRepository.GetRoomEventsByUserIdCalls())
func (mock *RepositoryMock) GetRoomEventsByUserIdCalls() []struct {
	UserId uint64
} {
	var calls []struct {
		UserId uint64
	}
	mock.lockGetRoomEventsByUserId.RLock()
	calls = mock.calls.GetRoomEventsByUserId
	mock.lockGetRoomEventsByUserId.RUnlock()
	return calls
}

// RescheduleRoomEvent calls RescheduleRoomEventFunc.
func (mock *RepositoryMock) RescheduleRoomEvent(eventId int64, event domain.PatchEvent) error {
	if mock.RescheduleRoomEventFunc == nil {
		panic("RepositoryMock.RescheduleRoomEventFunc: method is nil but Repository.RescheduleRoomEvent was just called")
	}
	callInfo := struct {
		EventId int64
		Event   domain.PatchEvent
	}{
		EventId: eventId,
		Event:   event,
	}
	mock.lockRescheduleRoomEvent.Lock()
	mock.calls.RescheduleRoomEvent = append(mock.calls.RescheduleRoomEvent, callInfo)
	mock.lockRescheduleRoomEvent.Unlock()
	return mock.RescheduleRoomEventFunc(eventId, event)
}

// RescheduleRoomEventCalls gets all the calls that were made to RescheduleRoomEvent.
// Check the length with:
//     len(mockedRepository.RescheduleRoomEventCalls())
func (mock *RepositoryMock) RescheduleRoomEventCalls() []struct {
	EventId int64
	Event   domain.PatchEvent
} {
	var calls []struct {
		EventId int64
		Event   domain.PatchEvent
	}
	mock.lockRescheduleRoomEvent.RLock()
	calls = mock.calls.RescheduleRoomEvent
	mock.lockRescheduleRoomEvent.RUnlock()
	return calls
}
