package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	commentcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/comment_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/tools/http_utils"
	"github.com/gorilla/mux"
)

const (
	NOTIFY_ID_PARAM = "notifyId"

	FAILURE_COMMENTS = "failed to get comments: %s"

	FAILURE_DELETE_COMMENT = "failed to delete comment: %s"
	FAILURE_POST_COMMENT   = "failed to post comment: %s"
	FAILURE_PATCH_COMMENT  = "failed to patch comment: %s"
	FAILURE_GET_COMMENT    = "failed to get comment: %s"

	SUCCESS_COMMENT_DELETED = "sucessfully deleted comment"
)

type CommentHandler struct {
	CommentController *commentcont.CommentController
	SessionController *sessioncont.SessionController
}

// GetNotifyComments godoc
// @Summary Get notify comments collection
// @Description Get notify comments collection
// @Produce  json
// @Param notifyId query int true "Notify to filter comments for, obligatory"
// @Success 200 {array} apimodels.Comment
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Security ApiKeyAuth
// @Param Authorization header string false "token with the bearer started"
// @Tags comment
// @Router /comments [get]
func (handler CommentHandler) GetNotifyComments(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	notifyIds, notifyIdOk := parameters[NOTIFY_ID_PARAM]
	if !notifyIdOk || len(notifyIds) != 1 {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_COMMENTS,
					fmt.Errorf("invalid query paramete notifyId")),
			}, http.StatusBadRequest) // 400
		return
	}

	notifyId, err := strconv.Atoi(notifyIds[0])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_COMMENTS,
					fmt.Errorf("invalid query paramete notifyId")),
			}, http.StatusBadRequest) // 400
		return
	}
	comments, err := handler.CommentController.GetNotifyComments(int64(notifyId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_COMMENTS,
					fmt.Errorf("invalid query paramete notifyId")),
			}, http.StatusNotFound) // 404
		return
	}
	http_utils.SetJSONResponse(w, comments, http.StatusOK) // 200
}

// GetComment godoc
// @Summary Get comment
// @Description Get comment by id
// @Produce  json
// @Param id path int true "Comment ID"
// @Success 200 {object} apimodels.Comment
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags comment
// @Router /comments/{id} [get]
func (handler CommentHandler) GetNotifyComment(w http.ResponseWriter, r *http.Request) {
	commentId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_GET_COMMENT,
					fmt.Errorf("invalid path variable id")),
			}, http.StatusBadRequest) // 400
		return
	}

	comment, err := handler.CommentController.GetComment(int64(commentId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_GET_COMMENT,
					fmt.Errorf("invalid path variable id")),
			}, http.StatusNotFound) // 404
		return
	}

	http_utils.SetJSONResponse(w, comment, http.StatusOK) // 200

}

// AddComment godoc
// @Summary Create new comment
// @Description Create comment and get id
// @Produce  json
// @Accept  json
// @Param NewComment body apimodels.PostComment true "New comment"
// @Success 201 {object} apimodels.SuccessPostComment
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param Authorization header string false "token with the bearer started"
// @Security ApiKeyAuth
// @Tags comment
// @Router /comments [post]
func (handler CommentHandler) AddNotifyComment(w http.ResponseWriter, r *http.Request) {
	comment := apimodels.PostComment{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_COMMENT, err),
			}, http.StatusBadRequest) // 400
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &comment)
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_COMMENT, err),
			}, http.StatusBadRequest) // 400
		return
	}

	// details, err := handler.SessionController.ExtractAccessTokenMetadata(r)
	// if err != nil {
	// 	http_utils.SetJSONResponse(w,
	// 		apimodels.MessageResponse{
	// 			Message: fmt.Sprintf(FAILURE_POST_EVENT, err),
	// 		}, http.StatusUnauthorized) // 401
	// 	return
	// }
	// userId, err := handler.SessionController.GetUserIdByAccessToken(details.Uuid)
	// if err != nil {
	// 	http_utils.SetJSONResponse(w,
	// 		apimodels.MessageResponse{
	// 			Message: fmt.Sprintf(FAILURE_AUTHOR_EVENTS, err),
	// 		}, http.StatusInternalServerError) // 500
	// 	return
	// }

	commentId, err := handler.CommentController.CreateComment(domain.PostComment{
		AuthorId: 0,
		NotifyId: comment.NotifyId,
		Message:  comment.Message,
	})
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_POST_COMMENT, err),
			}, http.StatusInternalServerError) // 500
		return
	}

	http_utils.SetJSONResponse(w, apimodels.SuccessPostComment{
		Id: commentId,
	}, http.StatusCreated) // 201
}

// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment by id
// @Param id path int true "Comment ID"
// @Success 200 {object} apimodels.MessageResponse
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Security ApiKeyAuth
// @Param Authorization header string false "token with the bearer started"
// @Tags comment
// @Router /comments/{id} [delete]
func (handler CommentHandler) DeleteNotifyComment(w http.ResponseWriter, r *http.Request) {
	commentId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_COMMENT, err),
			}, http.StatusBadRequest) // 400
		return
	}

	err = handler.CommentController.DeleteComment(int64(commentId))
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_EVENT, err),
			}, http.StatusConflict) // 409
		return
	}

	http_utils.SetJSONResponse(w,
		apimodels.MessageResponse{
			Message: SUCCESS_COMMENT_DELETED,
		}, http.StatusOK) // 200
}

func NewCommentHandler(CommentController *commentcont.CommentController, SessionController *sessioncont.SessionController) *CommentHandler {
	return &CommentHandler{
		CommentController: CommentController,
		SessionController: SessionController,
	}
}
