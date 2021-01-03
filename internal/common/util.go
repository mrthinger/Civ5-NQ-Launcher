package common

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
)

//DownloadFile Downloads url to filepath
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// start new bar
	bar := pb.Full.Start64(resp.ContentLength)
	// create proxy reader
	barReader := bar.NewProxyReader(resp.Body)

	// Write the body to file
	_, err = io.Copy(out, barReader)

	bar.Finish()

	return err
}

//GetEnv get env variable with backup
func GetEnv(env, backup string) string {

	if v, exists := os.LookupEnv(env); exists {
		return v
	}

	return backup

}

// GetFileLinks gets links to most up to date map & mod
// https://civ5-nq-launcher.herokuapp.com/currentLinks
func GetFileLinks(currentLinksEndpoint string) (links CurrentLinks) {

	// Make HTTP request
	res, err := http.Get(currentLinksEndpoint)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &links)
	if err != nil {
		panic(err.Error())
	}

	return
}

// DeleteFiles deletes everything matching glob including subdirs
func DeleteFiles(glob string) {
	if matches, err := filepath.Glob(glob); err == nil {
		for _, match := range matches {
			fmt.Println("Deleting: " + match)
			os.RemoveAll(match)
		}
	}
}

// ValidateFolderExists panics if no folder is found at path
func ValidateFolderExists(path string) {
	if val, err := os.Stat(path); os.IsNotExist(err) {
		panic("Folder not found @ " + path)
	} else if !val.IsDir() {
		panic("Found file where expected folder, why the heck is a file @ " + path)
	}
}
