package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/abbos-ron2/go-medium/config"
	"github.com/abbos-ron2/go-medium/rest"
	"github.com/abbos-ron2/go-medium/storage"
	"github.com/jackc/pgx/v4"
)

func main() {
	var cfg = config.Load()
	connStr := "postgres://" + cfg.DBUser + ":" + cfg.DBPass + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable"

	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	storage := storage.New(cfg, db)
	srv := rest.New(cfg, storage)

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
