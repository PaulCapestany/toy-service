// echo.go
//
// The echo handler returns the provided message with "[modified]" appended,
// along with metadata about version, commit, and env.

package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
)

const maxEchoBodyBytes = 1 << 20 // 1 MiB

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
	r.Body = http.MaxBytesReader(w, r.Body, maxEchoBodyBytes)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			log.Warn().Msg("Rejected /echo request: payload too large")
			writeJSONError(w, http.StatusRequestEntityTooLarge, "Payload too large (max 1MiB)")
			return
		}
		log.Error().Err(err).Msg("Failed to decode /echo request body")
		writeJSONError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	if req.Message == "" {
		writeJSONError(w, http.StatusBadRequest, "Invalid input")
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

// writeJSONError writes a JSON error payload with the provided status code.
func writeJSONError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
