package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	serverPtr := flag.Bool("server", false, "Starts executable in server mode")
	civPathPtr := flag.String("dir", DefaultCivDirectory, "Specifiy nonstandard civ folder (default is: "+DefaultCivDirectory+")")
	flag.Parse()

	if *serverPtr {
		fmt.Println("Server Mode Enabled")
		StartServer()
	}

	//CLI Start
	//Detect Civ folder (check flag then go to default )

	civBasePath := filepath.Clean(*civPathPtr)
	civDlcPath := filepath.Join(civBasePath, "Assets", "DLC")
	civMapsPath := filepath.Join(civBasePath, "Assets", "Maps")
	fmt.Println(civBasePath)

	validateFolderExists(civDlcPath)
	validateFolderExists(civMapsPath)

	//Delete the old files if they exist (using a regex)
	deleteFiles(filepath.Join(civMapsPath, "*[Ll]ek"))

	//get new file links from server (and parse them)

	//download them to cache/temp folder
	//unzip them to civ folder

}

//https://civ5-nq-launcher.herokuapp.com/currentLinks
func getFileLinks(endpoint string) (mapLink, modLink string) {

}

func unzip(zipPath, extractToPath) {

}

func deleteFiles(glob string) {

}

func validateFolderExists(path string) {
	if val, err := os.Stat(path); os.IsNotExist(err) {
		panic("Civ5 folder not found @ " + path)
	} else if !val.IsDir() {
		panic("Found file where expected folder, why the heck is a file @ " + path)
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
		href, exists := element.Attr("href")``
		if exists {
			link := "https://drive.google.com/u/0" + href
			DownloadFile("zipfile.zip", link)
			fmt.Println()
		}
	})
}
