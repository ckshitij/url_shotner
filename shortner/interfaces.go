package shortner

import "context"

type ShortnerService interface {
	ProcessURL(ctx context.Context, url string) (string, error)
	GetURL(shortURL string) (string, error)
}

type ShortnerStore interface {
	Insert(data URLShortData)
	GetURL(shortURL string) (string, error)
}
