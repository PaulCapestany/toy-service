package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestVersionHandler(t *testing.T) {
	t.Log("Test that /version returns service build metadata")

	r := chi.NewRouter()
	r.Get("/version", VersionHandler)

	req, err := http.NewRequest("GET", "/version", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "no-store", w.Header().Get("Cache-Control"))

	var resp map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	require.Equal(t, "toy-service", resp["name"])
	require.NotEmpty(t, resp["version"])
	require.NotEmpty(t, resp["commit"])
}
