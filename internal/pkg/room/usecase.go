package room

import (
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/models"
	"time"
)

type UseCase interface {
	GetAllRooms() ([]models.Room, error)
	GetRoomCalendar(roomId int64) ([]models.Event, error)
	GetRoomSchedule(roomId int64, date time.Time) (*models.Schedule, error)
	UpdateRoomBooking(roomId int64, event models.Event) error
	AddRoomBooking(roomId int64, event models.Event) (int64, error)
	MyRoomBooking(userId uint64) ([]models.Booking, error)
	DeleteRoomBooking(userId uint64, eventId int64) error
}
