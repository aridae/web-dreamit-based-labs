package notifyrepo

import "github.com/aridae/web-dreamit-api-based-labs/internal/domain"

type Repository interface {
	CreateNotify(notify domain.PostNotify) (int64, error)
	GetNotify(notifyId int64) (*domain.Notify, error)
	FilterNotifies(filter domain.OptionalNotifyFilter) ([]domain.Notify, error)
	DeleteNotify(notifyId int64) error
	GetNotifiesByEventId(eventId int64) ([]domain.Notify, error)
	GetEventNotifiesWithTag(eventId int64, tag string) ([]domain.Notify, error)
	GetEventNotifiesWithSubject(eventId int64, subject string) ([]domain.Notify, error)
	GetEventNotifiesWithTagSubject(eventId int64, subject, tag string) ([]domain.Notify, error)
}
