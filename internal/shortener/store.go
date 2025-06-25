package shortener

import (
	"errors"
	"sync"
)

type InMemoryStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string]string),
	}
}

func (s *InMemoryStore) Save(code, url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[code] = url
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
