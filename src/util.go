package main

import (
	"io"
	"net/http"
	"os"
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

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

//GetEnv get env variable with backup
func GetEnv(env, backup string) string {

	if v, exists := os.LookupEnv(env); exists {
		return v
	}

	return backup

}
