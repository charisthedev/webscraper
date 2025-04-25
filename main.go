package main

import (
	"fmt"

	"github.com/charisthedev/go-webscraper/src/utils"
)

func main(){
	var url string;
	utils.OnGo();
	fmt.Scan(&url)
	utils.Scraper(url)
}