package shortner

import "time"

type URLShortData struct {
	ShortURL  string
	URL       string
	Domain    string
	CreatedAt time.Time
}

type DomainCount struct {
	Domain string `json:"domain"`
	Count  int    `json:"freuency"`
}

type URLShortnerData struct {
	URL string `json:"url"`
}

type URLShortnerResponse struct {
	ShortURL string `json:"short_url"`
}

type MetricsData struct {
	Metrics []DomainCount `json:"top3_domains"`
}
