package shortner

import (
	"errors"
	"net/http"
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
	ErrInvalidURL: {
		StatusCode: http.StatusBadRequest,
	},
}
