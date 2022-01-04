package apiserver

import (
	"net/http"
	"reflect"
	"testing"

	notifycont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/notify_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
)

func TestNotifyHandler_GetNotifies(t *testing.T) {
	type fields struct {
		NotifyController  *notifycont.NotifyController
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
			handler := NotifyHandler{
				NotifyController:  tt.fields.NotifyController,
				SessionController: tt.fields.SessionController,
			}
			handler.GetNotifies(tt.args.w, tt.args.r)
		})
	}
}

func TestNotifyHandler_GetNotify(t *testing.T) {
	type fields struct {
		NotifyController  *notifycont.NotifyController
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
			handler := NotifyHandler{
				NotifyController:  tt.fields.NotifyController,
				SessionController: tt.fields.SessionController,
			}
			handler.GetNotify(tt.args.w, tt.args.r)
		})
	}
}

func TestNotifyHandler_AddNotify(t *testing.T) {
	type fields struct {
		NotifyController  *notifycont.NotifyController
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
			handler := NotifyHandler{
				NotifyController:  tt.fields.NotifyController,
				SessionController: tt.fields.SessionController,
			}
			handler.AddNotify(tt.args.w, tt.args.r)
		})
	}
}

func TestNotifyHandler_DeleteNotify(t *testing.T) {
	type fields struct {
		NotifyController  *notifycont.NotifyController
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
			handler := NotifyHandler{
				NotifyController:  tt.fields.NotifyController,
				SessionController: tt.fields.SessionController,
			}
			handler.DeleteNotify(tt.args.w, tt.args.r)
		})
	}
}

func TestNewNotifyHandler(t *testing.T) {
	type args struct {
		NotifyController  *notifycont.NotifyController
		SessionController *sessioncont.SessionController
	}
	tests := []struct {
		name string
		args args
		want *NotifyHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotifyHandler(tt.args.NotifyController, tt.args.SessionController); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotifyHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
