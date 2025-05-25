package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func (r *MuxRouter) AddSubRoute(path string) {
	subRouter := chi.NewRouter()
	r.Router.Mount(path, subRouter)
}

func (r *MuxRouter) RegisterResourceHandlers(basePath string, resourceHandlers []handlers.ResourseHandlers) {
	r.Router.Route(basePath, func(group chi.Router) {
		for _, resourseHdlr := range resourceHandlers {
			for _, handler := range resourseHdlr.ResourceHTTPHandlers() {
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
