package app

import (
	"backend/config"
	"backend/core/httpserver"
	"backend/core/postgres"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	// Database
	pg, err := postgres.New(cfg.Database.URL, postgres.MaxPoolSize(cfg.Database.PoolMax))
	if err != nil {
		log.Fatalln(fmt.Errorf("app:run - postgres: %w", err))
	}
	defer pg.Close()
	log.Println("app:run postgres started")

	// HTTP Server
	httpHandler := gin.New()
	httpServer := httpserver.New(httpHandler, httpserver.Addr(cfg.HTTP.Host, cfg.HTTP.Port))
	log.Println("app:run httpserver started (" + cfg.HTTP.Host + ":" + cfg.HTTP.Port + ")")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Signals
	select {
	case s := <-interrupt:
		log.Println("app:run os-signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Printf("app:run httpserver error: %s", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		log.Printf("app:run httpserver shutdown error: %s", err)
	}
}
