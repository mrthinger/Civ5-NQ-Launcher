package client

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mholt/archiver/v3"
	"github.com/mrthinger/Civ5-NQ-Launcher/internal/common"
	"github.com/sqweek/dialog"
	"golang.org/x/sys/windows/registry"
)

//StartClient starts client code
func StartClient() {
	cmdCivPath := *flag.String("dir", "", "Specifiy nonstandard civ folder")
	flag.Parse()

	var regCivPath string
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\Steam App 8930`, registry.QUERY_VALUE)
	if err != nil {
		regCivPath, _, err = k.GetStringValue("InstallLocation")
		if err != nil {
			log.Println(err)
		}
	}
	defer k.Close()

	var dialogCivPath string
	if cmdCivPath == "" && regCivPath == "" {
		dialogCivPath, err = dialog.Directory().Title("Pick Civ5 Folder (has the exe in it)").Browse()
		if err != nil {
			log.Fatal(err)
		}
	}

	var civPath string
	if cmdCivPath != "" {
		civPath = cmdCivPath
	} else if regCivPath != "" {
		civPath = regCivPath
	} else if dialogCivPath != "" {
		civPath = dialogCivPath
	} else {
		log.Fatal("No Civ5 folder found")
	}

	fmt.Println("Civ5 by MrThinger - Version: " + strconv.Itoa(common.CLIBuildNumber))

	//Check for updates
	common.SelfUpdate()

	civBasePath := filepath.Clean(civPath)
	civDlcPath := filepath.Join(civBasePath, "Assets", "DLC")
	civMapsPath := filepath.Join(civBasePath, "Assets", "Maps")

	common.ValidateFolderExists(civDlcPath)
	common.ValidateFolderExists(civMapsPath)
	fmt.Println("Valid Civ5 path found: " + civBasePath)

	//Delete the old files if they exist (using a regex)
	common.DeleteFiles(filepath.Join(civMapsPath, "[Ll][Ee][Kk]*"))
	common.DeleteFiles(filepath.Join(civDlcPath, "[Ll][Ee][Kk]*"))

	//get new file links from server (and parse them)
	links := common.GetFileLinks(common.DefaultCurrentLinksEndpoint)

	//download them to cache/temp folder
	tempDir := os.TempDir()
	tempMod := filepath.Join(tempDir, "civ5-mod.zip")
	tempMap := filepath.Join(tempDir, "civ5-map.zip")
	fmt.Println("Downloading: " + links.Mod)
	common.DownloadFile(tempMod, links.Mod)
	fmt.Println("Downloading: " + links.Map)
	common.DownloadFile(tempMap, links.Map)

	//unzip them to civ folder
	fmt.Println("Unzipping mod to: " + civDlcPath)
	archiver.Unarchive(tempMod, civDlcPath)
	fmt.Println("Unzipping map to: " + civMapsPath)
	archiver.Unarchive(tempMap, civMapsPath)

	log.Println("Success! Press Enter to exit...")
	fmt.Scanln()
}
