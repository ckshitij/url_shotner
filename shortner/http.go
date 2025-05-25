package shortner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.io/ckshitij/url-shortner/handlers"
)

type URLShortner struct {
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

	fmt.Println(r)
	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	var response URLShortnerResponse
	response.ShortURL = request.URL

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
