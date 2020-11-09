package testing

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestSetup(t *testing.T) {
	r := mux.NewRouter()
	httpServer := Setup(r)
	if httpServer == nil {
		t.Fatalf("It should not nil")
	}
}
