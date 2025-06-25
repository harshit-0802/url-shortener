package shortener

import (
	"errors"
	"net/url"
	"sort"
	"sync"
)

type DomainCount struct {
	Domain string
	Count  int
}

type InMemoryStore struct {
	data           map[string]string
	mu             sync.RWMutex
	domainCountMap map[string]int
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data:           make(map[string]string),
		domainCountMap: make(map[string]int),
	}
}

func (s *InMemoryStore) Save(code, url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[code] = url
	domain := extractDomain(url)
	s.domainCountMap[domain]++
	return nil
}

func (s *InMemoryStore) Load(code string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.data[code]
	if !ok {
		return "", errors.New("not found")
	}
	return url, nil
}

func (s *InMemoryStore) GetTopDomains(n int) []DomainCount {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []DomainCount
	for domain, count := range s.domainCountMap {
		result = append(result, DomainCount{Domain: domain, Count: count})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	if len(result) > n {
		return result[:n]
	}
	return result
}

func extractDomain(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "unknown"
	}
	return u.Hostname()
}
