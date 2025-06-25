package shortener

import (
	"encoding/json"
	"log"
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
		log.Printf("Invalid request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Shortening URL: %s", req.Url)

	code, err := h.svc.ShortenURL(req.Url)
	if err != nil {
		log.Printf("Failed to shorten URL %s: %v", req.Url, err)
		http.Error(w, "Failed to generate short URL.", http.StatusInternalServerError)
		return
	}

	shortUrl := r.Host + "/" + code
	log.Printf("Generated short URL: %s -> %s", req.Url, shortUrl)

	response := gen.ShortenResponse{ShortUrl: &shortUrl}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GET /{code}
func (h *Handler) RedirectUrl(w http.ResponseWriter, r *http.Request, code string) {
	log.Printf("Redirect requested for code: %s", code)

	url, err := h.svc.ResolveURL(code)
	if err != nil {
		log.Printf("Code not found: %s", code)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	log.Printf("Redirecting to original URL: %s", url)
	http.Redirect(w, r, url, http.StatusFound)
}

// GET /metrics
func (h *Handler) Metrics(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching top 3 shortened domains")

	topDomains := h.svc.GetTopDomains(3)

	var resp []gen.DomainCount
	for _, d := range topDomains {
		log.Printf("Domain: %s, Count: %d", d.Domain, d.Count)
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
