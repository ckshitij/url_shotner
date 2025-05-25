package server

import (
	"net/http"
	"time"
)

func NewHTTPServer() *http.Server {
	router := NewMuxRouter()
	router.SetDefaultMiddlewares()

	return &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router.Router,
	}
}
