package httptools

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	AddRoutes(*mux.Router)
}

func HandlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
