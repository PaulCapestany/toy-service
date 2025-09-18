// version.go
//
// Exposes a lightweight handler for retrieving build/version metadata.

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// VersionHandler handles GET /version requests.
// It returns the service name, semantic version, and git commit hash for quick checks.
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Handling /version request")

	cfg := LoadEnvConfig()

	resp := map[string]string{
		"name":    cfg.Name,
		"version": cfg.Version,
		"commit":  cfg.GitCommit,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Error().Err(err).Msg("Failed to write /version response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Debug().Msg("/version response successfully returned")
}

