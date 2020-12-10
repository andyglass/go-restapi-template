package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andyglass/go-restapi-template/logger"
)

var (
	version  string
	revision string
	log      = logger.Logger
)

func main() {
	srv := NewServer()

	// OS Signal handler
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-done
	log.Info("Web server graceful shutdown in progress")

	// Graceful termination DB connections
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// srv.DB.Close()
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Web server graceful shutdown has failed: %+v", err)
	}
	log.Info("Web server gracefully stopped")
}
