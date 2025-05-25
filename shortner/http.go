package shortner

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.io/ckshitij/url-shortner/handlers"
)

type URLShortner struct {
	service ShortnerService
}

func NewURLShortner(service ShortnerService) URLShortner {
	return URLShortner{
		service: service,
	}
}

func (h URLShortner) ResourceHTTPHandlers() []*handlers.HTTPHandler {
	return []*handlers.HTTPHandler{
		{
			Method:      http.MethodPost,
			Path:        "/url-shorten",
			MiddleWares: nil,
			Handler:     h.ShortURL,
		},
		{
			Method:      http.MethodGet,
			Path:        "/{shorten_str}",
			MiddleWares: nil,
			Handler:     h.FetchURL,
		},
	}
}

func (h *URLShortner) ShortURL(w http.ResponseWriter, r *http.Request) {
	var request URLShortnerData

	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	var response URLShortnerResponse

	data, err := h.service.ProcessURL(r.Context(), request.URL)
	if err != nil {
		HandleError(err, w)
		return
	}

	response.ShortURL = data

	// Respond with success
	handlers.WriteJSON(w, http.StatusCreated, response)
}

func (h *URLShortner) FetchURL(w http.ResponseWriter, r *http.Request) {

	// Decode JSON body
	shortenURL := chi.URLParam(r, "shorten_str")

	url, err := h.service.GetURL(r.Context(), shortenURL)
	if err != nil {
		HandleError(err, w)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

/*
curl -X POST http://localhost:8080/api/v1/url-shorten \
-H "Content-Type: application/json" \
-d '{
"url": "https://example.com/some/very/long/path",
}'
*/
