package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	invitecont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/invite_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/tools/http_utils"
	"github.com/gorilla/mux"
)

const (
	FAILURE_INVITES_BY_RECEIVER        = "failed to get invites for receiver: %s"
	FAILURE_INVITES_BY_EVENT           = "failed to get invites for event: %s"
	FAILURE_STATUS_INVITES_BY_EVENT    = "failed to filter event invites by status: %s"
	FAILURE_PATCH_EVENT_INVITES_STATUS = "failed to patch event invites status: %s"
	SUCCESS_PATCH_EVENT_INVITES_STATUS = "successfully patched event invites status"

	FAILURE_DELETE_INVITE = "failed to delete invite: %s"
	FAILURE_POST_INVITE   = "failed to post invite: %s"
	FAILURE_PATCH_INVITE  = "failed to patch invite: %s"
	FAILURE_GET_INVITE    = "failed to get invite: %s"

	SUCCESS_INVITE_DELETED = "invite sucessfully deleted"
	SUCCESS_INVITE_PATCHED = "invite sucessfully patched"
	FAILURE_GET_INVITES    = "resource unaccessible"
	RECEIVER_ID_PARAM      = "receiverId"
	EVENT_ID_PARAM         = "eventId"
	STATUS_ID_PARAM        = "status"
)

type InviteHandler struct {
	InviteController  *invitecont.InviteController
	SessionController *sessioncont.SessionController
}

// Getinvites godoc
// @Summary Get invites collection
// @Description Get invites collection filtered by query parameters
// @Produce  json
// @Param receiverId query int false "Receiver ID"
// @Param eventId query int false "Event ID"
// @Success 200 {array} apimodels.Invite
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Security ApiKeyAuth
// @Param Authorization header string false "token with the bearer started"
// @Tags invite
// @Router /invites [get]
func (handler InviteHandler) GetInvites(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	receiverId, okReceiverId := parameters[RECEIVER_ID_PARAM]
	eventId, okEventId := parameters[EVENT_ID_PARAM]
	statusId, okStatusId := parameters[STATUS_ID_PARAM]

	if okReceiverId && len(receiverId) == 1 {
		handler.getInvitesByReceiver(w, r)
		return
	}

	if okEventId && len(eventId) == 1 {
		if okStatusId && len(statusId) == 1 {
			handler.getStatusInvitesByEvent(w, r)
			return
		}

		handler.getInvitesByEvent(w, r)
		return
	}

	http_utils.SetJSONResponse(w,
		apimodels.MessageResponse{
			Message: FAILURE_GET_INVITES,
		}, http.StatusNotFound) // 404
}

func (handler InviteHandler) getInvitesByReceiver(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	receiverIds := parameters[RECEIVER_ID_PARAM]
	receiverIdstr, err := strconv.Atoi(receiverIds[0])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_INVITES_BY_RECEIVER, err),
			}, http.StatusBadRequest) // 400
	}

	receiverId := int64(receiverIdstr)
	invites, err := handler.InviteController.GetReceiverInvites(uint64(receiverId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_INVITES_BY_RECEIVER, err),
			}, http.StatusNotFound) // 404
	}

	inviteDTOs := make([]apimodels.Invite, len(invites))
	for _, invite := range invites {
		inviteDTOs[0] = apimodels.Invite{
			Id:         invite.Id,
			EventId:    invite.EventId,
			ReceiverId: invite.ReceiverId,
			StatusId:   invite.StatusId,
		}
	}

	http_utils.SetJSONResponse(w, inviteDTOs, http.StatusOK) // 200
}

