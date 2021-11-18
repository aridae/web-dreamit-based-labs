package controllers

import (
	"fmt"

	eventrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/event_repo"
	domain "github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

type EventUseCase interface {
	GetEvents() ([]domain.Event, error)
	GetRoomEvents(eventId int64) ([]domain.Event, error)
	RescheduleRoomEvent(eventId int64, event domain.PatchEvent) error
	AddRoomEvent(event domain.Event) (int64, error)
	MyRoomEvents(userId uint64) ([]domain.Event, error)
	DeleteRoomEvent(eventId int64) error
}

type EventController struct {
	EventRepo eventrepo.Repository
}

func NewEventController(EventRepo eventrepo.Repository) *EventController {
	return &EventController{
		EventRepo: EventRepo,
	}
}

func (r EventController) GetEvent(eventId int64) (*domain.Event, error) {
	return r.EventRepo.GetEvent(eventId)
}

func (r EventController) GetEvents() ([]domain.Event, error) {
	return r.EventRepo.GetEvents()
}

func (r EventController) GetRoomEvents(eventId int64) ([]domain.Event, error) {
	return r.EventRepo.GetRoomEventsByRoomId(eventId)
}

func (r EventController) GetAuthorEvents(userId uint64) ([]domain.Event, error) {
	return r.EventRepo.GetRoomEventsByUserId(userId)
}

func (r EventController) RescheduleRoomEvent(eventId int64, event domain.PatchEvent) error {

	return r.EventRepo.RescheduleRoomEvent(eventId, event)
}

func (r EventController) AddRoomEvent(event domain.PostEvent) (int64, error) {
	fmt.Println("in post event controller")
	return r.EventRepo.AddRoomEvent(event)
}

func (r EventController) MyRoomEvents(userId uint64) ([]domain.Event, error) {
	return r.EventRepo.GetRoomEventsByUserId(userId)
}

func (r EventController) DeleteRoomEvent(eventId int64) error {
	return r.EventRepo.DeleteRoomEvent(eventId)
}
