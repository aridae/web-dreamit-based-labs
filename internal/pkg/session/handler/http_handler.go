package handler

import (
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/server/tools/http_utils"
	"net/http"

	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/pkg/session"
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/pkg/tools/jwt_token"
)

type SessionHandler struct {
	sessUCase session.UseCase
}

func NewHandler(sessUCase session.UseCase) session.Handler {
	return &SessionHandler{
		sessUCase: sessUCase,
	}
}

func (h *SessionHandler) RefreshSession(w http.ResponseWriter, r *http.Request) {
	tokenAuth, err := jwt_token.ExtractRefreshTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w, "unauthorized", http.StatusOK)
		return
	}

	newToken, err := h.sessUCase.RefreshSession(tokenAuth.Uuid)
	if err != nil {
		http_utils.SetJSONResponse(w, "bad token", http.StatusOK)
		return
	}

	http_utils.SetJSONResponse(w, newToken, http.StatusOK)
}

func (h *SessionHandler) CheckSession(w http.ResponseWriter, r *http.Request) {
	tokenAuth, err := jwt_token.ExtractAccessTokenMetadata(r)
	if err != nil {
		http_utils.SetJSONResponse(w, "unauthorized", http.StatusOK)
		return
	}

	_, err = h.sessUCase.GetUserIdByAccessToken(tokenAuth.Uuid)
	if err != nil {
		http_utils.SetJSONResponse(w, "bad token", http.StatusOK)
		return
	}

	http_utils.SetJSONResponse(w, "OK", http.StatusOK)
}
