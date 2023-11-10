package main

import (
	"github.com/felipefernandesdev/api_opportunities/config"
	"github.com/felipefernandesdev/api_opportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("Main")
	//Initialize Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initializate error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()
}
