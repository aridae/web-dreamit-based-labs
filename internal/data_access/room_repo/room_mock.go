// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package roomrepo

import (
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

type RepositoryMock struct {
	// GetAllRoomsFunc mocks the GetAllRooms method.
	GetAllRoomsFunc func() ([]domain.Room, error)

	// GetRoomFunc mocks the GetRoom method.
	GetRoomFunc func(roomId int64) (*domain.Room, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetAllRooms holds details about calls to the GetAllRooms method.
		GetAllRooms []struct {
		}
		// GetRoom holds details about calls to the GetRoom method.
		GetRoom []struct {
			// RoomId is the roomId argument value.
			RoomId int64
		}
	}
	lockGetAllRooms sync.RWMutex
	lockGetRoom     sync.RWMutex
}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{
	}
}

// GetAllRooms calls GetAllRoomsFunc.
func (mock *RepositoryMock) GetAllRooms() ([]domain.Room, error) {
	if mock.GetAllRoomsFunc == nil {
		panic("RepositoryMock.GetAllRoomsFunc: method is nil but Repository.GetAllRooms was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetAllRooms.Lock()
	mock.calls.GetAllRooms = append(mock.calls.GetAllRooms, callInfo)
	mock.lockGetAllRooms.Unlock()
	return mock.GetAllRoomsFunc()
}

// GetAllRoomsCalls gets all the calls that were made to GetAllRooms.
// Check the length with:
//     len(mockedRepository.GetAllRoomsCalls())
func (mock *RepositoryMock) GetAllRoomsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAllRooms.RLock()
	calls = mock.calls.GetAllRooms
	mock.lockGetAllRooms.RUnlock()
	return calls
}

// GetRoom calls GetRoomFunc.
func (mock *RepositoryMock) GetRoom(roomId int64) (*domain.Room, error) {
	if mock.GetRoomFunc == nil {
		panic("RepositoryMock.GetRoomFunc: method is nil but Repository.GetRoom was just called")
	}
	callInfo := struct {
		RoomId int64
	}{
		RoomId: roomId,
	}
	mock.lockGetRoom.Lock()
	mock.calls.GetRoom = append(mock.calls.GetRoom, callInfo)
	mock.lockGetRoom.Unlock()
	return mock.GetRoomFunc(roomId)
}

// GetRoomCalls gets all the calls that were made to GetRoom.
// Check the length with:
//     len(mockedRepository.GetRoomCalls())
func (mock *RepositoryMock) GetRoomCalls() []struct {
	RoomId int64
} {
	var calls []struct {
		RoomId int64
	}
	mock.lockGetRoom.RLock()
	calls = mock.calls.GetRoom
	mock.lockGetRoom.RUnlock()
	return calls
}