func (handler InviteHandler) getInvitesByEvent(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	eventIds := parameters[RECEIVER_ID_PARAM]
	eventIdstr, err := strconv.Atoi(eventIds[0])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_INVITES_BY_EVENT, err),
			}, http.StatusBadRequest) // 400
	}
	statusIds := parameters[STATUS_ID_PARAM]
	statusIdstr, err := strconv.Atoi(statusIds[0])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_INVITES_BY_EVENT, err),
			}, http.StatusBadRequest) // 400
	}

	eventId := int64(eventIdstr)
	statusId := int64(statusIdstr)

	invites, err := handler.InviteController.GetStatusEventInvites(int64(eventId), int64(statusId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_INVITES_BY_EVENT, err),
			}, http.StatusInternalServerError) // 500
	}

	inviteDTOs := make([]apimodels.Invite, len(invites))
	for _, invite := range invites {
		inviteDTOs[0] = apimodels.Invite{
			Id:         invite.Id,
			EventId:    invite.EventId,
			ReceiverId: invite.ReceiverId,
			StatusId:   invite.StatusId,
		}
	}

	http_utils.SetJSONResponse(w, inviteDTOs, http.StatusOK) // 200
}

func (handler InviteHandler) getStatusInvitesByEvent(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	eventIds := parameters[RECEIVER_ID_PARAM]
	eventIdstr, err := strconv.Atoi(eventIds[0])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_INVITES_BY_EVENT, err),
			}, http.StatusBadRequest) // 400
	}

	eventId := int64(eventIdstr)
	invites, err := handler.InviteController.GetEventInvites(int64(eventId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_INVITES_BY_EVENT, err),
			}, http.StatusInternalServerError) // 500
	}

	inviteDTOs := make([]apimodels.Invite, len(invites))
	for _, invite := range invites {
		inviteDTOs[0] = apimodels.Invite{
			Id:         invite.Id,
			EventId:    invite.EventId,
			ReceiverId: invite.ReceiverId,
			StatusId:   invite.StatusId,
		}
	}

	http_utils.SetJSONResponse(w, inviteDTOs, http.StatusOK) // 200
}

// GetInvite godoc
// @Summary Get invite
// @Description Get invite by id
// @Produce  json
// @Param id path int true "Invite ID"
// @Success 200 {object} apimodels.Invite
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 404 {object} apimodels.MessageResponse
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags invite
// @Router /invites/{id} [get]
func (handler InviteHandler) GetInvite(w http.ResponseWriter, r *http.Request) {
	inviteId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_GET_INVITE, err),
			}, http.StatusBadRequest) // 400
		return
	}

	invite, err := handler.InviteController.GetInvite(int64(inviteId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_GET_INVITE, err),
			}, http.StatusNotFound) // 404
		return
	}

	http_utils.SetJSONResponse(w, invite, http.StatusOK) // 200
}

// AddInvite godoc
// @Summary Post new invite
// @Description Create invite and get id
// @Produce  json
// @Accept  json
// @Param NewInvite body apimodels.PostInvite true "New invite to add to the system"
// @Success 201 {object} apimodels.SuccessPostInvite
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags invite
// @Router /invites [post]
func (handler InviteHandler) AddInvite(w http.ResponseWriter, r *http.Request) {
	invite := apimodels.PostInvite{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_INVITE, err),
			}, http.StatusBadRequest) // 400
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &invite)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_INVITE, err),
			}, http.StatusBadRequest) // 400
		return
	}

	inviteId, err := handler.InviteController.CreateInvite(domain.PostInvite{
		EventId:    invite.EventId,
		ReceiverId: invite.ReceiverId,
	})
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_EVENT, err),
			}, http.StatusConflict) // 500
		return
	}

	http_utils.SetJSONResponse(w, apimodels.SuccessPostInvite{
		Id: inviteId,
	}, http.StatusCreated) // 201
}

// DeleteInvite godoc
// @Summary Delete invite by id
// @Description Delete invite by id
// @Produce  json
// @Accept  json
// @Param id path int true "Invite ID"
// @Success 200 {object} apimodels.MessageResponse
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags invite
// @Router /invites/{id} [delete]
func (handler InviteHandler) DeleteInvite(w http.ResponseWriter, r *http.Request) {
	inviteId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_INVITE, err),
			}, http.StatusBadRequest) // 400
		return
	}

	err = handler.InviteController.DeleteInvite(int64(inviteId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_INVITE, err),
			}, http.StatusInternalServerError) // 500
		return
	}

	http_utils.SetJSONResponse(w,
		apimodels.MessageResponse{
			Message: SUCCESS_INVITE_DELETED,
		}, http.StatusOK) // 200
}

