package handlers

import "net/http"

type HTTPHandler struct {
	Method      string
	Path        string
	MiddleWares []func(http.Handler) http.Handler
	Handler     http.HandlerFunc
}

type ResourseHandlers interface {
	ResourceHTTPHandlers() []*HTTPHandler
}
