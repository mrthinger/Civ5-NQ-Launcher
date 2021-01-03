package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mrthinger/Civ5-NQ-Launcher/internal/common"
)

const (
	//DefaultPort for server if "PORT" env variable is not specified
	DefaultPort = "8080"
)

//StartServer starts backend server
func StartServer() {
	r := gin.Default()
	port := common.GetEnv("PORT", DefaultPort)

	r.GET("/currentLinks", func(c *gin.Context) {
		c.JSON(200, common.CurrentLinks{
			Mod: "https://storage.googleapis.com/civ5-mods/lek-mod/LEKMOD_V26.2.zip",
			Map: "https://storage.googleapis.com/civ5-mods/lek-map/Lekmap%20v2.2.zip",
		})
	})

	r.GET("/selfUpdate", func(c *gin.Context) {
		c.JSON(200, common.BuildInfo{
			Version:     common.CLIBuildNumber,
			DownloadURL: "https://storage.googleapis.com/civ5-mods/nqlauncher/NQLauncher.exe",
		})
	})

	r.Run(":" + port)
}
