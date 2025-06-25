package shortener

import (
	"testing"
)

func TestInMemoryStore_SaveAndLoad(t *testing.T) {
	store := NewInMemoryStore()

	code := "abc123"
	url := "https://example.com/path"

	// Save URL
	err := store.Save(code, url)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Load URL
	gotURL, err := store.Load(code)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if gotURL != url {
		t.Errorf("expected URL %q, got %q", url, gotURL)
	}
}

func TestInMemoryStore_Load_NotFound(t *testing.T) {
	store := NewInMemoryStore()

	_, err := store.Load("nonexistent")
	if err == nil {
		t.Fatal("expected error for missing code, got nil")
	}
}

func TestGetTopDomains(t *testing.T) {
	store := NewInMemoryStore()

	// Add URLs with known domains and frequencies
	urls := []string{
		"https://youtube.com/video1",
		"https://youtube.com/video2",
		"https://youtube.com/video3",
		"https://en.wikipedia.org/article1",
		"https://wikipedia.org/article2",
		"https://example.com/page1",
	}

	for i, u := range urls {
		code := "code" + string(rune(i+'a'))
		if err := store.Save(code, u); err != nil {
			t.Fatalf("Save failed: %v", err)
		}
	}

	top := store.GetTopDomains(2)

	expected := []DomainCount{
		{Domain: "youtube.com", Count: 3},
		{Domain: "wikipedia.org", Count: 2},
	}

	if len(top) != len(expected) {
		t.Fatalf("Expected %d top domains, got %d", len(expected), len(top))
	}

	for i, domain := range expected {
		if top[i].Domain != domain.Domain || top[i].Count != domain.Count {
			t.Errorf("Mismatch at rank %d: got %+v, want %+v", i+1, top[i], domain)
		}
	}
}
