package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gocolly/colly/v2"
)

type article struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Summary string `json:"summary"`
}

func main() {
	server := newDemoSite()
	defer server.Close()

	articles := map[string]*article{}
	c := colly.NewCollector()

	c.OnHTML("article.card", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.ChildText("h2"))
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		articles[link] = &article{
			Title: title,
			URL:   link,
		}
		if err := e.Request.Visit(link); err != nil {
			log.Println("visit detail:", err)
		}
	})

	c.OnHTML("div.detail", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()
		if item, ok := articles[url]; ok {
			item.Summary = strings.TrimSpace(e.ChildText("p"))
		}
	})

	if err := c.Visit(server.URL); err != nil {
		log.Fatal(err)
	}

	var result []article
	for _, item := range articles {
		result = append(result, *item)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func newDemoSite() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
<html>
  <body>
    <article class="card">
      <h2>Gin Router Basics</h2>
      <a href="/detail/gin">Read</a>
    </article>
    <article class="card">
      <h2>gRPC Hello Service</h2>
      <a href="/detail/grpc">Read</a>
    </article>
  </body>
</html>`)
	})

	mux.HandleFunc("/detail/gin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<div class="detail"><p>Gin is great for quick JSON APIs and middleware-based web services.</p></div>`)
	})

	mux.HandleFunc("/detail/grpc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<div class="detail"><p>gRPC works well for service-to-service communication with generated clients.</p></div>`)
	})

	return httptest.NewServer(mux)
}
