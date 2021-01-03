package main

import (
	"fmt"
	"strconv"

	"github.com/mrthinger/Civ5-NQ-Launcher/internal/common"
	"github.com/mrthinger/Civ5-NQ-Launcher/internal/server"
)

func main() {

	fmt.Println("NQLauncher by MrThinger - Version: " + strconv.Itoa(common.CLIBuildNumber))

	fmt.Println("Server Mode Enabled")
	server.StartServer()

}
