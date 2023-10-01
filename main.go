package main

import (
	"github.com/m-cassani/jobs-api/config"
	"github.com/m-cassani/jobs-api/router"
)

var (
	logger *config.Logger
)

func main() {
	// Create Logger for Main
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
