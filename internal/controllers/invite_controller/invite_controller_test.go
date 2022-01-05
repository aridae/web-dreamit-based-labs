package controllers

import (
	"reflect"
	"testing"

	inviterepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/invite_repo"
	domain "github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

func TestInviteController_GetStatusEventInvites(t *testing.T) {
	type fields struct {
		InviteRepo inviterepo.Repository
	}
	type args struct {
		eventId  int64
		statusId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Invite
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &InviteController{
				InviteRepo: tt.fields.InviteRepo,
			}
			got, err := c.GetStatusEventInvites(tt.args.eventId, tt.args.statusId)
			if (err != nil) != tt.wantErr {
				t.Errorf("InviteController.GetStatusEventInvites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InviteController.GetStatusEventInvites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInviteController_GetEventInvites(t *testing.T) {
	type fields struct {
		InviteRepo inviterepo.Repository
	}
	type args struct {
		eventId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Invite
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &InviteController{
				InviteRepo: tt.fields.InviteRepo,
			}
			got, err := c.GetEventInvites(tt.args.eventId)
			if (err != nil) != tt.wantErr {
				t.Errorf("InviteController.GetEventInvites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InviteController.GetEventInvites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInviteController_ExpireEventInvites(t *testing.T) {
	type fields struct {
		InviteRepo inviterepo.Repository
	}
	type args struct {
		eventId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &InviteController{
				InviteRepo: tt.fields.InviteRepo,
			}
			if err := c.ExpireEventInvites(tt.args.eventId); (err != nil) != tt.wantErr {
				t.Errorf("InviteController.ExpireEventInvites() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInviteController_GetReceiverInvites(t *testing.T) {
	type fields struct {
		InviteRepo inviterepo.Repository
	}
	type args struct {
		recId uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Invite
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &InviteController{
				InviteRepo: tt.fields.InviteRepo,
			}
			got, err := c.GetReceiverInvites(tt.args.recId)
			if (err != nil) != tt.wantErr {
				t.Errorf("InviteController.GetReceiverInvites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InviteController.GetReceiverInvites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInviteController_GetInvite(t *testing.T) {
	type fields struct {
		InviteRepo inviterepo.Repository
	}
	type args struct {
		inviteId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Invite
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &InviteController{
				InviteRepo: tt.fields.InviteRepo,
			}
			got, err := c.GetInvite(tt.args.inviteId)
			if (err != nil) != tt.wantErr {
				t.Errorf("InviteController.GetInvite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InviteController.GetInvite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInviteController_CreateInvite(t *testing.T) {
	type fields struct {
		InviteRepo inviterepo.Repository
	}
	type args struct {
		inv domain.PostInvite
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &InviteController{
				InviteRepo: tt.fields.InviteRepo,
			}
			got, err := c.CreateInvite(tt.args.inv)
			if (err != nil) != tt.wantErr {
				t.Errorf("InviteController.CreateInvite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InviteController.CreateInvite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInviteController_DeleteInvite(t *testing.T) {
	type fields struct {
		InviteRepo inviterepo.Repository
	}
	type args struct {
		inviteId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &InviteController{
				InviteRepo: tt.fields.InviteRepo,
			}
			if err := c.DeleteInvite(tt.args.inviteId); (err != nil) != tt.wantErr {
				t.Errorf("InviteController.DeleteInvite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInviteController_AcceptInvite(t *testing.T) {
	type fields struct {
		InviteRepo inviterepo.Repository
	}
	type args struct {
		inviteId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &InviteController{
				InviteRepo: tt.fields.InviteRepo,
			}
			if err := c.AcceptInvite(tt.args.inviteId); (err != nil) != tt.wantErr {
				t.Errorf("InviteController.AcceptInvite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInviteController_DeclineInvite(t *testing.T) {
	type fields struct {
		InviteRepo inviterepo.Repository
	}
	type args struct {
		inviteId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &InviteController{
				InviteRepo: tt.fields.InviteRepo,
			}
			if err := c.DeclineInvite(tt.args.inviteId); (err != nil) != tt.wantErr {
				t.Errorf("InviteController.DeclineInvite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInviteController_ExpireInvite(t *testing.T) {
	type fields struct {
		InviteRepo inviterepo.Repository
	}
	type args struct {
		inviteId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &InviteController{
				InviteRepo: tt.fields.InviteRepo,
			}
			if err := c.ExpireInvite(tt.args.inviteId); (err != nil) != tt.wantErr {
				t.Errorf("InviteController.ExpireInvite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewInviteController(t *testing.T) {
	type args struct {
		InviteRepo inviterepo.Repository
	}
	tests := []struct {
		name string
		args args
		want *InviteController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInviteController(tt.args.InviteRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInviteController() = %v, want %v", got, tt.want)
			}
		})
	}
}
