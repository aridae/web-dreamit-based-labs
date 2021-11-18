package room

import "net/http"

type Handler interface {
	GetAllRooms(w http.ResponseWriter, r *http.Request)
	GetRoomEvents(w http.ResponseWriter, r *http.Request)
	UpdateRoomEvent(w http.ResponseWriter, r *http.Request)
	AddRoomEvent(w http.ResponseWriter, r *http.Request)
	DeleteRoomEvent(w http.ResponseWriter, r *http.Request)
	MyRoomEvents(w http.ResponseWriter, r *http.Request)
}
