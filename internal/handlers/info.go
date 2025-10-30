// info.go
//
// The info handler returns service metadata based on environment variables
// and hardcoded defaults.

package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

// InfoHandler handles GET /info requests.
// It returns details about the service configuration and runtime environment.
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Handling /info request")

	cfg := LoadEnvConfig()

	secretVal, secretSet := os.LookupEnv("FAKE_SECRET")
	fakeSecretPresent := secretSet && secretVal != ""

	resp := struct {
		Name              string `json:"name"`
		Version           string `json:"version"`
		Env               string `json:"env"`
		LogVerbosity      string `json:"logVerbosity"`
		FakeSecretPresent bool   `json:"fakeSecretPresent"`
		FakeSecretLength  int    `json:"fakeSecretLength,omitempty"`
		Commit            string `json:"commit"`
	}{
		Name:              cfg.Name,
		Version:           cfg.Version,
		Env:               cfg.Env,
		LogVerbosity:      cfg.LogVerbosity,
		FakeSecretPresent: fakeSecretPresent,
		Commit:            cfg.GitCommit,
	}

	if fakeSecretPresent {
		resp.FakeSecretLength = len(secretVal)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Error().Err(err).Msg("Failed to write /info response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Debug().Msg("/info response successfully returned")
}
