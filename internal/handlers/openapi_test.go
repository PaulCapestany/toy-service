package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

// loadOpenAPISpec loads and parses the OpenAPI spec from the local spec/openapi.yaml file.
func loadOpenAPISpec(t *testing.T) *openapi3.T {
	t.Helper()

	specPath := filepath.Join("..", "..", "spec", "openapi.yaml")
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	swagger, err := loader.LoadFromFile(specPath)
	require.NoError(t, err, "Failed to load OpenAPI spec")

	// Validate the spec to ensure correctness (including examples).
	err = swagger.Validate(context.Background())
	require.NoError(t, err, "OpenAPI spec validation failed")

	return swagger
}

// newTestServer starts a httptest server with the handlers registered.
func newTestServer() *httptest.Server {
	r := chi.NewRouter()
	r.Get("/healthz", HealthzHandler)
	r.Post("/echo", EchoHandler)
	r.Get("/info", InfoHandler)
	r.Get("/version", VersionHandler)

	return httptest.NewServer(r)
}

func TestOpenAPIConformance(t *testing.T) {
	log.Debug().Msg("Starting OpenAPI conformance tests")

	swagger := loadOpenAPISpec(t)
	server := newTestServer()
	defer server.Close()

	healthzPath := "/healthz"
	echoPath := "/echo"
	infoPath := "/info"
	versionPath := "/version"

	// 1. Test /healthz
	t.Run("GET /healthz", func(t *testing.T) {
		resp, err := http.Get(server.URL + healthzPath)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		validateResponse(t, swagger, "get", healthzPath, resp.StatusCode, body)
	})

	// 2. Test /info
	t.Run("GET /info", func(t *testing.T) {
		resp, err := http.Get(server.URL + infoPath)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		validateResponse(t, swagger, "get", infoPath, resp.StatusCode, body)
	})

	// 3. Test /version
	t.Run("GET /version", func(t *testing.T) {
		resp, err := http.Get(server.URL + versionPath)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		validateResponse(t, swagger, "get", versionPath, resp.StatusCode, body)
	})

	// 4. Test /echo with valid input
	t.Run("POST /echo with valid input", func(t *testing.T) {
		reqBody := `{"message":"Hello"}`
		resp, err := http.Post(server.URL+echoPath, "application/json", bytes.NewBuffer([]byte(reqBody)))
		require.NoError(t, err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		validateResponse(t, swagger, "post", echoPath, resp.StatusCode, body)
	})

	// 5. Test /echo with invalid input (missing required "message" field)
	t.Run("POST /echo with invalid input", func(t *testing.T) {
		reqBody := `{"msg":"NoMessageField"}`
		resp, err := http.Post(server.URL+echoPath, "application/json", bytes.NewBuffer([]byte(reqBody)))
		require.NoError(t, err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)

		validateResponse(t, swagger, "post", echoPath, resp.StatusCode, body)
	})
}

// validateResponse uses the loaded swagger and the provided method/path/status to look up the expected schema.
// It then validates the response body against that schema.
func validateResponse(t *testing.T, swagger *openapi3.T, method, path string, statusCode int, body []byte) {
	t.Helper()

	// Use Paths.Find to locate the PathItem associated with this path.
	pathItem := swagger.Paths.Find(path)
	require.NotNil(t, pathItem, "Path %s not found in OpenAPI spec", path)

	operation := getOperationForMethod(t, pathItem, method)
	require.NotNil(t, operation, "No operation defined for %s %s", method, path)

	responseRef := operation.Responses.Status(statusCode)
	require.NotNil(t, responseRef, "No response defined for %d on %s %s", statusCode, method, path)
	require.NotNil(t, responseRef.Value, "Response value is nil for %d on %s %s", statusCode, method, path)

	// Check if there's a JSON schema defined
	jsonContent, hasJSON := responseRef.Value.Content["application/json"]
	if !hasJSON {
		t.Fatalf("No application/json response schema found for %d on %s %s", statusCode, method, path)
	}

	schemaRef := jsonContent.Schema
	require.NotNil(t, schemaRef, "No schema defined for application/json in %d response on %s %s", statusCode, method, path)

	// Parse JSON response into interface{}
	var data interface{}
	err := json.Unmarshal(body, &data)
	require.NoError(t, err, "Failed to unmarshal response body to JSON: %s", string(body))

	// Validate data against the schema
	err = schemaRef.Value.VisitJSON(data)
	require.NoError(t, err, "Response body does not match OpenAPI schema for %s %s %d: %s", method, path, statusCode, string(body))
}

func getOperationForMethod(t *testing.T, pathItem *openapi3.PathItem, method string) *openapi3.Operation {
	t.Helper()
	switch method {
	case "get":
		return pathItem.Get
	case "post":
		return pathItem.Post
	case "put":
		return pathItem.Put
	case "delete":
		return pathItem.Delete
	case "options":
		return pathItem.Options
	case "head":
		return pathItem.Head
	case "patch":
		return pathItem.Patch
	case "trace":
		return pathItem.Trace
	default:
		t.Fatalf("Unsupported method: %s", method)
		return nil
	}
}
