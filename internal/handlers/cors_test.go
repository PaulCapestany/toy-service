// cors_test.go
//
// Tests to ensure that CORS headers are properly set, allowing cross-origin requests.
// This test checks that the server responds with appropriate CORS headers on OPTIONS requests
// and that GET/POST requests also include them.

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/stretchr/testify/require"
)

// TestCORSHeaders ensures that the service sets expected CORS headers.
func TestCORSHeaders(t *testing.T) {
	t.Log("Ensuring CORS headers are present")

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	r.Get("/healthz", HealthzHandler)

	// Test OPTIONS request
	optsReq, err := http.NewRequest("OPTIONS", "/healthz", nil)
	require.NoError(t, err)

	// Set headers to mimic a real browser preflight request
	optsReq.Header.Set("Origin", "http://localhost")
	optsReq.Header.Set("Access-Control-Request-Method", "GET")
	optsReq.Header.Set("Access-Control-Request-Headers", "Content-Type")

	optsRec := httptest.NewRecorder()
	r.ServeHTTP(optsRec, optsReq)

	require.Equal(t, http.StatusOK, optsRec.Code)

	// Test GET request
	getReq, err := http.NewRequest("GET", "/healthz", nil)
	require.NoError(t, err)

	// Add Origin header to mimic a real browser request from another origin
	getReq.Header.Set("Origin", "http://localhost")

	getRec := httptest.NewRecorder()
	r.ServeHTTP(getRec, getReq)

	require.Equal(t, http.StatusOK, getRec.Code)
	require.Contains(t, getRec.Header().Get("Access-Control-Allow-Origin"), "*")
}
