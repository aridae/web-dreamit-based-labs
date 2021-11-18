package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	usercont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/user_controller"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/tools/http_utils"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserController    *usercont.UserController
	SessionController *sessioncont.SessionController
}

const (
	FAILURE_USERS  = "failed to get all users: %s"
	FAILURE_USER   = "failed to get user: %s"
	FAILURE_SIGNUP = "failed to signup user: %s"
	FAILURE_LOGIN  = "failed to login user: %s"
)

func NewUserHandler(UserController *usercont.UserController, SessionController *sessioncont.SessionController) *UserHandler {
	return &UserHandler{
		UserController:    UserController,
		SessionController: SessionController,
	}
}

// GetUser godoc
// @Summary Get user
// @Description Get user by id
// @Tags user
// @Accept  json
// @Produce  json
// @Succes 200 {object} apimodels.userProfile
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Failure 404 {object} apimodels.MessageResponse
// @Param id path int true "User ID"
// @Security ApiKeyAuth
// @Param Authorization header string false "token with the bearer started"
// @Router /users/{id} [get]
func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http_utils.SetJSONResponse(w,
			apimodels.MessageResponse{
				Message: fmt.Sprintf(FAILURE_DELETE_EVENT, err),
			}, http.StatusBadRequest) // 400
		return
	}

	userProfile, err := u.UserController.GetSelfProfile(uint64(userId))
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: fmt.Sprintf(FAILURE_USER, err),
		}, http.StatusNotFound) // 404
		return
	}

	http_utils.SetJSONResponse(w, userProfile, http.StatusOK)
}

// GetUsers godoc
// @Summary Get users collection
// @Description Get users collection
// @Tags user
// @Accept  json
// @Produce  json
// @Succes 200 {array} apimodels.UserProfile
// @Failure 500 {object} apimodels.MessageResponse
// @Failure 401 {object} apimodels.MessageResponse
// @Security ApiKeyAuth
// @Param Authorization header string false "token with the bearer started"
// @Router /users [get]
func (u *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserController.GetUsers()
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: fmt.Sprintf(FAILURE_USERS, err),
		}, http.StatusInternalServerError)
		return
	}
	http_utils.SetJSONResponse(w, users, http.StatusOK)
}

// SignUp godoc
// @Summary Signing user up
// @Description Signing user up by adding him to the database
// @Tags user
// @Accept  json
// @Produce  json
// @Succes 200 {object} apimodels.Token
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 409 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param request body apimodels.SignupUserRequest true "User sign up data"
// @Router /users/signup [post]
func (u *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	signupUser := &apimodels.SignupUserRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: "Invalid json provided",
		}, http.StatusBadRequest) // 400
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &signupUser)
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: "Invalid json provided",
		}, http.StatusBadRequest) // 400
		return
	}

	userId, err := u.UserController.SignUp(&domain.SignupUserData{
		Login:    signupUser.Login,
		Email:    signupUser.Email,
		Password: signupUser.Password,
	})
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: "Conflicting credetntials",
		}, http.StatusConflict) // 409
		return
	}

	token, err := u.SessionController.CreateNewSession(uint64(userId))
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: "can't create token",
		}, http.StatusInternalServerError) // 500
		return
	}

	http_utils.SetJSONResponse(w, token, http.StatusOK)
}

// LogIn godoc
// @Summary Logging user in
// @Description Logging user in by retrieving his data from the database
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} apimodels.Token
// @Failure 400 {object} apimodels.MessageResponse
// @Failure 403 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Param request body apimodels.LoginUserRequest true "User log in data"
// @Router /users/login [post]
func (u *UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	loginUser := &apimodels.LoginUserRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: "Invalid json provided",
		}, http.StatusBadRequest) // 400
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &loginUser)
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: "Invalid json provided",
		}, http.StatusBadRequest) // 400
		return
	}

	userId, err := u.UserController.LogIn(&domain.LoginUserData{
		EmailOrLogin: loginUser.EmailOrLogin,
		Password:     loginUser.Password,
	})
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: "Invalid login data",
		}, http.StatusForbidden) // 403
		return
	}

	token, err := u.SessionController.CreateNewSession(uint64(userId))
	if err != nil {
		http_utils.SetJSONResponse(w, apimodels.MessageResponse{
			Message: "can't create token",
		}, http.StatusInternalServerError) // 500
		return
	}

	http_utils.SetJSONResponse(w, token, http.StatusOK)
}

// Logout godoc
// @Summary Logging user out
// @Description Logging user in out
// @Tags user
// @Produce  json
// @Success 200 {object} apimodels.MessageResponse
// @Failure 500 {object} apimodels.MessageResponse
// @Router /users/logout [post]
func (u *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http_utils.SetJSONResponse(w, apimodels.MessageResponse{
		Message: "OK",
	}, http.StatusOK)
}

// как это аннотировать я не поняла..
func (u *UserHandler) LogInKeycloak(w http.ResponseWriter, r *http.Request) {
	index := strings.Index(r.RequestURI, "code")
	code := string([]rune(r.RequestURI)[index+5:])
	userId, err := u.UserController.LogInKeycloak(code)

	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusUnauthorized)
		return
	}
	http_utils.SetJSONResponse(w, userId, http.StatusOK)
}
