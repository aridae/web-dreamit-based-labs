package apiserver

import (
	"net/http"
	"reflect"
	"testing"

	roomcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/room_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
)

func TestRoomHandler_GetAllRooms(t *testing.T) {
	type fields struct {
		RoomController    *roomcont.RoomController
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
			r2 := RoomHandler{
				RoomController:    tt.fields.RoomController,
				SessionController: tt.fields.SessionController,
			}
			r2.GetAllRooms(tt.args.w, tt.args.r)
		})
	}
}

func TestRoomHandler_GetRoom(t *testing.T) {
	type fields struct {
		RoomController    *roomcont.RoomController
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
			r2 := RoomHandler{
				RoomController:    tt.fields.RoomController,
				SessionController: tt.fields.SessionController,
			}
			r2.GetRoom(tt.args.w, tt.args.r)
		})
	}
}

func TestNewRoomHandler(t *testing.T) {
	type args struct {
		RoomController    *roomcont.RoomController
		SessionController *sessioncont.SessionController
	}
	tests := []struct {
		name string
		args args
		want *RoomHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoomHandler(tt.args.RoomController, tt.args.SessionController); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoomHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
