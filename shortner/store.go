package shortner

import (
	"context"
	"sync"

	"github.com/ckshitij/collections/heap"
)

type URLStore struct {
	shortnerMap map[string]string
	domainCount map[string]int
	topDomains  [2][]DomainCount
	activeInd   int8
	sync.RWMutex
}

func NewURLShortnerStore() *URLStore {
	return &URLStore{
		shortnerMap: make(map[string]string),
		domainCount: make(map[string]int),
		topDomains: [2][]DomainCount{
			make([]DomainCount, 1),
			make([]DomainCount, 1),
		},
		activeInd: 0,
	}
}

func (u *URLStore) Insert(ctx context.Context, data URLShortData) {
	u.Lock()
	defer u.Unlock()
	u.shortnerMap[data.ShortURL] = data.URL
	u.domainCount[data.Domain]++

	// create a new fresh heap
	priorityQueue := heap.NewHeap(func(a, b DomainCount) bool {
		return a.Count > b.Count
	})

	for domain, count := range u.domainCount {
		priorityQueue.Push(DomainCount{Domain: domain, Count: count})
	}
	temp := []DomainCount{}
	for i := 0; i < 3 && priorityQueue.Size() > 0; i++ {
		temp = append(temp, priorityQueue.Pop())
	}
	u.topDomains[u.activeInd^1] = temp
	u.activeInd ^= 1
}

func (u *URLStore) GetURL(ctx context.Context, shortURL string) (string, error) {
	u.RLock()
	defer u.RUnlock()
	url, ok := u.shortnerMap[shortURL]
	if !ok {
		return "", ErrURLNotFound
	}
	return url, nil
}

func (u *URLStore) TopDomains(ctx context.Context) *MetricsData {
	return &MetricsData{
		Metrics: u.topDomains[u.activeInd],
	}
}
