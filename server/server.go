package server

import (
	"fmt"
	"net/http"
	"text/template"
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

	tmpl := template.Must(template.ParseGlob("index.html"))

	router.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Name": "World",
		}

		err := tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

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
