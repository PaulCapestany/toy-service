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
	t.Log("Test that /info returns service metadata")

	r := chi.NewRouter()
	r.Get("/info", InfoHandler)

	req, err := http.NewRequest("GET", "/info", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	require.Equal(t, "toy-service", resp["name"])
	require.NotEmpty(t, resp["version"])
	require.NotEmpty(t, resp["env"])
	require.NotEmpty(t, resp["logVerbosity"])
	require.NotEmpty(t, resp["fakeSecret"])
	require.NotEmpty(t, resp["commit"])
}
