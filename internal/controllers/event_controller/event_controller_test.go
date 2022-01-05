package controllers

import (
	"reflect"
	"testing"

	eventrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/event_repo"
	domain "github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

func TestNewEventController(t *testing.T) {
	type args struct {
		EventRepo eventrepo.Repository
	}
	tests := []struct {
		name string
		args args
		want *EventController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventController(tt.args.EventRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventController_GetEvent(t *testing.T) {
	type fields struct {
		EventRepo eventrepo.Repository
	}
	type args struct {
		eventId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Event
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := EventController{
				EventRepo: tt.fields.EventRepo,
			}
			got, err := r.GetEvent(tt.args.eventId)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventController.GetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventController.GetEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventController_GetEvents(t *testing.T) {
	type fields struct {
		EventRepo eventrepo.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Event
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := EventController{
				EventRepo: tt.fields.EventRepo,
			}
			got, err := r.GetEvents()
			if (err != nil) != tt.wantErr {
				t.Errorf("EventController.GetEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventController.GetEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventController_GetRoomEvents(t *testing.T) {
	type fields struct {
		EventRepo eventrepo.Repository
	}
	type args struct {
		eventId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Event
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := EventController{
				EventRepo: tt.fields.EventRepo,
			}
			got, err := r.GetRoomEvents(tt.args.eventId)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventController.GetRoomEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventController.GetRoomEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventController_GetAuthorEvents(t *testing.T) {
	type fields struct {
		EventRepo eventrepo.Repository
	}
	type args struct {
		userId uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Event
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := EventController{
				EventRepo: tt.fields.EventRepo,
			}
			got, err := r.GetAuthorEvents(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventController.GetAuthorEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventController.GetAuthorEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventController_RescheduleRoomEvent(t *testing.T) {
	type fields struct {
		EventRepo eventrepo.Repository
	}
	type args struct {
		eventId int64
		event   domain.PatchEvent
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
			r := EventController{
				EventRepo: tt.fields.EventRepo,
			}
			if err := r.RescheduleRoomEvent(tt.args.eventId, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("EventController.RescheduleRoomEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventController_AddRoomEvent(t *testing.T) {
	type fields struct {
		EventRepo eventrepo.Repository
	}
	type args struct {
		event domain.PostEvent
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
			r := EventController{
				EventRepo: tt.fields.EventRepo,
			}
			got, err := r.AddRoomEvent(tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventController.AddRoomEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EventController.AddRoomEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventController_MyRoomEvents(t *testing.T) {
	type fields struct {
		EventRepo eventrepo.Repository
	}
	type args struct {
		userId uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Event
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := EventController{
				EventRepo: tt.fields.EventRepo,
			}
			got, err := r.MyRoomEvents(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventController.MyRoomEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventController.MyRoomEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventController_DeleteRoomEvent(t *testing.T) {
	type fields struct {
		EventRepo eventrepo.Repository
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
			r := EventController{
				EventRepo: tt.fields.EventRepo,
			}
			if err := r.DeleteRoomEvent(tt.args.eventId); (err != nil) != tt.wantErr {
				t.Errorf("EventController.DeleteRoomEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
