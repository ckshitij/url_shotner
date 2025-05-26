package shortner

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestURLStore_InsertAndGetURL(t *testing.T) {
	store := NewURLShortnerStore()
	ctx := context.Background()

	data := URLShortData{
		ShortURL:  "abc123",
		URL:       "https://example.com/page",
		Domain:    "example.com",
		CreatedAt: time.Now(), // timestamp irrelevant here
	}

	store.Insert(ctx, data)

	// Verify GetURL returns correct value
	gotURL, err := store.GetURL(ctx, "abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotURL != data.URL {
		t.Errorf("expected URL %q, got %q", data.URL, gotURL)
	}

	// Verify GetURL returns error for unknown key
	_, err = store.GetURL(ctx, "notfound")
	if err != ErrURLNotFound {
		t.Errorf("expected ErrURLNotFound, got %v", err)
	}
}

func TestURLStore_TopDomains(t *testing.T) {
	store := NewURLShortnerStore()
	ctx := context.Background()

	// Insert multiple domains
	domains := []string{"a.com", "b.com", "a.com", "c.com", "a.com", "b.com"}
	for i, domain := range domains {
		data := URLShortData{
			ShortURL:  fmt.Sprintf("key%c", i),
			URL:       "https://" + domain,
			Domain:    domain,
			CreatedAt: time.Now(),
		}
		store.Insert(ctx, data)
	}

	metrics := store.TopDomains(ctx)
	if len(metrics.Metrics) == 0 {
		t.Fatal("expected at least one top domain")
	}

	// The domain "a.com" should be the top domain with highest count (3)
	topDomain := metrics.Metrics[0]
	if topDomain.Domain != "a.com" {
		t.Errorf("expected top domain 'a.com', got '%s'", topDomain.Domain)
	}
	if topDomain.Count != 3 {
		t.Errorf("expected count 3 for top domain, got %d", topDomain.Count)
	}
}
