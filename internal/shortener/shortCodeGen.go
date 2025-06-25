package shortener

import (
	"crypto/sha1"
	"encoding/base64"
)

type ShortCodeGenerator interface {
	Generate(input string) (string, error)
}

type SHA1Base64Generator struct{}

func NewSHA1Base64Generator() *SHA1Base64Generator {
	return &SHA1Base64Generator{}
}

func (g *SHA1Base64Generator) Generate(input string) (string, error) {
	// Create SHA1 hash of the input URL
	h := sha1.New()
	_, err := h.Write([]byte(input))
	if err != nil {
		return "", err
	}

	hash := h.Sum(nil)

	// Encode hash to base64 URL-safe string
	encoded := base64.URLEncoding.EncodeToString(hash)

	return encoded[:10], nil
}
