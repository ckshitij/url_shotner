package server

import (
	"fmt"
	"net/http"
	"time"

	"github.io/ckshitij/url-shortner/config"
	"github.io/ckshitij/url-shortner/handlers"
	"github.io/ckshitij/url-shortner/shortner"
)

func NewHTTPServer(cfg *config.ServiceConfig) *http.Server {
	router := NewMuxRouter()
	router.SetDefaultMiddlewares()

	router.EnableCorsConfig()
	urlShortner := shortner.NewURLShortnerModule()

	resourceHandlers := []handlers.ResourseHandlers{
		urlShortner,
	}
	router.RegisterResourceHandlers("/api/v1", resourceHandlers)

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
		Handler:      router.Router,
	}
}
