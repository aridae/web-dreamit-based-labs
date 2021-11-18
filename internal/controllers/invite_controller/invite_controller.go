package controllers

import (
	inviterepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/invite_repo"
	domain "github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

type InviteUseCase interface {
	GetEventInvites(eventId int64) ([]domain.Invite, error)
	GetStatusInvites(eventId int64, statusId int64) ([]domain.Invite, error)
	ExpireEventInvites(eventId int64) error
	GetReceiverInvites(recId uint64) ([]domain.Invite, error)
	GetInvite(inviteId int64) (*domain.Invite, error)
	CreateInvite(domain.PostInvite) (int64, error)
	AcceptInvite(inviteId int64) error
	DeclineInvite(inviteId int64) error
	ExpireInvite(inviteId int64) error
}

type InviteController struct {
	InviteRepo inviterepo.Repository
}

func (c *InviteController) GetStatusEventInvites(eventId int64, statusId int64) ([]domain.Invite, error) {
	return c.InviteRepo.GetStatusInvitesByEventId(eventId, statusId)
}

func (c *InviteController) GetEventInvites(eventId int64) ([]domain.Invite, error) {
	return c.InviteRepo.GetInvitesByEventId(int64(eventId))
}

func (c *InviteController) ExpireEventInvites(eventId int64) error {
	return c.InviteRepo.UpdateInvitesStatusByEventId(int64(eventId), domain.EXPIRED_STATUS)
}

func (c *InviteController) GetReceiverInvites(recId uint64) ([]domain.Invite, error) {
	return c.InviteRepo.GetInvitesByReceiverId(recId)
}

func (c *InviteController) GetInvite(inviteId int64) (*domain.Invite, error) {
	return c.InviteRepo.GetInviteById(int64(inviteId))
}

func (c *InviteController) CreateInvite(inv domain.PostInvite) (int64, error) {
	return c.InviteRepo.CreateInvite(inv)
}

func (c *InviteController) DeleteInvite(inviteId int64) error {
	return c.InviteRepo.DeleteInviteById(inviteId)
}

func (c *InviteController) AcceptInvite(inviteId int64) error {
	return c.InviteRepo.UpdateInviteStatusById(inviteId, domain.ACCEPTED_STATUS)
}

func (c *InviteController) DeclineInvite(inviteId int64) error {
	return c.InviteRepo.UpdateInviteStatusById(inviteId, domain.DECLINED_STATUS)
}

func (c *InviteController) ExpireInvite(inviteId int64) error {
	return c.InviteRepo.UpdateInviteStatusById(inviteId, domain.EXPIRED_STATUS)
}

func NewInviteController(InviteRepo inviterepo.Repository) *InviteController {
	return &InviteController{
		InviteRepo: InviteRepo,
	}
}
