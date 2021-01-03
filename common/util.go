package common

import (
	"io"
	"net/http"
	"os"

	"github.com/cheggaaa/pb/v3"
	"github.com/inconshreveable/go-update"
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

//SelfUpdate use binary located at newBinURL to replace self with
func SelfUpdate(newBinURL, targetPath string) error {
	resp, err := http.Get(newBinURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{
		TargetPath:  targetPath,
		OldSavePath: "",
	})
	if err != nil {
		return err
	}
	return err

}
