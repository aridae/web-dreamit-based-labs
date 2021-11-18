package room

import "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"

type Repository interface {
	GetAllRooms() ([]api_models.Room, error)
	GetRoomEventsByRoomId(roomId int64) ([]api_models.Event, error)
	UpdateRoomEventsByRoomId(roomId int64, event api_models.Event) error
	AddRoomEvent(event api_models.Event) (int64, error)
	GetRoomEventsByUserId(userId uint64) ([]api_models.Event, error)
	DeleteRoomEvent(userId uint64, eventId int64) error
}
