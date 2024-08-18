package main

import (
	"log"
	"net/http"
	"webCrawler/internal/handlers"
)

func main() {
	http.HandleFunc("/crawl", handlers.CrawlHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
