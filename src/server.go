package main

import (
	"github.com/gin-gonic/gin"
)

const (
	//DefaultPort for server if "PORT" env variable is not specified
	DefaultPort = "8080"
)

//StartServer starts backend server
func StartServer() {
	r := gin.Default()
	port := GetEnv("PORT", DefaultPort)

	r.GET("/currentLinks", func(c *gin.Context) {
		c.JSON(200, CurrentLinks{
			Mod: "https://storage.googleapis.com/civ5-mods/lek-mod/LEKMOD_V25.1.zip",
			Map: "https://storage.googleapis.com/civ5-mods/lek-map/Lekmap%20v1.4.2.zip",
		})
	})

	r.GET("/selfUpdate", func(c *gin.Context) {
		c.JSON(200, BuildInfo{
			Version:     CLIBuildNumber,
			DownloadURL: "https://storage.googleapis.com/civ5-mods/nqlauncher/NQLauncher.exe",
		})
	})

	r.Run(":" + port)
}
