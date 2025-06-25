package main

import (
	"log"
	"net/http"

	"github.com/harshit-0802/url-shortener/gen"
	"github.com/harshit-0802/url-shortener/internal/shortener"
)

func main() {
	store := shortener.NewInMemoryStore()
	urlGenerator := shortener.NewSHA1Base64Generator()
	svc := shortener.NewService(store, urlGenerator)
	handler := shortener.NewHandler(svc)
	router := gen.HandlerWithOptions(handler, gen.ChiServerOptions{})

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
