package shortner

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleError(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Invalid URL Error",
			err:            ErrInvalidURL,
			expectedStatus: http.StatusBadRequest,
			expectedError:  ErrInvalidURL.Error(),
		},
		{
			name:           "URL Not Found Error",
			err:            ErrURLNotFound,
			expectedStatus: http.StatusNotFound,
			expectedError:  ErrURLNotFound.Error(),
		},
		{
			name:           "Unknown Error",
			err:            errors.New("some unexpected error"),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "some unexpected error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			HandleError(tt.err, rr)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response map[string]string
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("Failed to decode JSON response: %v", err)
			}

			if response["error"] != tt.expectedError {
				t.Errorf("Expected error message '%s', got '%s'", tt.expectedError, response["error"])
			}
		})
	}
}
