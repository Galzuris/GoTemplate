package app

import (
	"backend/config"
	"log"
)

func Run(cfg *config.Config) {
	log.Println("host: " + cfg.HTTP.Host)
	log.Println("port: " + cfg.HTTP.Port)
}
