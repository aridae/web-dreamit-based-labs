package room

import "net/http"

type Handler interface {
	GetAllRooms(w http.ResponseWriter, r *http.Request)
	GetRoomCalendar(w http.ResponseWriter, r *http.Request)
	GetRoomSchedule(w http.ResponseWriter, r *http.Request)
	UpdateRoomBooking(w http.ResponseWriter, r *http.Request)
	AddRoomBooking(w http.ResponseWriter, r *http.Request)
	DeleteRoomBooking(w http.ResponseWriter, r *http.Request)
	MyRoomBooking(w http.ResponseWriter, r *http.Request)
}
