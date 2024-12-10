// main.go
//
// The main entrypoint for the toy microservice server.
// This file sets up the HTTP server, routes, and logging.
//
// Initially, it only provides a /healthz endpoint for basic health checks.
// Future steps will add echo, info endpoints, environment variable handling,
// semantic versioning integration, and more.

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/paulcapestany/toy-service/internal/handlers"
)

func main() {
	// Set up structured logging with zerolog
	// Adjust time format for readability
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Starting toy-service server")

	r := chi.NewRouter()

	// Register routes
	r.Get("/healthz", handlers.HealthzHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Info().Msgf("Listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed")
		}
	}()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Received shutdown signal")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Graceful shutdown failed")
	} else {
		log.Info().Msg("Server gracefully stopped")
	}
}
