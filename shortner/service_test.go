package shortner

import (
	"context"
	"fmt"
	"testing"

	"github.io/ckshitij/url-shortner/config"
)

// Mock store
type mockShortnerStore struct {
	insertCalled bool
	insertData   URLShortData
	urlMap       map[string]string
	topDomains   *MetricsData
}

func (m *mockShortnerStore) Insert(ctx context.Context, data URLShortData) {
	m.insertCalled = true
	m.insertData = data
	if m.urlMap == nil {
		m.urlMap = make(map[string]string)
	}
	m.urlMap[data.ShortURL] = data.URL
}

func (m *mockShortnerStore) GetURL(ctx context.Context, shortURL string) (string, error) {
	url, ok := m.urlMap[shortURL]
	if !ok {
		return "", ErrURLNotFound
	}
	return url, nil
}

func (m *mockShortnerStore) TopDomains(ctx context.Context) *MetricsData {
	return m.topDomains
}

func init() {
	config.Config = &config.ServiceConfig{
		Server: config.ServerConfig{
			Host: "localhost",
			Port: "8080",
		},
	}
}

func TestProcessURL_Success(t *testing.T) {
	store := &mockShortnerStore{}
	service := NewShortnerService(store)

	longURL := "https://example.com/page"
	shortURL, err := service.ProcessURL(context.Background(), longURL)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !store.insertCalled {
		t.Error("Expected Insert to be called")
	}

	if store.insertData.URL != longURL {
		t.Errorf("Expected stored URL to be '%s', got '%s'", longURL, store.insertData.URL)
	}

	expectedPrefix := fmt.Sprintf("http://%s:%s/api/v1/", config.Config.Server.Host, config.Config.Server.Port)
	if shortURL == "" || shortURL[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("Expected short URL to start with '%s', got '%s'", expectedPrefix, shortURL)
	}
}

func TestProcessURL_Invalid(t *testing.T) {
	store := &mockShortnerStore{}
	service := NewShortnerService(store)

	_, err := service.ProcessURL(context.Background(), "bad-url")
	if err == nil {
		t.Fatal("Expected error for invalid URL, got nil")
	}
	if err != ErrInvalidURL {
		t.Errorf("Expected ErrInvalidURL, got %v", err)
	}
}

func TestGetURL_Success(t *testing.T) {
	store := &mockShortnerStore{
		urlMap: map[string]string{
			"abc123": "https://example.com/page",
		},
	}
	service := NewShortnerService(store)

	result, err := service.GetURL(context.Background(), "abc123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result != "https://example.com/page" {
		t.Errorf("Expected result to be 'https://example.com/page', got '%s'", result)
	}
}

func TestGetURL_NotFound(t *testing.T) {
	store := &mockShortnerStore{}
	service := NewShortnerService(store)

	_, err := service.GetURL(context.Background(), "notfound")
	if err == nil {
		t.Fatal("Expected error for unknown short URL, got nil")
	}
	if err != ErrURLNotFound {
		t.Errorf("Expected ErrURLNotFound, got %v", err)
	}
}

func TestValidateAndExtractDomain(t *testing.T) {
	service := NewShortnerService(&mockShortnerStore{})

	tests := []struct {
		input       string
		expectError bool
		expected    string
	}{
		{"https://example.com", false, "example.com"},
		{"https://www.example.com", false, "example.com"},
		{"invalid-url", true, ""},
		{"ftp://host", false, "host"},
		{"http://", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			domain, err := service.(*serviceImpl).validateAndExtractDomain(tt.input)
			if tt.expectError && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !tt.expectError && domain != tt.expected {
				t.Errorf("Expected domain '%s', got '%s'", tt.expected, domain)
			}
		})
	}
}
