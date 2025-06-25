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
