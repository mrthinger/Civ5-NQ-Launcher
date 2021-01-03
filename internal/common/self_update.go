package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/inconshreveable/go-update"
)

// SelfUpdate updates self
func SelfUpdate() {

	//self update
	latestBuildInfo := getSelfUpdateInfo(DefaultSelfUpdateEndpoint)
	if latestBuildInfo.Version > CLIBuildNumber {
		fmt.Println("Update detected! Updating...")
		exPath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		applyUpdate(latestBuildInfo.DownloadURL, exPath)
		fmt.Println("Update finished! Restarting!")
		exec.Command("cmd", "/C", "start", exPath).Start()
		os.Exit(0)
	}
}

func getSelfUpdateInfo(latestBuildEndpoint string) (latestBuild BuildInfo) {

	// Make HTTP request
	res, err := http.Get(latestBuildEndpoint)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &latestBuild)
	if err != nil {
		panic(err.Error())
	}

	return
}

//applyUpdate use binary located at newBinURL to replace self with
func applyUpdate(newBinURL, targetPath string) error {
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
