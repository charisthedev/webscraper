package main

import (
	"fmt"

	"github.com/charisthedev/go-webscraper/src/utils"
)

func main(){
	var url string;
	fmt.Printf("Enter url to scrape:")
	fmt.Scan(&url)
	utils.Scraper(url)
}