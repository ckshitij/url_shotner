package server

import (
	"fmt"
	"net/http"
	"time"

	"github.io/ckshitij/url-shortner/config"
)

func NewHTTPServer(cfg *config.ServiceConfig) *http.Server {
	router := NewMuxRouter()
	router.SetDefaultMiddlewares()

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
		Handler:      router.Router,
	}
}
