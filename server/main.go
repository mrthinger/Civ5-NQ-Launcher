package main

import (
	"fmt"
	"strconv"

	"github.com/mrthinger/Civ5-NQ-Launcher/common"
)

func main() {

	fmt.Println("NQLauncher by MrThinger - Version: " + strconv.Itoa(common.CLIBuildNumber))

	fmt.Println("Server Mode Enabled")
	StartServer()

}
