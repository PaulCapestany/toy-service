// env.go
//
// Provides helper functions to retrieve environment variables with defaults.

package handlers

import "os"

// getEnv retrieves the value of the environment variable named by the key,
// or returns the provided default if the variable is not set.
func getEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

// EnvConfig holds environment-dependent configuration values.
type EnvConfig struct {
	Env          string
	LogVerbosity string
	FakeSecret   string
	Version      string
	GitCommit    string
	Name         string
}

// LoadEnvConfig loads configuration from environment variables.
// TODO: generate/pull these values in dynamically.
func LoadEnvConfig() EnvConfig {
    return EnvConfig{
        Env:          getEnv("SERVICE_ENV", "dev"),
        LogVerbosity: getEnv("LOG_VERBOSITY", "info"),
        FakeSecret:   getEnv("FAKE_SECRET", "redacted"),
		Version:      getEnv("VERSION", "v0.3.26"),
        GitCommit:    getEnv("GIT_COMMIT", "unknown"),
        Name:         "toy-service",
    }
}
