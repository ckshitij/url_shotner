package shortner

import (
	"context"
	"fmt"
	"sync"
)

type URLStore struct {
	shortnerMap map[string]string
	domainCount map[string]int
	sync.RWMutex
}

func NewURLShortnerStore() *URLStore {
	return &URLStore{
		shortnerMap: make(map[string]string),
		domainCount: make(map[string]int),
	}
}

func (u *URLStore) Insert(ctx context.Context, data URLShortData) {
	u.Lock()
	defer u.Unlock()
	u.shortnerMap[data.ShortURL] = data.URL
	u.domainCount[data.Domain]++
	fmt.Printf("%+v\n", u.shortnerMap)
}

func (u *URLStore) GetURL(ctx context.Context, shortURL string) (string, error) {
	url, ok := u.shortnerMap[shortURL]
	fmt.Printf("%+v\n", u.shortnerMap)
	if !ok {
		return "", ErrURLNotFound
	}
	return url, nil
}