// PatchEventInvitesStatus godoc
// @Summary Update invites status for the event
// @Description Set invites status for the event provided by query parameters
// @Produce  json
// @Accept  json
// @Success 200 {object} apimodels.MessageResponse
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param status body apimodels.PatchInvite true "Patch with status to update"
// @Param eventId query int true "Event ID"
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags invite
// @Router /invites [patch]
func (handler InviteHandler) PatchEventInvitesStatus(w http.ResponseWriter, r *http.Request) {

	// fetch event ID from query
	parameters := r.URL.Query()
	eventIds, eventsIdsOk := parameters[EVENT_ID_PARAM]

	// validate parameters
	if !eventsIdsOk || len(eventIds) != 1 {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_EVENT_INVITES_STATUS,
					fmt.Errorf("bad query event ID provided")),
			}, http.StatusBadRequest) // 400
		return
	}
	eventId, err := strconv.Atoi(eventIds[0])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_EVENT_INVITES_STATUS,
					fmt.Errorf("bad query event ID provided")),
			}, http.StatusBadRequest) // 400
		return
	}

	// fetch patch from body
	invite := apimodels.PatchInvite{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_EVENT_INVITES_STATUS, err),
			}, http.StatusBadRequest) // 400
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &invite)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_EVENT_INVITES_STATUS, err),
			}, http.StatusBadRequest) // 400
		return
	}
	if invite.StatusId != domain.EXPIRED_STATUS {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_EVENT_INVITES_STATUS,
					fmt.Errorf("action unsupported")),
			}, http.StatusForbidden) // 403
		return
	}

	// route controller method based on parameters privided
	err = handler.InviteController.ExpireEventInvites(int64(eventId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_INVITE, err),
			}, http.StatusInternalServerError) // 500
		return
	}

	http_utils.SetJSONResponse(w, apimodels.MessageResponse{
		Message: SUCCESS_PATCH_EVENT_INVITES_STATUS,
	}, http.StatusOK) // 200

}

// PatchInviteStatus godoc
// @Summary Update invite status
// @Description Set invite status by id
// @Produce  json
// @Accept  json
// @Success 200 {object} apimodels.MessageResponse
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param status body apimodels.PatchInvite true "Patch with status to update"
// @Param id path int true "Invite ID"
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags invite
// @Router /invites/{id} [patch]
func (handler InviteHandler) PatchInviteStatus(w http.ResponseWriter, r *http.Request) {

	// fetch invite ID from query
	inviteId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_INVITE, err),
			}, http.StatusBadRequest) // 400
		return
	}

	// fetch patch from body
	invite := apimodels.PatchInvite{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_INVITE, err),
			}, http.StatusBadRequest) // 400
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &invite)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_PATCH_INVITE, err),
			}, http.StatusBadRequest) // 400
		return
	}

	var controllerErr error
	switch invite.StatusId {
	case domain.EXPIRED_STATUS:
		controllerErr = handler.InviteController.ExpireInvite(int64(inviteId))
	case domain.ACCEPTED_STATUS:
		controllerErr = handler.InviteController.AcceptInvite(int64(inviteId))
	case domain.DECLINED_STATUS:
		controllerErr = handler.InviteController.DeclineInvite(int64(inviteId))
	default:
		controllerErr = fmt.Errorf("unsuppurted action")
	}

	if controllerErr != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_INVITE, err),
			}, http.StatusInternalServerError) // 500
		return
	}

	http_utils.SetJSONResponse(w, apimodels.MessageResponse{
		Message: SUCCESS_INVITE_PATCHED,
	}, http.StatusOK) // 200

}

func NewInviteHandler(InviteController *invitecont.InviteController, SessionController *sessioncont.SessionController) *InviteHandler {
	return &InviteHandler{
		InviteController:  InviteController,
		SessionController: SessionController,
	}
}
