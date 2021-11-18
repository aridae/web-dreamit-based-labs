package room

import "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"

type UseCase interface {
	GetAllRooms() ([]api_models.Room, error)
	GetRoomEvents(roomId int64) ([]api_models.Event, error)
	UpdateRoomEvent(roomId int64, event api_models.Event) error
	AddRoomEvent(event api_models.Event) (int64, error)
	MyRoomEvents(userId uint64) ([]api_models.Event, error)
	DeleteRoomEvent(userId uint64, eventId int64) error
}
