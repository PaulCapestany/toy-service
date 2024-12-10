// healthz.go
//
// The healthz handler provides a simple health check endpoint.
// It returns a static JSON response {"status":"ok"} if the server is running.
// This endpoint is used for readiness/liveness probes and initial verification
// that the service is functioning properly.

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// HealthzHandler handles GET /healthz requests.
// It returns a JSON object indicating server health status.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Handling /healthz request")

	resp := map[string]string{"status": "ok"}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Error().Err(err).Msg("Failed to write /healthz response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Debug().Msg("/healthz response successfully returned")
}
