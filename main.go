package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize router using default gin configs
	r := gin.Default()

	// Define a route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
