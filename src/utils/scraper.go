package utils

import (
	"io"
	"log"
	"net/http"
)

func Scraper (url string) {
	resp, err :=http.Get(url);
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body);
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body);
	log.Printf(sb)
}