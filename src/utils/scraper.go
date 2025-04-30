package utils

import (
	"encoding/json"
	"log"
	"net/url"
	"os"

	"github.com/gocolly/colly"
)

type Product struct {
	Name       	string
	Badge       string
	Image    	string
	Rating      string
	Price 		string
}

func Scraper(rawurl string) {
	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal("Failed to parse URL:", err)
	}

	domain := parsedURL.Host

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"),
	)

	var products []Product

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	c.OnHTML("div.s-result-item", func(e *colly.HTMLElement) {
		product := Product{
			Name:      e.ChildText("h2.a-size-medium span"),
			Badge:      e.ChildText("span.a-badge-text"),
			Image:   e.ChildAttr("img.s-image", "src"),
			Rating:     e.ChildText("i.a-icon-star-small span.a-icon-alt"),
			Price: 		e.ChildText("span.a-price > span.a-offscreen"),
		}

		// Only append if title exists (to avoid empty divs)
		if product.Name != "" {
			products = append(products, product)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("Scraping completed for", r.Request.URL)

		file, err := os.Create("products.json")
		if err != nil {
			log.Fatal("Could not create JSON file:", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(products); err != nil {
			log.Fatal("Could not write JSON data:", err)
		}
		log.Println("Saved data to products.json")
	})

	c.Visit(rawurl)
}
