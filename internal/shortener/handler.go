package shortener

import (
	"encoding/json"
	"net/http"

	"github.com/harshit-0802/url-shortener/gen"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// POST /shorten
func (h *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var req gen.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Url == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	code, err := h.svc.ShortenURL(req.Url)
	if err != nil {
		http.Error(w, "Failed to generate short URL.", http.StatusInternalServerError)
		return
	}

	shortUrl := r.Host + "/" + code
	response := gen.ShortenResponse{ShortUrl: &shortUrl}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GET /{code}
func (h *Handler) RedirectUrl(w http.ResponseWriter, r *http.Request, code string) {
	http.Redirect(w, r, "", http.StatusFound)
}
