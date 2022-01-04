package apiserver

import (
	"net/http"
	"reflect"
	"testing"

	eventcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/event_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
)

func TestEventHandler_GetEventsCollection(t *testing.T) {
	type fields struct {
		EventController   *eventcont.EventController
		SessionController *sessioncont.SessionController
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := EventHandler{
				EventController:   tt.fields.EventController,
				SessionController: tt.fields.SessionController,
			}
			handler.GetEventsCollection(tt.args.w, tt.args.r)
		})
	}
}

func TestEventHandler_GetEvent(t *testing.T) {
	type fields struct {
		EventController   *eventcont.EventController
		SessionController *sessioncont.SessionController
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := EventHandler{
				EventController:   tt.fields.EventController,
				SessionController: tt.fields.SessionController,
			}
			handler.GetEvent(tt.args.w, tt.args.r)
		})
	}
}

func TestEventHandler_PostEvent(t *testing.T) {
	type fields struct {
		EventController   *eventcont.EventController
		SessionController *sessioncont.SessionController
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := EventHandler{
				EventController:   tt.fields.EventController,
				SessionController: tt.fields.SessionController,
			}
			handler.PostEvent(tt.args.w, tt.args.r)
		})
	}
}

func TestEventHandler_DeleteEvent(t *testing.T) {
	type fields struct {
		EventController   *eventcont.EventController
		SessionController *sessioncont.SessionController
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := EventHandler{
				EventController:   tt.fields.EventController,
				SessionController: tt.fields.SessionController,
			}
			handler.DeleteEvent(tt.args.w, tt.args.r)
		})
	}
}

func TestEventHandler_PatchEvent(t *testing.T) {
	type fields struct {
		EventController   *eventcont.EventController
		SessionController *sessioncont.SessionController
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := EventHandler{
				EventController:   tt.fields.EventController,
				SessionController: tt.fields.SessionController,
			}
			handler.PatchEvent(tt.args.w, tt.args.r)
		})
	}
}

func TestNewEventHandler(t *testing.T) {
	type args struct {
		EventController   *eventcont.EventController
		SessionController *sessioncont.SessionController
	}
	tests := []struct {
		name string
		args args
		want *EventHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventHandler(tt.args.EventController, tt.args.SessionController); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
