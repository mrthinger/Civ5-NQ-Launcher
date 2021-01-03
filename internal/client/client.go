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
	"golang.org/x/sys/windows/registry"
)

//StartClient starts client code
func StartClient() {

	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Firaxis\Civilization5`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	regCivPath, _, err := k.GetStringValue("LastKnownPath")
	if err != nil {
		log.Fatal(err)
	}

	civPathPtr := flag.String("dir", regCivPath, "Specifiy nonstandard civ folder (default is: "+regCivPath+")")
	flag.Parse()

	fmt.Println("NQLauncher by MrThinger - Version: " + strconv.Itoa(common.CLIBuildNumber))

	//Check for updates
	common.SelfUpdate()

	//Detect Civ folder (check flag then go to default )

	civBasePath := filepath.Clean(*civPathPtr)
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

}
