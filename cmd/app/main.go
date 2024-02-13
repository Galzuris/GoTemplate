package main

import (
	"backend/config"
	"backend/internal/app"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Read config error: %s", err)
	}

	app.Run(cfg)
}
