package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestInfoHandler(t *testing.T) {
	t.Log("Test that /info returns service metadata without exposing secrets")

	r := chi.NewRouter()
	r.Get("/info", InfoHandler)

	req, err := http.NewRequest("GET", "/info", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "no-store", w.Header().Get("Cache-Control"))

	var resp struct {
		Name              string `json:"name"`
		Version           string `json:"version"`
		Env               string `json:"env"`
		LogVerbosity      string `json:"logVerbosity"`
		FakeSecretPresent bool   `json:"fakeSecretPresent"`
		FakeSecretLength  int    `json:"fakeSecretLength"`
		Commit            string `json:"commit"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	require.Equal(t, "toy-service", resp.Name)
	require.NotEmpty(t, resp.Version)
	require.NotEmpty(t, resp.Env)
	require.NotEmpty(t, resp.LogVerbosity)
	require.False(t, resp.FakeSecretPresent)
	require.Equal(t, 0, resp.FakeSecretLength)
	require.NotEmpty(t, resp.Commit)
}

func TestInfoHandlerWithSecret(t *testing.T) {
	t.Log("Test that /info reports secret presence without revealing value")

	const secret = "super-secret"

	r := chi.NewRouter()
	r.Get("/info", InfoHandler)

	t.Setenv("FAKE_SECRET", secret)

	req, err := http.NewRequest("GET", "/info", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var resp struct {
		FakeSecretPresent bool `json:"fakeSecretPresent"`
		FakeSecretLength  int  `json:"fakeSecretLength"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	require.True(t, resp.FakeSecretPresent)
	require.Equal(t, len(secret), resp.FakeSecretLength)
}
