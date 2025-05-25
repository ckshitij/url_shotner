package shortner

import "context"

type ShortnerService interface {
	ProcessURL(ctx context.Context, url string) (string, error)
	GetURL(ctx context.Context, shortURL string) (string, error)
}

type ShortnerStore interface {
	Insert(ctx context.Context, data URLShortData)
	GetURL(ctx context.Context, shortURL string) (string, error)
}
