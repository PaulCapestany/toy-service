// echo.go
//
// The echo handler returns the provided message with "[modified]" appended,
// along with metadata about version, commit, and env.

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Env     string `json:"env"`
}

// EchoHandler handles POST /echo requests.
// It echoes back the input message, appending " [modified]", and returns
// version, commit, and environment info.
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Handling /echo request")

	var req EchoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error().Err(err).Msg("Failed to decode /echo request body")
		http.Error(w, `{"error":"Invalid input"}`, http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		http.Error(w, `{"error":"Invalid input"}`, http.StatusBadRequest)
		return
	}

	cfg := LoadEnvConfig()

	resp := EchoResponse{
		Message: req.Message + " [modified]",
		Version: cfg.Version,
		Commit:  cfg.GitCommit,
		Env:     cfg.Env,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Error().Err(err).Msg("Failed to write /echo response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Debug().Msg("/echo response successfully returned")
}
