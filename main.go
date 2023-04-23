package main

import (
	"firstproject/API/config"
	"firstproject/API/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return // exit the program
	}

	//initialize router
	router.Initialize()
}
