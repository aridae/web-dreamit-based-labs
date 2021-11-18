package controllers

import (
	notifyrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/notify_repo"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

type NotifyUseCase interface {
	GetNotify(notifyId int64) (*domain.Notify, error)
	CreateNotify(domain.Notify) (int64, error)
	DeleteNotify(notifyId int64) error

	FilterNotifies(filter domain.OptionalNotifyFilter) ([]domain.Notify, error)
	GetEventNotifies(eventId int64) ([]domain.Notify, error)

	GetSubjectEventNotifies(eventId int64, subject string) ([]domain.Notify, error)
	GettagEventNotifies(eventId int64, tag string) ([]domain.Notify, error)
}

type NotifyController struct {
	NotifyRepo notifyrepo.Repository
}

func (c *NotifyController) GetNotify(notifyId int64) (*domain.Notify, error) {
	return c.NotifyRepo.GetNotify(notifyId)
}

func (c *NotifyController) CreateNotify(notify domain.PostNotify) (int64, error) {
	return c.NotifyRepo.CreateNotify(notify)
}

func (c *NotifyController) DeleteNotify(notifyId int64) error {
	return c.NotifyRepo.DeleteNotify(notifyId)
}

func (c *NotifyController) GetEventNotifies(eventId int64) ([]domain.Notify, error) {
	return c.NotifyRepo.GetNotifiesByEventId(eventId)
}

func (c *NotifyController) GetSubjectEventNotifies(eventId int64, subject string) ([]domain.Notify, error) {
	return c.NotifyRepo.GetEventNotifiesWithSubject(eventId, subject)
}

func (c *NotifyController) GetTagEventNotifies(eventId int64, tag string) ([]domain.Notify, error) {
	return c.NotifyRepo.GetEventNotifiesWithTag(eventId, tag)
}

func (c *NotifyController) FilterNotifies(filter domain.OptionalNotifyFilter) ([]domain.Notify, error) {
	return c.NotifyRepo.FilterNotifies(filter)
}

func NewNotifyController(NotifyRepo notifyrepo.Repository) *NotifyController {
	return &NotifyController{
		NotifyRepo: NotifyRepo,
	}
}
