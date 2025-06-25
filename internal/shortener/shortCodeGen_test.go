package shortener

import (
	"testing"
)

func TestSHA1Base64Generator_Generate(t *testing.T) {
	gen := NewSHA1Base64Generator()

	url1 := "https://example.com/some/path"
	url2 := "https://example.com/some/otherpath"

	code1, err := gen.Generate(url1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(code1) != 10 {
		t.Errorf("expected code length 10, got %d", len(code1))
	}

	// Calling generate on the same url should produce same output
	code1Again, err := gen.Generate(url1)
	if err != nil {
		t.Fatalf("unexpected error on repeated generate: %v", err)
	}
	if code1 != code1Again {
		t.Errorf("expected same code for same input, got %s and %s", code1, code1Again)
	}

	code2, err := gen.Generate(url2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if code1 == code2 {
		t.Errorf("expected different urls, got %s", code1)
	}
}
