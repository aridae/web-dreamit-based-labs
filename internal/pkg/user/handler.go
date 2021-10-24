package user

import (
	"net/http"
)

type Handler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
	LogInKeycloak(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	GetSelfProfile(w http.ResponseWriter, r *http.Request)
	DeleteSelfProfile(w http.ResponseWriter, r *http.Request)
}
