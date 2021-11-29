package controllers

import (
	domain "github.com/aridae/web-dreamit-api-based-labs/internal/domain"

	roomrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/room_repo"
)

type RoomController struct {
	RoomRepo roomrepo.Repository
}

func NewRoomController(RoomRepo roomrepo.Repository) *RoomController {
	return &RoomController{
		RoomRepo: RoomRepo,
	}
}

func (r RoomController) GetRoom(roomId int64) (*domain.Room, error) {
	return r.RoomRepo.GetRoom(roomId)
}
func (r RoomController) GetAllRooms() ([]domain.Room, error) {
	return r.RoomRepo.GetAllRooms()
}
