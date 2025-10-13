package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestConfigHandler_PresenceFalse(t *testing.T) {
	prev := os.Getenv("FAKE_SECRET")
	t.Cleanup(func() { _ = os.Setenv("FAKE_SECRET", prev) })
	_ = os.Unsetenv("FAKE_SECRET")

	r := chi.NewRouter()
	r.Get("/internal/config", ConfigHandler)

	req, err := http.NewRequest("GET", "/internal/config", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	var resp ConfigSummary
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	require.False(t, resp.FakeSecretPresent)
	require.Equal(t, 0, resp.FakeSecretLen)
}

func TestConfigHandler_PresenceTrue(t *testing.T) {
	prev := os.Getenv("FAKE_SECRET")
	t.Cleanup(func() { _ = os.Setenv("FAKE_SECRET", prev) })
	require.NoError(t, os.Setenv("FAKE_SECRET", "supersecret"))

	r := chi.NewRouter()
	r.Get("/internal/config", ConfigHandler)

	req, err := http.NewRequest("GET", "/internal/config", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	var resp ConfigSummary
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	require.True(t, resp.FakeSecretPresent)
	require.Equal(t, len("supersecret"), resp.FakeSecretLen)
}
