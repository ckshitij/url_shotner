package shortner

import (
	"encoding/json"
	"net/http"

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
	}
}

func (h *URLShortner) ShortURL(w http.ResponseWriter, r *http.Request) {
	var request URLShortnerRequest

	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	var response URLShortnerResponse

	data, err := h.service.ProcessURL(r.Context(), request.URL)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if val, ok := errorStatusMap[err]; ok {
			statusCode = val.StatusCode
		}
		handlers.WriteError(w, statusCode, err.Error())
	}

	response.ShortURL = data

	// Respond with success
	handlers.WriteJSON(w, http.StatusCreated, response)
}

/*
curl -X POST http://localhost:8080/api/v1/url-shorten \
-H "Content-Type: application/json" \
-d '{
"url": "https://example.com/some/very/long/path",
}'
*/
