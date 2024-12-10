package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestEchoHandler(t *testing.T) {
	t.Log("Test that /echo returns a modified message and service metadata")

	r := chi.NewRouter()
	r.Post("/echo", EchoHandler)

	reqBody := `{"message":"Hello"}`
	req, err := http.NewRequest("POST", "/echo", bytes.NewBuffer([]byte(reqBody)))
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var resp EchoResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	require.Contains(t, resp.Message, "Hello [modified]")
	require.NotEmpty(t, resp.Version)
	require.NotEmpty(t, resp.Commit)
	require.NotEmpty(t, resp.Env)
}
