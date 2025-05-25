package shortner

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.io/ckshitij/url-shortner/config"
)

type serviceImpl struct {
	store ShortnerStore
}

func NewShortnerService(store ShortnerStore) ShortnerService {
	return &serviceImpl{
		store: store,
	}
}

func (s *serviceImpl) ProcessURL(ctx context.Context, url string) (string, error) {
	domain, err := s.validateAndExtractDomain(url)
	if err != nil {
		return "", err
	}

	shortURL := ShortenURL(url)

	data := URLShortData{
		ShortURL:  shortURL,
		URL:       url,
		Domain:    domain,
		CreatedAt: time.Now(),
	}

	s.store.Insert(data)
	shortenURL := fmt.Sprintf("http://%s:%s/%s", config.Config.Server.Host, config.Config.Server.Port, shortURL)
	return shortenURL, nil
}

func (s *serviceImpl) GetURL(shortURL string) (string, error) {
	return s.store.GetURL(shortURL)
}

func (s *serviceImpl) validateAndExtractDomain(rawURL string) (string, error) {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return "", ErrInvalidURL
	}

	// Must contain scheme and host to be valid
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", ErrInvalidURL
	}

	// Optional: normalize www-prefixed domains
	host := strings.ToLower(parsedURL.Host)
	host = strings.TrimPrefix(host, "www.")

	return host, nil
}
