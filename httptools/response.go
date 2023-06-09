package httptools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, code int, v any) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func JsonOKResponse(w http.ResponseWriter, v any) {
	JsonResponse(w, http.StatusOK, v)
}

func ResponseUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}

func ResponseBadRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

func ResponseForbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
}

func ResponseNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func ResponseInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func CachedResponse(w http.ResponseWriter, store string, maxAge int) {
	w.Header().Set("Cache-Control", fmt.Sprintf("%s, max-age=%d", store, maxAge))
}
