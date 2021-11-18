package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	notifycont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/notify_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/tools/http_utils"
	"github.com/gorilla/mux"
)

type NotifyHandler struct {
	NotifyController  *notifycont.NotifyController
	SessionController *sessioncont.SessionController
}

const (
	FAILURE_NOTIFIES = "failed to get notifies: %s"

	FAILURE_DELETE_NOTIFY = "failed to delete notify: %s"
	FAILURE_POST_NOTIFY   = "failed to post notify: %s"
	FAILURE_PATCH_NOTIFY  = "failed to patch notify: %s"
	FAILURE_GET_NOTIFY    = "failed to get notify: %s"

	SUCCESS_NOTIFY_DELETED = "sucessfully deleted notify"

	TAGS_PARAM    = "tags"
	SUBJECT_PARAM = "subject"
)

// parseParameters godoc
// Extracts and validates parameters from request query.
func (handler NotifyHandler) parseParameters(r *http.Request) (*domain.OptionalNotifyFilter, error) {
	parameters := r.URL.Query()

	var filter domain.OptionalNotifyFilter

	// obligatory
	eventIds, eventIdOk := parameters[EVENT_ID_PARAM]
	if !eventIdOk || len(eventIds) != 1 {
		return nil, fmt.Errorf("invalid event ID parameter")
	}
	eventId, err := strconv.Atoi(eventIds[0])
	if err != nil {
		return nil, err
	}
	filter.EventId = int64(eventId)

	// optional
	tags, tagsOk := parameters[TAGS_PARAM]
	if tagsOk && len(tags) > 0 {
		filter.Tags = tags
	}

	subject, subOk := parameters[SUBJECT_PARAM]
	if subOk && len(subject) == 1 {
		filter.Subject = subject[0]
	}

	return &filter, nil
}

// GetNotifies godoc
// @Summary Get notifies collection
// @Description Get notifies collection filtered by query parameters
// @Produce  json
// @Param tags query []string false "Tags to filter notifies, any, optional"
// @Param subject query string false "Subject to filter notifies, single-match, optional"
// @Param eventId query int true "Event to filter notifies for, obligatory"
// @Success 200 {array} apimodels.Notify
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 404 {object} apimodels.MessageResponse
// @Security ApiKeyAuth
// @Param Authorization header string false "token with the bearer started"
// @Tags notify
// @Router /notifies [get]
func (handler NotifyHandler) GetNotifies(w http.ResponseWriter, r *http.Request) {
	filter, err := handler.parseParameters(r)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_NOTIFIES, err),
			}, http.StatusBadRequest) // 400
		return
	}

	notifies, err := handler.NotifyController.FilterNotifies(*filter)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_NOTIFIES, err),
			}, http.StatusNotFound) // 404
		return
	}
	http_utils.SetJSONResponse(w, notifies, http.StatusOK) // 200

}

// GetNotify godoc
// @Summary Get notify
// @Description Get notify by id
// @Produce  json
// @Param id path int true "Notify ID"
// @Success 200 {object} apimodels.Notify
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 404 {object} apimodels.MessageResponse
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags notify
// @Router /notifies/{id} [get]
func (handler NotifyHandler) GetNotify(w http.ResponseWriter, r *http.Request) {
	notifyId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_GET_INVITE, err),
			}, http.StatusBadRequest) // 400
		return
	}

	notify, err := handler.NotifyController.GetNotify(int64(notifyId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_GET_INVITE, err),
			}, http.StatusNotFound) // 404
		return
	}

	http_utils.SetJSONResponse(w, notify, http.StatusOK) // 200
}

// AddNotify godoc
// @Summary Create new notify
// @Description Create notify and get id
// @Produce  json
// @Accept  json
// @Param id body apimodels.PostNotify true "New notify to add to the system"
// @Success 201 {object} apimodels.SuccessPostNotify
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags notify
// @Router /notifies [post]
func (handler NotifyHandler) AddNotify(w http.ResponseWriter, r *http.Request) {
	notify := apimodels.PostNotify{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_NOTIFY, err),
			}, http.StatusBadRequest) // 400
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &notify)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_NOTIFY, err),
			}, http.StatusBadRequest) // 400
		return
	}

	notifyId, err := handler.NotifyController.CreateNotify(domain.PostNotify{
		EventId: notify.EventId,
	})
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_EVENT, err),
			}, http.StatusConflict) // 500
		return
	}

	http_utils.SetJSONResponse(w, apimodels.SuccessPostNotify{
		Id: notifyId,
	}, http.StatusCreated) // 201
}

// DeleteNotify godoc
// @Summary Delete notify by id
// @Description Delete notify by id
// @Produce  json
// @Accept  json
// @Param id path int true "Notify ID"
// @Success 200 {object} apimodels.MessageResponse
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 404 {object} apimodels.MessageResponse
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags notify
// @Router /notifies/{id} [delete]
func (handler NotifyHandler) DeleteNotify(w http.ResponseWriter, r *http.Request) {
	notifyId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_NOTIFY, err),
			}, http.StatusBadRequest) // 400
		return
	}

	err = handler.NotifyController.DeleteNotify(int64(notifyId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_NOTIFY, err),
			}, http.StatusConflict) // 409
		return
	}

	http_utils.SetJSONResponse(w,
		apimodels.MessageResponse{
			Message: SUCCESS_NOTIFY_DELETED,
		}, http.StatusOK) // 200

}

func NewNotifyHandler(NotifyController *notifycont.NotifyController, SessionController *sessioncont.SessionController) *NotifyHandler {
	return &NotifyHandler{
		NotifyController:  NotifyController,
		SessionController: SessionController,
	}
}
