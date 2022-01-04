package apiserver

import (
	"net/http"
	"reflect"
	"testing"

	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	usercont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/user_controller"
)

func TestNewUserHandler(t *testing.T) {
	type args struct {
		UserController    *usercont.UserController
		SessionController *sessioncont.SessionController
	}
	tests := []struct {
		name string
		args args
		want *UserHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.UserController, tt.args.SessionController); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_GetUser(t *testing.T) {
	type fields struct {
		UserController    *usercont.UserController
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
			u := &UserHandler{
				UserController:    tt.fields.UserController,
				SessionController: tt.fields.SessionController,
			}
			u.GetUser(tt.args.w, tt.args.r)
		})
	}
}

func TestUserHandler_GetUsers(t *testing.T) {
	type fields struct {
		UserController    *usercont.UserController
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
			u := &UserHandler{
				UserController:    tt.fields.UserController,
				SessionController: tt.fields.SessionController,
			}
			u.GetUsers(tt.args.w, tt.args.r)
		})
	}
}

func TestUserHandler_SignUp(t *testing.T) {
	type fields struct {
		UserController    *usercont.UserController
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
			u := &UserHandler{
				UserController:    tt.fields.UserController,
				SessionController: tt.fields.SessionController,
			}
			u.SignUp(tt.args.w, tt.args.r)
		})
	}
}

func TestUserHandler_LogIn(t *testing.T) {
	type fields struct {
		UserController    *usercont.UserController
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
			u := &UserHandler{
				UserController:    tt.fields.UserController,
				SessionController: tt.fields.SessionController,
			}
			u.LogIn(tt.args.w, tt.args.r)
		})
	}
}

func TestUserHandler_Logout(t *testing.T) {
	type fields struct {
		UserController    *usercont.UserController
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
			u := &UserHandler{
				UserController:    tt.fields.UserController,
				SessionController: tt.fields.SessionController,
			}
			u.Logout(tt.args.w, tt.args.r)
		})
	}
}

func TestUserHandler_LogInKeycloak(t *testing.T) {
	type fields struct {
		UserController    *usercont.UserController
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
			u := &UserHandler{
				UserController:    tt.fields.UserController,
				SessionController: tt.fields.SessionController,
			}
			u.LogInKeycloak(tt.args.w, tt.args.r)
		})
	}
}
