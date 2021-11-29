package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	eventcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/event_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/tools/http_utils"
	"github.com/gorilla/mux"
)

const (
	FAILURE_AUTHOR_EVENTS = "failed to get author events: %s"
	FAILURE_ROOM_EVENTS   = "failed to get room events: %s"
	FAILURE_EVENTS        = "failed to get events: %s"

	FAILURE_DELETE_EVENT = "failed to delete event: %s"
	FAILURE_POST_EVENT   = "failed to post event: %s"
	FAILURE_PATCH_EVENT  = "failed to patch event: %s"
	FAILURE_GET_EVENT    = "failed to get event: %s"

	SUCCESS_EVENT_DELETED = "successfully deleted event"
	SUCCESS_EVENT_PATCHED = "successfully patched event"

	AUTHOR_ID_PARAM = "authorId"
	ROOM_ID_PARAM   = "roomId"
)

type EventHandler struct {
	EventController   *eventcont.EventController
	SessionController *sessioncont.SessionController
}

// GetEventsCollection godoc
// @Summary Get events collection
// @Description Get events collection filtered by query parameters
// @Produce  json
// @Param authorId query int false "Author ID"
// @Param roomId query int false "Room ID"
// @Success 200 {array} apimodels.Event
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Security ApiKeyAuth
// @Param Authorization header string false "token with the bearer started"
// @Tags event
// @Router /events [get]
func (handler EventHandler) GetEventsCollection(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()

	if authorIds, ok := parameters[AUTHOR_ID_PARAM]; ok && len(authorIds) == 1 {
		handler.getAuthorEventsCollection(w, r)
		return
	}

	if roomIds, ok := parameters[ROOM_ID_PARAM]; ok && len(roomIds) == 1 {
		handler.getRoomEventsCollection(w, r)
		return
	}

	handler.getEventsCollection(w, r)
}

func (handler EventHandler) getEventsCollection(w http.ResponseWriter, r *http.Request) {
	events, err := handler.EventController.GetEvents()
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_EVENTS, err),
			}, http.StatusInternalServerError) // 500
		return
	}

	http_utils.SetJSONResponse(w, events, http.StatusOK) // 200
}

func (handler EventHandler) getAuthorEventsCollection(w http.ResponseWriter, r *http.Request) {
	authorId, err := strconv.Atoi(r.URL.Query()[AUTHOR_ID_PARAM][0])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_AUTHOR_EVENTS, err),
			}, http.StatusBadRequest) // 400
		return
	}

	// иначе берем коллекцию
	events, err := handler.EventController.GetAuthorEvents(uint64(authorId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_AUTHOR_EVENTS, err),
			}, http.StatusBadRequest) // 400
		return
	}

	http_utils.SetJSONResponse(w, events, http.StatusOK) // 200
}

func (handler EventHandler) getRoomEventsCollection(w http.ResponseWriter, r *http.Request) {
	roomId, err := strconv.Atoi(r.URL.Query()[ROOM_ID_PARAM][0])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_ROOM_EVENTS, err),
			}, http.StatusBadRequest) // 400
		return
	}

	// иначе берем коллекцию
	events, err := handler.EventController.GetRoomEvents(int64(roomId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_ROOM_EVENTS, err),
			}, http.StatusBadRequest) // 400
		return
	}

	http_utils.SetJSONResponse(w, events, http.StatusOK) // 200
}

// GetEvent godoc
// @Summary Get event
// @Description Get event by id
// @Produce  json
// @Param id path int true "Event ID"
// @Success 200 {object} apimodels.Event
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags event
// @Router /events/{id} [get]
func (handler EventHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_GET_EVENT, err),
			}, http.StatusBadRequest) // 400
		return
	}

	event, err := handler.EventController.GetEvent(int64(eventId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_GET_EVENT, err),
			}, http.StatusNotFound) // 404
		return
	}

	http_utils.SetJSONResponse(w, event, http.StatusOK) // 200
}

// PostEvent godoc
// @Summary Post event
// @Description Create event and get id
// @Produce  json
// @Accept  json
// @Param NewEvent body apimodels.PostEvent true "New event to add to the system"
// @Success 201 {object} apimodels.SuccessPostEvent
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 409 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags event
// @Router /events [post]
func (handler EventHandler) PostEvent(w http.ResponseWriter, r *http.Request) {

	fmt.Println("in post event")
	details, err := handler.SessionController.ExtractAccessTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_EVENT, err),
			}, http.StatusUnauthorized)
		return
	}
	userId, err := handler.SessionController.GetUserIdByAccessToken(details.Uuid)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_AUTHOR_EVENTS, err),
			}, http.StatusInternalServerError) // 500
		return
	}

	event := apimodels.PostEvent{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_EVENT, err),
			}, http.StatusBadRequest) // 400
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &event)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_EVENT, err),
			}, http.StatusBadRequest) // 400
		return
	}

	eventId, err := handler.EventController.AddRoomEvent(domain.PostEvent{
		RoomId:   event.RoomId,
		Title:    event.Title,
		Start:    event.Start,
		End:      event.End,
		AuthorId: userId,
	})
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_EVENT, err),
			}, http.StatusConflict) // 409
		return
	}

	http_utils.SetJSONResponse(w, apimodels.SuccessPostEvent{
		Id: eventId,
	}, http.StatusCreated)
}

// DeleteEvent godoc
// @Summary Delete event
// @Description Delete event by id
// @Param id path int true "Event ID"
// @Success 200 {object} apimodels.MessageResponse
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Security ApiKeyAuth
// @Param Authorization header string false "token with the bearer started"
// @Tags event
// @Router /events/{id} [delete]
func (handler EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_EVENT, err),
			}, http.StatusBadRequest) // 400
		return
	}

	err = handler.EventController.DeleteRoomEvent(int64(eventId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_EVENT, err),
			}, http.StatusConflict) // 409
		return
	}

	http_utils.SetJSONResponse(w,
		apimodels.MessageResponse{
			Message: SUCCESS_EVENT_DELETED,
		}, http.StatusOK) // 200
}

// PatchEvent godoc
// @Summary Reschedule event
// @Description Patch event roomId, start and end datetime by event id
// @Produce  json
// @Accept  json
// @Param id path int true "Event ID"
// @Param request body apimodels.PatchEvent true "Patch editions to apply"
// @Success 200 {object} apimodels.MessageResponse
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 409 {object} apimodels.MessageResponse
// @Security ApiKeyAuth
// @Param Authorization header string false "token with the bearer started"
// @Tags event
// @Router /events/{id} [patch]
func (handler EventHandler) PatchEvent(w http.ResponseWriter, r *http.Request) {

	eventId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_EVENT, err),
			}, http.StatusBadRequest)
		return
	}

	event := apimodels.PatchEvent{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_EVENT, err),
			}, http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &event)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_EVENT, err),
			}, http.StatusBadRequest)
		return
	}

	err = handler.EventController.RescheduleRoomEvent(int64(eventId), domain.PatchEvent{
		Id:     int64(eventId),
		Start:  event.Start,
		End:    event.End,
		RoomId: event.RoomId,
	})
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_EVENT, err),
			}, http.StatusConflict) // 409
		return
	}

	http_utils.SetJSONResponse(w,
		apimodels.MessageResponse{
			Message: SUCCESS_EVENT_PATCHED,
		}, http.StatusOK) // 200
}

func NewEventHandler(EventController *eventcont.EventController, SessionController *sessioncont.SessionController) *EventHandler {
	return &EventHandler{
		EventController:   EventController,
		SessionController: SessionController,
	}
}
