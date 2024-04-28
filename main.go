package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title       string
	Description string
}

func main() {
	res, err := http.Get("https://www.scrapethissite.com/pages/")
	if err != nil {
		log.Fatal("Failed to connect to the target page", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("HTTP Error %d: %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Failed to parse the HTML document", err)
	}

	var articles []Article

	doc.Find(".page").Each(func(i int, p *goquery.Selection) {
		article := Article{}
		article.Title = strings.TrimSpace(p.Find(".page-title").Text())
		article.Description = strings.TrimSpace(p.Find(".session-desc").Text())

		articles = append(articles, article)
	})

	json, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(json))
}
