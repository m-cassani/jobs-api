package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router using default gin configs
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run the server
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
