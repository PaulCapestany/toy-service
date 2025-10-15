package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReloadHandler_Success(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("SECRET_FILE_DIR", dir)
	t.Setenv("FAKE_SECRET", "old-secret")

	secretValue := "new-secret\n"
	require.NoError(t, os.WriteFile(filepath.Join(dir, "FAKE_SECRET"), []byte(secretValue), 0o600))

	req := httptest.NewRequest(http.MethodPost, "/-/reload", nil)
	rr := httptest.NewRecorder()

	ReloadHandler(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var resp reloadResponse
	require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))

	assert.Equal(t, "ok", resp.Status)
	assert.Equal(t, len("new-secret"), resp.FakeSecretLen)
	assert.Equal(t, "new-secret", os.Getenv("FAKE_SECRET"))
}

func TestReloadHandler_ReadFailure(t *testing.T) {
	t.Setenv("SECRET_FILE_DIR", filepath.Join(t.TempDir(), "missing"))
	t.Setenv("FAKE_SECRET", "should-stay")

	req := httptest.NewRequest(http.MethodPost, "/-/reload", nil)
	rr := httptest.NewRecorder()

	ReloadHandler(rr, req)

	require.Equal(t, http.StatusInternalServerError, rr.Code)
	require.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var body map[string]string
	require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &body))
	assert.Equal(t, "failed to read secret file", body["error"])

	assert.Equal(t, "should-stay", os.Getenv("FAKE_SECRET"))
}
