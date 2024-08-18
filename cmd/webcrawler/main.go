package main

import (
	"log"
	"net/http"
	"webCrawler/config"
	"webCrawler/internal/handlers"
)

func main() {
	http.HandleFunc(config.WebCrawlerEndpoint, handlers.CrawlHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/readiness", readinessHandler)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	// Just returning OK status, as we are using any dependency like DB to check the readiness.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}
