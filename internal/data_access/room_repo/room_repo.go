package roomrepo

import "github.com/aridae/web-dreamit-api-based-labs/internal/domain"

type Repository interface {
	GetRoom(roomId int64) (*domain.Room, error)
	GetAllRooms() ([]domain.Room, error)
}
