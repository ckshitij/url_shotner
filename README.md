# URL Shortener Service

A simple URL shortener service built in Go, providing REST APIs to shorten URLs, redirect to original URLs, and view domain usage metrics. The service stores data in-memory.

---

## Features

- Shorten any valid URL and get a unique short URL.
- Redirect from short URL to the original URL.
- In-memory storage of URL mappings and domain frequency metrics.
- Metrics API returns the top 3 domains with the most shortened URLs.

---

## API Endpoints

### 1. Shorten URL

**POST** `/api/v1/url-shorten`

Request Body:

```json
{
  "url": "https://example.com/some/long/url"
}
```

Response:

```json
{
  "short_url": "http://localhost:8088/api/v1/abc123"
}
```

### 2. Redirect to Original URL

**GET** `/api/v1/{short_url}`

Response: HTTP 302 redirect to the original URL.


### 3. Get Top 3 Domains by Shortening Frequency

**GET** `/api/v1/metrics`

Response:

```json
{
  "top3_domains": [
    {"domain": "youtube.com", "freuency": 6},
    {"domain": "wikipedia.org", "freuency": 4},
    {"domain": "stackoverflow.com", "freuency": 2}
  ]
}
```

## How to Run

### Clone the repo

```sh
git clone https://github.com/ckshitij/url_shotner.git
cd url_shortner
```

### Start service

```sh
go mod tidy
go run main.go
```

### Use Docker

#### Download Image and Run locally

- Go to link https://hub.docker.com/r/ckshitij/url-shortener/tags OR Use below command

```sh
docker pull ckshitij/url-shortener:0.0.2
docker run -p 8088:8088 ckshitij/url-shortener:0.0.2
```

- Open browser and type http://localhost:8088 , now you can use the shortner.

#### Build Docker

```sh
cd url_shortner
docker build --file Dockerfile -tag url-shotner:0.0.2 .
docker run -p 8088:8088 url-shotner:0.0.2
```
