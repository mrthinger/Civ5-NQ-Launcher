package main

import (
	"github.com/gin-gonic/gin"
)

//StartServer starts backend server
func StartServer() {
	r := gin.Default()
	port := GetEnv("port", "8080")
	r.GET("/currentLek", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"currentLink": "https://",
		})
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
