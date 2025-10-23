package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
	Tags   string `json:"tags"`
}

func main() {
	res, err := http.Get("http://quotes.toscrape.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]Quote, 0)

	doc.Find(".quote").Each(func(i int, sel *goquery.Selection) {
		row := Quote{
			Text:   sel.Find(".text").Text(),
			Author: sel.Find(".author").Text(),
			Tags:   sel.Find(".tags .tag").Text(),
		}
		rows = append(rows, row)
	})

	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bts))
}
