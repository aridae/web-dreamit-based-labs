package eventrepo

import (
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

type Repository interface {
	GetEvent(eventId int64) (*domain.Event, error)
	GetEvents() ([]domain.Event, error)
	GetRoomEventsByRoomId(roomId int64) ([]domain.Event, error)
	RescheduleRoomEvent(eventId int64, event domain.PatchEvent) error
	AddRoomEvent(event domain.PostEvent) (int64, error)
	GetRoomEventsByUserId(userId uint64) ([]domain.Event, error)
	DeleteRoomEvent(eventId int64) error
}
