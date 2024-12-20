// main.go
//
// The main entrypoint for the toy microservice server.
// This file sets up the HTTP server, routes, and logging.

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/paulcapestany/toy-service/internal/handlers"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Starting toy-service server (last build took 3min)")

	r := chi.NewRouter()

	// Apply CORS middleware to allow local dev connections from toy-web
	// Verbose logging is performed on handler initialization and request
	// Just allow any origin during local dev. This can be narrowed down as needed.
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"X-Request-Id"},
		AllowCredentials: false,
		MaxAge:           300, // 5 minutes
	}))

	// Register routes
	r.Get("/healthz", handlers.HealthzHandler)
	r.Post("/echo", handlers.EchoHandler)
	r.Get("/info", handlers.InfoHandler)

	srv := startServer(r)
	gracefulShutdown(srv)
}

func startServer(r *chi.Mux) *http.Server {
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

	return srv
}

func gracefulShutdown(srv *http.Server) {
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
