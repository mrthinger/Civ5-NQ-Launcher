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
		c.JSON(200, gin.H{
			"mod": "https://storage.googleapis.com/civ5-mods/lek-mod/LEKMOD_V25.1.zip",
			"map": "https://storage.googleapis.com/civ5-mods/lek-map/Lekmap%20v1.4.2.zip",
		})
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
