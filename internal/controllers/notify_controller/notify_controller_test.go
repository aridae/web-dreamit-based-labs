package controllers

import (
	"reflect"
	"testing"

	notifyrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/notify_repo"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

func TestNotifyController_GetNotify(t *testing.T) {
	type fields struct {
		NotifyRepo notifyrepo.Repository
	}
	type args struct {
		notifyId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Notify
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NotifyController{
				NotifyRepo: tt.fields.NotifyRepo,
			}
			got, err := c.GetNotify(tt.args.notifyId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotifyController.GetNotify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotifyController.GetNotify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifyController_CreateNotify(t *testing.T) {
	type fields struct {
		NotifyRepo notifyrepo.Repository
	}
	type args struct {
		notify domain.PostNotify
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
			c := &NotifyController{
				NotifyRepo: tt.fields.NotifyRepo,
			}
			got, err := c.CreateNotify(tt.args.notify)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotifyController.CreateNotify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NotifyController.CreateNotify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifyController_DeleteNotify(t *testing.T) {
	type fields struct {
		NotifyRepo notifyrepo.Repository
	}
	type args struct {
		notifyId int64
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
			c := &NotifyController{
				NotifyRepo: tt.fields.NotifyRepo,
			}
			if err := c.DeleteNotify(tt.args.notifyId); (err != nil) != tt.wantErr {
				t.Errorf("NotifyController.DeleteNotify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotifyController_GetEventNotifies(t *testing.T) {
	type fields struct {
		NotifyRepo notifyrepo.Repository
	}
	type args struct {
		eventId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Notify
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NotifyController{
				NotifyRepo: tt.fields.NotifyRepo,
			}
			got, err := c.GetEventNotifies(tt.args.eventId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotifyController.GetEventNotifies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotifyController.GetEventNotifies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifyController_GetSubjectEventNotifies(t *testing.T) {
	type fields struct {
		NotifyRepo notifyrepo.Repository
	}
	type args struct {
		eventId int64
		subject string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Notify
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NotifyController{
				NotifyRepo: tt.fields.NotifyRepo,
			}
			got, err := c.GetSubjectEventNotifies(tt.args.eventId, tt.args.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotifyController.GetSubjectEventNotifies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotifyController.GetSubjectEventNotifies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifyController_GetTagEventNotifies(t *testing.T) {
	type fields struct {
		NotifyRepo notifyrepo.Repository
	}
	type args struct {
		eventId int64
		tag     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Notify
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NotifyController{
				NotifyRepo: tt.fields.NotifyRepo,
			}
			got, err := c.GetTagEventNotifies(tt.args.eventId, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotifyController.GetTagEventNotifies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotifyController.GetTagEventNotifies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifyController_FilterNotifies(t *testing.T) {
	type fields struct {
		NotifyRepo notifyrepo.Repository
	}
	type args struct {
		filter domain.OptionalNotifyFilter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Notify
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NotifyController{
				NotifyRepo: tt.fields.NotifyRepo,
			}
			got, err := c.FilterNotifies(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotifyController.FilterNotifies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotifyController.FilterNotifies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNotifyController(t *testing.T) {
	type args struct {
		NotifyRepo notifyrepo.Repository
	}
	tests := []struct {
		name string
		args args
		want *NotifyController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotifyController(tt.args.NotifyRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotifyController() = %v, want %v", got, tt.want)
			}
		})
	}
}
