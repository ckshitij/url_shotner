package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.io/ckshitij/url-shortner/handlers"
)

// Mock implementation of ResourseHandlers
type mockResource struct{}

func (m *mockResource) ResourceHTTPHandlers() []*handlers.HTTPHandler {
	return []*handlers.HTTPHandler{
		{
			Method:  http.MethodGet,
			Path:    "/ping",
			Handler: func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) },
		},
		{
			Method:  http.MethodPost,
			Path:    "/echo",
			Handler: func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("echo")) },
		},
	}
}

func TestNewMuxRouter(t *testing.T) {
	router := NewMuxRouter()
	if router == nil || router.Router == nil {
		t.Fatal("Expected non-nil MuxRouter with Router initialized")
	}
}

func TestSetDefaultMiddlewares(t *testing.T) {
	router := NewMuxRouter()
	router.SetDefaultMiddlewares()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	router.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}
}

func TestEnableCorsConfig(t *testing.T) {
	router := NewMuxRouter()
	router.EnableCorsConfig()

	// Register a dummy OPTIONS route so CORS middleware is triggered
	router.Router.Options("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Simulate a CORS preflight request
	req, _ := http.NewRequest(http.MethodOptions, "/", nil)
	req.Header.Set("Origin", "http://example.com")
	req.Header.Set("Access-Control-Request-Method", "POST")

	rr := httptest.NewRecorder()
	router.Router.ServeHTTP(rr, req)

	allowOrigin := rr.Header().Get("Access-Control-Allow-Origin")
	if allowOrigin != "*" {
		t.Errorf("Expected 'Access-Control-Allow-Origin' to be '*', got '%s'", allowOrigin)
	}

	// Optional: test allowed methods
	allowMethods := rr.Header().Get("Access-Control-Allow-Methods")
	if allowMethods == "" {
		t.Error("Expected 'Access-Control-Allow-Methods' to be set")
	}
}

func TestRegisterResourceHandlers(t *testing.T) {
	router := NewMuxRouter()
	resource := &mockResource{}

	router.RegisterResourceHandlers("/api", []handlers.ResourseHandlers{resource})

	tests := []struct {
		method string
		path   string
		want   string
	}{
		{http.MethodGet, "/api/ping", "pong"},
		{http.MethodPost, "/api/echo", "echo"},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest(tt.method, tt.path, strings.NewReader(""))
		rr := httptest.NewRecorder()
		router.Router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status 200 for %s %s, got %d", tt.method, tt.path, rr.Code)
		}
		if strings.TrimSpace(rr.Body.String()) != tt.want {
			t.Errorf("Expected body '%s', got '%s'", tt.want, rr.Body.String())
		}
	}
}
