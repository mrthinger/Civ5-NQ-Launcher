package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	serverPtr := flag.Bool("server", false, "Starts executable in server mode")
	flag.Parse()

	if *serverPtr {
		fmt.Println("Server Mode Enabled")
		StartServer()
	}

}

func getFile() {
	// Make HTTP request
	response, err := http.Get("https://drive.google.com/uc?export=download&id=16I9i3atnDlJ3J8D8EPTQsg7p5__3luTS")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find and print image URLs
	document.Find("a#uc-download-link").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			link := "https://drive.google.com/u/0" + href
			DownloadFile("zipfile.zip", link)
			fmt.Println()
		}
	})
}
