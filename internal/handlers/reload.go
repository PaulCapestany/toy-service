package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "strings"

    "github.com/rs/zerolog/log"
)

// ReloadHandler reads the FAKE_SECRET value from a mounted secret file and updates
// the process environment so handlers that read os.Getenv can see the new value.
//
// By default it looks under /etc/backend-secret/FAKE_SECRET, but the base path can
// be overridden via SECRET_FILE_DIR. This pairs with the Helm chart which mounts
// the Secret at /etc/backend-secret by default.
func ReloadHandler(w http.ResponseWriter, r *http.Request) {
    base := os.Getenv("SECRET_FILE_DIR")
    if base == "" {
        base = "/etc/backend-secret"
    }
    path := base + "/FAKE_SECRET"

    data, err := os.ReadFile(path)
    if err != nil {
        log.Error().Err(err).Str("path", path).Msg("failed reading secret file")
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        _ = json.NewEncoder(w).Encode(map[string]string{"error": "failed to read secret file"})
        return
    }

    // Trim trailing newlines/whitespace if present (kube Secret keys are raw bytes)
    val := strings.TrimRight(string(data), "\r\n")
    // Update process env so subsequent os.Getenv reads see the new value
    if err := os.Setenv("FAKE_SECRET", val); err != nil {
        log.Error().Err(err).Msg("failed setting env")
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        _ = json.NewEncoder(w).Encode(map[string]string{"error": "failed setting env"})
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
