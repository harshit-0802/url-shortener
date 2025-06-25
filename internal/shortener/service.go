package shortener

import (
	"errors"
)

type Store interface {
	Save(code, url string) error
	Load(code string) (string, error)
	GetTopDomains(n int) []DomainCount
}

type Service struct {
	store     Store
	generator ShortCodeGenerator
}

func NewService(store Store, generator ShortCodeGenerator) *Service {
	return &Service{
		store:     store,
		generator: generator,
	}
}

func (s *Service) ShortenURL(url string) (string, error) {
	// Generate code and save mapping
	code, err := s.generator.Generate(url)
	if err != nil {
		return "", err
	}

	err = s.store.Save(code, url)
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

func (s *Service) GetTopDomains(n int) []DomainCount {
	return s.store.GetTopDomains(n)
}
