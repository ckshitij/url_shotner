package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.io/ckshitij/url-shortner/handlers"
)

type MuxRouter struct {
	Router *chi.Mux
}

func NewMuxRouter() *MuxRouter {
	router := chi.NewRouter()

	return &MuxRouter{
		Router: router,
	}
}

func (r *MuxRouter) SetDefaultMiddlewares() {
	middlewares := []func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60 * time.Second),
	}
	r.Router.Use(middlewares...)
}

func (r *MuxRouter) EnableCorsConfig() {
	r.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Use "*" to allow all origins, or restrict to ["http://localhost:3000"]
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any major browsers
	}))
}

func (r *MuxRouter) RegisterResourceHandlers(basePath string, resourceHandlers []handlers.ResourseHandlers) {
	r.Router.Route(basePath, func(group chi.Router) {
		for _, resourseHdlr := range resourceHandlers {
			for _, handler := range resourseHdlr.ResourceHTTPHandlers() {
				log.Printf("Registed endpoint : %s %s%s", handler.Method, basePath, handler.Path)
				switch handler.Method {
				case http.MethodGet:
					group.With(handler.MiddleWares...).Get(handler.Path, handler.Handler)
				case http.MethodPost:
					group.With(handler.MiddleWares...).Post(handler.Path, handler.Handler)
				case http.MethodPut:
					group.With(handler.MiddleWares...).Put(handler.Path, handler.Handler)
				case http.MethodPatch:
					group.With(handler.MiddleWares...).Patch(handler.Path, handler.Handler)
				case http.MethodDelete:
					group.With(handler.MiddleWares...).Delete(handler.Path, handler.Handler)
				case http.MethodOptions:
					group.With(handler.MiddleWares...).Options(handler.Path, handler.Handler)
				}
			}
		}
	})
}
