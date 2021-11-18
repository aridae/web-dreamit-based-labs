package apiserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	roomcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/room_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/tools/http_utils"
	"github.com/gorilla/mux"
)

type RoomHandler struct {
	RoomController    *roomcont.RoomController
	SessionController *sessioncont.SessionController
}

const (
	FAILURE_ALL_ROOMS = "failed to get all rooms: %s"
	FAILURE_ROOM      = "failed to get room: %s"
)

// GetAllRooms godoc
// @Summary Get all rooms
// @Description Get all rooms available in the system for booking
// @Tags room
// @Success 200 {array} apimodels.Room
// @Failure 500 {object} apimodels.MessageResponse
// @Router /rooms [get]
func (r2 RoomHandler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := r2.RoomController.GetAllRooms()
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: fmt.Sprintf(FAILURE_ALL_ROOMS, err),
		}, http.StatusInternalServerError)
		return
	}
	http_utils.SetJSONResponse(w, rooms, http.StatusOK)
}

// GetRoom godoc
// @Summary Get room
// @Description Get room by id
// @Tags room
// @Success 200 {array} apimodels.Room
// @Success 404 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param id path int true "Room ID"
// @Router /rooms/{id} [get]
func (r2 RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {

	roomId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_EVENT, err),
			}, http.StatusBadRequest) // 400
		return
	}

	room, err := r2.RoomController.GetRoom(int64(roomId))
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: fmt.Sprintf(FAILURE_ROOM, err),
		}, http.StatusInternalServerError)
		return
	}
	http_utils.SetJSONResponse(w, room, http.StatusOK)
}

func NewRoomHandler(RoomController *roomcont.RoomController, SessionController *sessioncont.SessionController) *RoomHandler {
	return &RoomHandler{
		RoomController:    RoomController,
		SessionController: SessionController,
	}
}
