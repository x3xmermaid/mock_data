package testing

import (
	"net/http/httptest"

	"github.com/gorilla/mux"
)

const (
	// IntentionallyError is the intentional error message for testing.
	IntentionallyError = "Error created intentionally."
)

// Setup sets up a test HTTP server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
// It is inspired by go-octokit.
func Setup(r *mux.Router) *httptest.Server {
	// test server
	server := httptest.NewServer(r)
	return server
}
