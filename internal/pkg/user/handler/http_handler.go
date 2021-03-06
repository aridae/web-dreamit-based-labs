package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aridae/web-dreamit-api-based-labs/internal/server/tools/http_utils"

	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/session"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/user"
)

type UserHandler struct {
	userUCase user.UseCase
	sessUCase session.UseCase
}

func NewHandler(userUCase user.UseCase, sessionUCase session.UseCase) user.Handler {
	return &UserHandler{
		userUCase: userUCase,
		sessUCase: sessionUCase,
	}
}

func (u *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	signupUser := &api_models.SignupUserRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w, "Invalid json provided", http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &signupUser)
	if err != nil {
		http_utils.SetJSONResponse(w, "Invalid json provided", http.StatusBadRequest)
		return
	}

	userId, err := u.userUCase.SignUp(signupUser)
	if err != nil {
		http_utils.SetJSONResponse(w, "Conflict", http.StatusConflict)
		return
	}

	token, err := u.sessUCase.CreateNewSession(userId)
	if err != nil {
		http_utils.SetJSONResponse(w, "Can't create token", http.StatusInternalServerError)
		return
	}

	http_utils.SetJSONResponse(w, token, http.StatusOK)

}

func (u *UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	loginUser := &api_models.LoginUserRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w, "Invalid json provided", http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &loginUser)
	if err != nil {
		http_utils.SetJSONResponse(w, "Invalid json provided", http.StatusBadRequest)
		return
	}

	userId, err := u.userUCase.LogIn(loginUser)
	if err != nil {
		http_utils.SetJSONResponse(w, "Invalid login data", http.StatusUnauthorized)
		return
	}

	token, err := u.sessUCase.CreateNewSession(userId)
	if err != nil {
		http_utils.SetJSONResponse(w, "Can't create token", http.StatusInternalServerError)
		return
	}

	http_utils.SetJSONResponse(w, token, http.StatusOK)
}

func (u *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	Uuid := r.URL.Query().Get("Uuid")

	err := u.sessUCase.DestroySession(Uuid)
	if err != nil {
		http_utils.SetJSONResponse(w, "can't remove session", http.StatusBadRequest)
		return
	}

	http_utils.SetJSONResponse(w, "OK", http.StatusOK)
}

// ?????? ?????? ???????????????????????? ?? ???? ????????????..
func (u *UserHandler) LogInKeycloak(w http.ResponseWriter, r *http.Request) {
	index := strings.Index(r.RequestURI, "code")
	code := string([]rune(r.RequestURI)[index+5:])
	userId, err := u.userUCase.LogInKeycloak(code)

	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusUnauthorized)
		return
	}

	token, err := u.sessUCase.CreateNewSession(userId)
	if err != nil {
		http_utils.SetJSONResponse(w, err, http.StatusInternalServerError)
		return
	}
	http_utils.SetJSONResponse(w, token, http.StatusOK)
}
