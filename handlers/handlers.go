package handlers

import (
	"encoding/json"
	"net/http"
)

type HTTPHandler struct {
	Method      string
	Path        string
	MiddleWares []func(http.Handler) http.Handler
	Handler     http.HandlerFunc
}

type ResourseHandlers interface {
	ResourceHTTPHandlers() []*HTTPHandler
}

// writeJSON writes any response data as JSON with the given status code.
func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeError sends an error response with a message and appropriate status code.
func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{
		"error": message,
	})
}
