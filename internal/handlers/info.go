// info.go
//
// The info handler returns service metadata based on environment variables
// and hardcoded defaults.

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// InfoHandler handles GET /info requests.
// It returns details about the service configuration and runtime environment.
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Handling /info request")

	cfg := LoadEnvConfig()

	resp := map[string]string{
		"name":         cfg.Name,
		"version":      cfg.Version,
		"env":          cfg.Env,
		"logVerbosity": cfg.LogVerbosity,
		"fakeSecret":   cfg.FakeSecret,
		"commit":       cfg.GitCommit,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Error().Err(err).Msg("Failed to write /info response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Debug().Msg("/info response successfully returned")
}
