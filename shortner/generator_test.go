package shortner

import (
	"testing"
)

func TestEncodeBase62(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected string
	}{
		{"Zero", 0, "0"},
		{"One", 1, "1"},
		{"SixtyOne", 61, "Z"},
		{"SixtyTwo", 62, "10"},
		{"LargeNumber", 3844, "100"}, // 62^2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EncodeBase62(tt.input)
			if got != tt.expected {
				t.Errorf("EncodeBase62(%d) = %s; want %s", tt.input, got, tt.expected)
			}
		})
	}
}

func TestHashURL(t *testing.T) {
	url := "https://example.com"
	hash1 := HashURL(url)
	hash2 := HashURL(url)

	if hash1 != hash2 {
		t.Errorf("HashURL should return consistent hashes, got %d and %d", hash1, hash2)
	}

	// Optional: make sure it's non-zero
	if hash1 == 0 {
		t.Error("HashURL returned 0; expected non-zero hash")
	}
}

func TestShortenURL(t *testing.T) {
	url := "https://example.com"
	short1 := ShortenURL(url)
	short2 := ShortenURL(url)

	if short1 != short2 {
		t.Errorf("ShortenURL should return consistent values; got %s and %s", short1, short2)
	}

	if short1 == "" {
		t.Error("ShortenURL returned empty string")
	}

	t.Logf("Shortened URL for %s is %s", url, short1)
}
