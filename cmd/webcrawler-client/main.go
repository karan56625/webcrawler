package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"webCrawler/config"
)

func main() {
	// Define command-line flags
	urlFlag := flag.String("url", "", "The URL to crawl")
	flag.Parse()

	if *urlFlag == "" {
		log.Fatal("URL flag is required")
	}

	webcrawlerHost, exists := os.LookupEnv("WEBCRAWLER_HOST")
	if !exists {
		fmt.Println("WEBCRAWLER_HOST is not set")
		fmt.Println("Using local host as webcrawler host")
		webcrawlerHost = "localhost"
	}
	webcrawlerPort, exists := os.LookupEnv("WEBCRAWLER_PORT")
	if !exists {
		fmt.Println("WEBCRAWLER_PORT is not set")
		fmt.Println("Using default webcrawler port ", config.Port)
		webcrawlerPort = config.Port
	}

	// Define the web crawler service endpoint
	endpoint := "http://" + webcrawlerHost + ":" + webcrawlerPort + config.WebCrawlerEndpoint

	// Create the request to start crawling
	resp, err := http.Get(endpoint + "?url=" + *urlFlag)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	fmt.Println("Crawl Result:")
	fmt.Println(string(body))
}
