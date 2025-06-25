package shortener

import (
	"errors"
)

type Store interface {
	Save(code, url string) error
	Load(code string) (string, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) ShortenURL(url string) (string, error) {
	// Generate code and save mapping
	code := "abc123" // TODO: Use a generator here
	err := s.store.Save(code, url)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (s *Service) ResolveURL(code string) (string, error) {
	url, err := s.store.Load(code)
	if err != nil {
		return "", errors.New("not found")
	}
	return url, nil
}
