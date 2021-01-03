package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/mholt/archiver/v3"
	"golang.org/x/sys/windows/registry"
)

func main() {

	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Firaxis\Civilization5`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	regCivPath, _, err := k.GetStringValue("LastKnownPath")
	if err != nil {
		log.Fatal(err)
	}

	serverPtr := flag.Bool("server", false, "Starts executable in server mode")
	civPathPtr := flag.String("dir", regCivPath, "Specifiy nonstandard civ folder (default is: "+regCivPath+")")
	flag.Parse()

	fmt.Println("NQLauncher by MrThinger - Version: " + strconv.Itoa(CLIBuildNumber))

	if *serverPtr {
		fmt.Println("Server Mode Enabled")
		StartServer()
	}

	//CLI Start
	//self update
	latestBuildInfo := getSelfUpdateInfo(DefaultSelfUpdateEndpoint)
	if latestBuildInfo.Version > CLIBuildNumber {
		fmt.Println("Update detected! Updating...")
		exPath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		SelfUpdate(latestBuildInfo.DownloadURL, exPath)
		fmt.Println("Update finished! Restarting!")
		exec.Command("cmd", "/C", "start", exPath).Start()
		os.Exit(0)
	}

	//Detect Civ folder (check flag then go to default )

	civBasePath := filepath.Clean(*civPathPtr)
	civDlcPath := filepath.Join(civBasePath, "Assets", "DLC")
	civMapsPath := filepath.Join(civBasePath, "Assets", "Maps")

	validateFolderExists(civDlcPath)
	validateFolderExists(civMapsPath)
	fmt.Println("Valid Civ5 path found: " + civBasePath)

	//Delete the old files if they exist (using a regex)
	deleteFiles(filepath.Join(civMapsPath, "[Ll][Ee][Kk]*"))
	deleteFiles(filepath.Join(civDlcPath, "[Ll][Ee][Kk]*"))

	//get new file links from server (and parse them)
	links := getFileLinks(DefaultCurrentLinksEndpoint)

	//download them to cache/temp folder
	tempDir := os.TempDir()
	tempMod := filepath.Join(tempDir, "civ5-mod.zip")
	tempMap := filepath.Join(tempDir, "civ5-map.zip")
	fmt.Println("Downloading: " + links.Mod)
	DownloadFile(tempMod, links.Mod)
	fmt.Println("Downloading: " + links.Map)
	DownloadFile(tempMap, links.Map)

	//unzip them to civ folder
	fmt.Println("Unzipping mod to: " + civDlcPath)
	archiver.Unarchive(tempMod, civDlcPath)
	fmt.Println("Unzipping map to: " + civMapsPath)
	archiver.Unarchive(tempMap, civMapsPath)

}

//https://civ5-nq-launcher.herokuapp.com/currentLinks
func getFileLinks(currentLinksEndpoint string) (links CurrentLinks) {

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

func deleteFiles(glob string) {
	if matches, err := filepath.Glob(glob); err == nil {
		for _, match := range matches {
			fmt.Println("Deleting: " + match)
			os.RemoveAll(match)
		}
	}
}

func validateFolderExists(path string) {
	if val, err := os.Stat(path); os.IsNotExist(err) {
		panic("Civ5 folder not found @ " + path)
	} else if !val.IsDir() {
		panic("Found file where expected folder, why the heck is a file @ " + path)
	}
}
