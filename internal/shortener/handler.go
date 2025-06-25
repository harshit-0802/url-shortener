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
	url, err := h.svc.ResolveURL(code)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *Handler) Metrics(w http.ResponseWriter, r *http.Request) {
	// Get top 3 domains
	topDomains := h.svc.GetTopDomains(3)

	var resp []gen.DomainCount
	for _, d := range topDomains {
		resp = append(resp, gen.DomainCount{
			Domain: &d.Domain,
			Count:  &d.Count,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode metrics response", http.StatusInternalServerError)
		return
	}

}
