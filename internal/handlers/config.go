// config.go
//
// Exposes a safe, non-secret summary of runtime configuration for internal verification.
// Specifically, it indicates whether FAKE_SECRET is present without revealing its value.

package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

type ConfigSummary struct {
	FakeSecretPresent bool `json:"fakeSecretPresent"`
	FakeSecretLen     int  `json:"fakeSecretLen"`
}

// ConfigHandler handles GET /internal/config requests.
// It returns a JSON object indicating whether FAKE_SECRET is set and its length.
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	v := os.Getenv("FAKE_SECRET")
	present := v != ""
	if present {
		log.Debug().Int("fakeSecretLen", len(v)).Msg("FAKE_SECRET present")
	} else {
		log.Debug().Msg("FAKE_SECRET not set")
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ConfigSummary{
		FakeSecretPresent: present,
		FakeSecretLen:     len(v),
	}); err != nil {
		log.Error().Err(err).Msg("Failed to write /internal/config response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
