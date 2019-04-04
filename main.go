package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExampleScrape parses data from html
func ExampleScrape(emailBody string) {
	// emailBody := `<div id="app" > <div class="wrapper"> <div class="item txt"> Text </div>	  <div class="item img"> Image </div></div></div>`
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(emailBody))
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(emailBody)

	// Find the review items
	doc.Find(".wrapper").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		key := s.Find(".txt").Text()
		value := s.Find(".img").Text()
		fmt.Printf("Row %d: %s - %s\n", i, key, value)
	})
}

func main() {
	emailBodyPtr := flag.String("email", "", "raw HTML of the email body")
	flag.Parse()
	fmt.Println(*emailBodyPtr)
	ExampleScrape(*emailBodyPtr)
}
