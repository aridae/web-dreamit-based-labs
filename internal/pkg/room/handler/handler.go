package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/pkg/models"
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/pkg/room"
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/pkg/session"
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/server/tools/http_utils"
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/pkg/tools/jwt_token"
	"net/http"
	"strconv"
)

type RoomHandler struct {
	RoomUCase room.UseCase
	SessionUCase session.UseCase
}

func (r2 RoomHandler) DeleteRoomBooking(w http.ResponseWriter, r *http.Request) {
	bookingId, err := strconv.Atoi(mux.Vars(r)["bookingId"])
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	details, err := jwt_token.ExtractAccessTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	userId, err := r2.SessionUCase.GetUserIdByAccessToken(details.Uuid)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	err = r2.RoomUCase.DeleteRoomBooking(userId, int64(bookingId))
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	http_utils.SetJSONResponse(w, "OK", http.StatusOK)

}

func (r2 RoomHandler) MyRoomBooking(w http.ResponseWriter, r *http.Request) {
	details, err := jwt_token.ExtractAccessTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	userId, err := r2.SessionUCase.GetUserIdByAccessToken(details.Uuid)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	result, err := r2.RoomUCase.MyRoomBooking(userId)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	http_utils.SetJSONResponse(w, result, http.StatusOK)
}

func (r2 RoomHandler) AddRoomBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	roomId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	event := models.Event{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w, "Invalid json provided", http.StatusBadRequest)
	}
	defer r.Body.Close()

	details, err := jwt_token.ExtractAccessTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}
	err = json.Unmarshal(body, &event)
	if err != nil {
		http_utils.SetJSONResponse(w, "Invalid json provided", http.StatusBadRequest)
		return
	}
	event.Author, err = r2.SessionUCase.GetUserIdByAccessToken(details.Uuid)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}
	eventId, err := r2.RoomUCase.AddRoomBooking(int64(roomId), event)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	http_utils.SetJSONResponse(w, eventId, http.StatusOK)
}

func (r2 RoomHandler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	rooms, err := r2.RoomUCase.GetAllRooms()
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}
	http_utils.SetJSONResponse(w, rooms, http.StatusOK)
}

func (r2 RoomHandler) GetRoomCalendar(w http.ResponseWriter, r *http.Request) {
	roomId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	calendar, err := r2.RoomUCase.GetRoomCalendar(int64(roomId))
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}
	http_utils.SetJSONResponse(w, calendar, http.StatusOK)
}

func (r2 RoomHandler) GetRoomSchedule(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (r2 RoomHandler) UpdateRoomBooking(w http.ResponseWriter, r *http.Request) {
	details, err := jwt_token.ExtractAccessTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}
	roomId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	event := models.Event{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w, "Invalid json provided", http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &event)
	if err != nil {
		http_utils.SetJSONResponse(w, "Invalid json provided", http.StatusBadRequest)
		return
	}

	event.Author, _ = r2.SessionUCase.GetUserIdByAccessToken(details.Uuid)
	err = r2.RoomUCase.UpdateRoomBooking(int64(roomId), event)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}

	http_utils.SetJSONResponse(w, "ok", http.StatusOK)
}

func NewHandler(UCase room.UseCase, SessionUCase session.UseCase) room.Handler {
	return &RoomHandler{
		RoomUCase: UCase,
		SessionUCase: SessionUCase,
	}
}
