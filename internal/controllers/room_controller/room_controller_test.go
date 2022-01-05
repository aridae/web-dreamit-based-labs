package controllers

import (
	"reflect"
	"testing"

	roomrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/room_repo"
	domain "github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

func TestNewRoomController(t *testing.T) {
	type args struct {
		RoomRepo roomrepo.Repository
	}
	tests := []struct {
		name string
		args args
		want *RoomController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoomController(tt.args.RoomRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoomController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomController_GetRoom(t *testing.T) {
	type fields struct {
		RoomRepo roomrepo.Repository
	}
	type args struct {
		roomId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Room
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RoomController{
				RoomRepo: tt.fields.RoomRepo,
			}
			got, err := r.GetRoom(tt.args.roomId)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoomController.GetRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoomController.GetRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomController_GetAllRooms(t *testing.T) {
	type fields struct {
		RoomRepo roomrepo.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Room
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RoomController{
				RoomRepo: tt.fields.RoomRepo,
			}
			got, err := r.GetAllRooms()
			if (err != nil) != tt.wantErr {
				t.Errorf("RoomController.GetAllRooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoomController.GetAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
