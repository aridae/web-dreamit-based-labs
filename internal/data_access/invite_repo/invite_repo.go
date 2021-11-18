package inviterepo

import (
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

// репозитории принимают, и возващают модель контроллера?
// контроллер вообще ничего не знает о модели репозитория
// но! репозитории поттягивают данные из бд в модель репозитория
// а потом мапят на модель котроллера
type Repository interface {
	CreateInvite(invite domain.PostInvite) (int64, error)
	DeleteInviteById(inviteId int64) error
	UpdateInviteStatusById(inviteId int64, status int64) error
	GetInviteById(inviteId int64) (*domain.Invite, error)
	GetInvitesByReceiverId(recId uint64) ([]domain.Invite, error)
	GetInvitesByEventId(eventId int64) ([]domain.Invite, error)
	GetStatusInvitesByEventId(eventId int64, statusId int64) ([]domain.Invite, error)
	UpdateInvitesStatusByEventId(eventId int64, status int64) error // одной транзакцией
	DeleteInvitesByEventId(eventId int64) error                     // одной транзакцией
}
