package main

import (

	"github.com/JhonatanRian/go-api/config"
	"github.com/JhonatanRian/go-api/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("go-api")

	err := config.Init()
	if err != nil {
		logger.Errorf("Failed to load config: %v", err)
		return
	}
	router.Initialize()
}
