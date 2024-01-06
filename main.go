package main

import (
	"github.com/VituSuperMEg/gopportuntes/config"
	"github.com/VituSuperMEg/gopportuntes/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Init Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
