package shortner

import (
	"errors"
	"net/http"

	"github.io/ckshitij/url-shortner/handlers"
)

var (
	ErrInvalidURL  = errors.New("URL must include scheme and domain (e.g., https://example.com)")
	ErrURLNotFound = errors.New("url not found")
)

// for now just status code extend it for
// error code etc.
type ErrorHTTPData struct {
	StatusCode int
}

var errorStatusMap = map[error]ErrorHTTPData{
	ErrInvalidURL:  {StatusCode: http.StatusBadRequest},
	ErrURLNotFound: {StatusCode: http.StatusNotFound},
}

func HandleError(err error, w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError
	if val, ok := errorStatusMap[err]; ok {
		statusCode = val.StatusCode
	}
	handlers.WriteError(w, statusCode, err.Error())
}
