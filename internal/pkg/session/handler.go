package session

import (
	"net/http"
)

type Handler interface {
	RefreshSession(w http.ResponseWriter, r *http.Request)
	CheckSession(w http.ResponseWriter, r *http.Request)
}
