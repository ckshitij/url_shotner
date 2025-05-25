package shortner

import "time"

type URLShortData struct {
	ShortURL  string
	URL       string
	Domain    string
	CreatedAt time.Time
}

type DomainCount struct {
	Domain string
	Count  int64
}

type URLShortnerRequest struct {
	URL string `json:"url"`
}

type URLShortnerResponse struct {
	ShortURL string `json:"short_url"`
}
