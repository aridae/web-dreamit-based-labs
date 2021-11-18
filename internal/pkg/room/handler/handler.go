package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/room"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/session"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/tools/http_utils"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/jwt_token"
	"github.com/gorilla/mux"
)

type RoomHandler struct {
	RoomUCase    room.UseCase
	SessionUCase session.UseCase
}

const (
	FAILURE_MY_ROOM_BOOKING    = "failed to get your rooms: %s"
	FAILURE_ALL_ROOMS          = "failed to get all rooms: %s"
	FAILURE_ROOM_EVENTS        = "failed to get room events: %s"
	FAILURE_DELETE_EVENT       = "failed to delete room event: %s"
	FAILURE_ADD_EVENT          = "failed to delete room event: %s"
	FAILURE_UPDATE_EVENT       = "failed to delete room event: %s"
	SUCCESS_ROOM_EVENT_DELETED = "successfully deleted room"
)

func (r2 RoomHandler) DeleteRoomEvent(w http.ResponseWriter, r *http.Request) {
	eventId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_EVENT, err),
			}, http.StatusBadRequest) // 400
		return
	}

	details, err := jwt_token.ExtractAccessTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_EVENT, err),
			}, http.StatusUnauthorized) // 401
		return
	}

	userId, err := r2.SessionUCase.GetUserIdByAccessToken(details.Uuid)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_EVENT, err),
			}, http.StatusConflict) // 409
		return
	}

	err = r2.RoomUCase.DeleteRoomEvent(userId, int64(eventId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_EVENT, err),
			}, http.StatusConflict)
		return
	}

	http_utils.SetJSONResponse(w,
		api_models.MessageResponse{
			Message: SUCCESS_ROOM_EVENT_DELETED,
		}, http.StatusOK)
}

func (r2 RoomHandler) MyRoomEvents(w http.ResponseWriter, r *http.Request) {
	details, err := jwt_token.ExtractAccessTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_MY_ROOM_BOOKING, err),
			}, http.StatusUnauthorized) // 401
		return
	}

	userId, err := r2.SessionUCase.GetUserIdByAccessToken(details.Uuid)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_MY_ROOM_BOOKING, err),
			}, http.StatusBadRequest) // 400
		return
	}

	events, err := r2.RoomUCase.MyRoomEvents(userId)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_MY_ROOM_BOOKING, err),
			}, http.StatusConflict) // 409
		return
	}

	http_utils.SetJSONResponse(w, events, http.StatusOK) // 200
}

func (r2 RoomHandler) AddRoomEvent(w http.ResponseWriter, r *http.Request) {

	event := api_models.Event{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_ADD_EVENT, err),
			}, http.StatusBadRequest)
	}
	defer r.Body.Close()

	details, err := jwt_token.ExtractAccessTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusConflict)
		return
	}
	err = json.Unmarshal(body, &event)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_ADD_EVENT, err),
			}, http.StatusBadRequest)
		return
	}
	event.Author, err = r2.SessionUCase.GetUserIdByAccessToken(details.Uuid)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_ADD_EVENT, err),
			}, http.StatusConflict)
		return
	}
	eventId, err := r2.RoomUCase.AddRoomEvent(event)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_ADD_EVENT, err),
			}, http.StatusConflict)
		return
	}

	http_utils.SetJSONResponse(w, eventId, http.StatusOK)
}

func (r2 RoomHandler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	rooms, err := r2.RoomUCase.GetAllRooms()
	if err != nil {
		http_utils.SetJSONResponse(w, api_models.MessageResponse{
			Message: fmt.Sprintf(FAILURE_ALL_ROOMS, err),
		}, http.StatusConflict)
		return
	}
	http_utils.SetJSONResponse(w, rooms, http.StatusOK)
}

func (r2 RoomHandler) GetRoomEvents(w http.ResponseWriter, r *http.Request) {
	roomId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_ROOM_EVENTS, err),
			}, http.StatusBadRequest)
		return
	}

	events, err := r2.RoomUCase.GetRoomEvents(int64(roomId))
	if err != nil {
		http_utils.SetJSONResponse(w, api_models.MessageResponse{
			Message: fmt.Sprintf(FAILURE_ROOM_EVENTS, err),
		}, http.StatusConflict)
		return
	}
	http_utils.SetJSONResponse(w, events, http.StatusOK)
}

func (r2 RoomHandler) UpdateRoomEvent(w http.ResponseWriter, r *http.Request) {
	details, err := jwt_token.ExtractAccessTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_UPDATE_EVENT, err),
			}, http.StatusUnauthorized)
		return
	}
	roomId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_UPDATE_EVENT, err),
			}, http.StatusConflict)
		return
	}

	event := api_models.Event{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_UPDATE_EVENT, err),
			}, http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &event)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_UPDATE_EVENT, err),
			}, http.StatusBadRequest)
		return
	}

	event.Author, _ = r2.SessionUCase.GetUserIdByAccessToken(details.Uuid)
	err = r2.RoomUCase.UpdateRoomEvent(int64(roomId), event)
	if err != nil {
		http_utils.SetJSONResponse(w,
			api_models.MessageResponse{
				Message: fmt.Sprintf(FAILURE_UPDATE_EVENT, err),
			}, http.StatusConflict)
		return
	}

	http_utils.SetJSONResponse(w, "ok", http.StatusOK)
}

func NewHandler(UCase room.UseCase, SessionUCase session.UseCase) room.Handler {
	return &RoomHandler{
		RoomUCase:    UCase,
		SessionUCase: SessionUCase,
	}
}
