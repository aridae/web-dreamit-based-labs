package apiserver

import (
	"net/http"
	"reflect"
	"testing"

	invitecont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/invite_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
)

func TestInviteHandler_GetInvites(t *testing.T) {
	type fields struct {
		InviteController  *invitecont.InviteController
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
			handler := InviteHandler{
				InviteController:  tt.fields.InviteController,
				SessionController: tt.fields.SessionController,
			}
			handler.GetInvites(tt.args.w, tt.args.r)
		})
	}
}

func TestInviteHandler_GetInvite(t *testing.T) {
	type fields struct {
		InviteController  *invitecont.InviteController
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
			handler := InviteHandler{
				InviteController:  tt.fields.InviteController,
				SessionController: tt.fields.SessionController,
			}
			handler.GetInvite(tt.args.w, tt.args.r)
		})
	}
}

func TestInviteHandler_AddInvite(t *testing.T) {
	type fields struct {
		InviteController  *invitecont.InviteController
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
			handler := InviteHandler{
				InviteController:  tt.fields.InviteController,
				SessionController: tt.fields.SessionController,
			}
			handler.AddInvite(tt.args.w, tt.args.r)
		})
	}
}

func TestInviteHandler_DeleteInvite(t *testing.T) {
	type fields struct {
		InviteController  *invitecont.InviteController
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
			handler := InviteHandler{
				InviteController:  tt.fields.InviteController,
				SessionController: tt.fields.SessionController,
			}
			handler.DeleteInvite(tt.args.w, tt.args.r)
		})
	}
}

func TestInviteHandler_PatchEventInvitesStatus(t *testing.T) {
	type fields struct {
		InviteController  *invitecont.InviteController
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
			handler := InviteHandler{
				InviteController:  tt.fields.InviteController,
				SessionController: tt.fields.SessionController,
			}
			handler.PatchEventInvitesStatus(tt.args.w, tt.args.r)
		})
	}
}

func TestInviteHandler_PatchInviteStatus(t *testing.T) {
	type fields struct {
		InviteController  *invitecont.InviteController
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
			handler := InviteHandler{
				InviteController:  tt.fields.InviteController,
				SessionController: tt.fields.SessionController,
			}
			handler.PatchInviteStatus(tt.args.w, tt.args.r)
		})
	}
}

func TestNewInviteHandler(t *testing.T) {
	type args struct {
		InviteController  *invitecont.InviteController
		SessionController *sessioncont.SessionController
	}
	tests := []struct {
		name string
		args args
		want *InviteHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInviteHandler(tt.args.InviteController, tt.args.SessionController); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInviteHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
