package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/abbos-ron2/go-medium/config"
	"github.com/abbos-ron2/go-medium/rest"
)

func main() {
	var cfg = config.Load()

	// quitSignal := make(chan os.Signal, 1)
	// signal.Notify(quitSignal, os.Interrupt)
	// ctx, cancel := context.WithCancel(context.Background())

	srv := rest.New(cfg)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
