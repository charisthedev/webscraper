package utils

import "net/http"

func Scraper (url string) {
	http.Get(url);
}