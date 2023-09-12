package main

import (
	"context"
	"log"
	takehome "tally-takehome/internal"
	"tally-takehome/internal/utils"
)

func main() {
	config, err := utils.LoadConfig[utils.Config]("./", "config.yaml")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	service, err := takehome.NewENSMonitoringService(ctx, config)
	if err != nil {
		log.Fatal("failed to start indexing service", "error", err)
	}

	service.Run(ctx)

	cancel()
}
