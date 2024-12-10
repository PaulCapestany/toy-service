// healthz_test.go
//
// Tests for the /healthz endpoint to ensure it returns the expected response.
// We use the standard library testing package and the testify library for assertions.

package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestHealthzHandler(t *testing.T) {
	t.Log("Test that /healthz returns a 200 and {'status':'ok'}")

	r := chi.NewRouter()
	r.Get("/healthz", HealthzHandler)

	req, err := http.NewRequest("GET", "/healthz", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	require.Equal(t, "ok", resp["status"], "Expected status to be 'ok'")
}
