package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"webCrawler/models"

	"golang.org/x/net/html"
)

func TestCrawlHandler_MissingURLParameter(t *testing.T) {
	req, err := http.NewRequest("GET", "/crawl", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CrawlHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := "Missing URL parameter\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCrawlHandler_InvalidURLParameter(t *testing.T) {
	req, err := http.NewRequest("GET", "/crawl?url=://invalid-url", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CrawlHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := "Invalid URL parameter\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCrawlHandler_ValidURLParameter(t *testing.T) {
	// Set up environment variables for the test
	os.Setenv("NUMBER_OF_WORKER", "2")
	os.Setenv("WORKER_QUEUE_LENGTH", "10")

	defer os.Unsetenv("NUMBER_OF_WORKER")
	defer os.Unsetenv("WORKER_QUEUE_LENGTH")

	// Set up the mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<html><body><a href="/about">About</a><a href="/contact">Contact</a></body></html>`)
	}))
	defer server.Close()

	req, err := http.NewRequest("GET", "/crawl?url="+server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CrawlHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := strings.Join([]string{
		strings.TrimPrefix(server.URL, "http://"),
		"- /",
		"\t- /about",
		"\t- /contact",
	}, "\n") + "\n\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestResolveURL(t *testing.T) {
	base, _ := url.Parse("http://example.com/path/")
	tests := []struct {
		href     string
		expected string
	}{
		{"about", "/path/about"},
		{"/about", "/about"},
		{"http://example.com/contact", "/contact"},
		{"https://anotherdomain.com", ""},
	}

	for _, test := range tests {
		result := resolveURL(test.href, base)
		if result != test.expected {
			t.Errorf("ResolveURL(%s, %s) = %s; expected %s", test.href, base.String(), result, test.expected)
		}
	}
}

func TestAddNodeToSiteMap(t *testing.T) {
	sitemap = &models.Node{}
	visited = make(map[string]bool)
	uri := "example.com/path/about"

	addNodeToSiteMap(uri)

	if !visited[uri] {
		t.Errorf("AddNodeToSiteMap did not mark URI as visited")
	}

	if len(sitemap.Children) == 0 || sitemap.Children[0].URI != "example.com" {
		t.Errorf("AddNodeToSiteMap did not correctly add URI to sitemap")
	}
}

func TestExtractLinks(t *testing.T) {
	htmlData := `<html><body><a href="about">About</a><a href="/contact">Contact</a></body></html>`
	doc, _ := html.Parse(strings.NewReader(htmlData))
	base, _ := url.Parse("http://example.com/")
	sitemap = new(models.Node)
	links := extractLinks(doc, base)
	expectedLinks := []string{"/about", "/contact"}
	if len(links) != len(expectedLinks) {
		t.Errorf("ExtractLinks returned %d links; expected %d", len(links), len(expectedLinks))
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Errorf("ExtractLinks returned unexpected link: got %s want %s", link, expectedLinks[i])
		}
	}
}

func TestFormatSitemap(t *testing.T) {
	root := &models.Node{URI: "example.com"}
	about := &models.Node{URI: "about"}
	contact := &models.Node{URI: "contact"}
	root.Children = []*models.Node{about, contact}

	result := formatSitemap(root, "example.com")
	expected := "example.com\n- /example.com\n\t- /about\n\t- /contact\n"

	if result != expected {
		t.Errorf("FormatSitemap returned unexpected result: got %s want %s", result, expected)
	}
}
