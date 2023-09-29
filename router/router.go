package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize router using default gin configs
	router := gin.Default()

	// Define a route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
