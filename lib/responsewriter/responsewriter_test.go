package responsewriter

import (
	"net/http"
	"testing"
)

type mockedHTTPWriter struct{}

func (w *mockedHTTPWriter) Header() http.Header {
	return http.Header{}
}

func (w *mockedHTTPWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (w *mockedHTTPWriter) WriteHeader(statusCode int) {
	return
}

func TestResponseOK(t *testing.T) {
	w := &mockedHTTPWriter{}
	rf := &ResponseFormat{}
	t.Run("OK", func(t *testing.T) {
		rf.ResponseOK(http.StatusOK, nil, w)
	})
	t.Run("NOK", func(t *testing.T) {
		data := make(chan int)
		rf.ResponseOK(http.StatusNotImplemented, data, w)
	})
}

func TestResponseNOK(t *testing.T) {
	w := &mockedHTTPWriter{}
	rf := &ResponseFormat{}
	t.Run("OK", func(t *testing.T) {
		rf.ResponseNOK(http.StatusNotImplemented, nil, w)
	})
	t.Run("NOK", func(t *testing.T) {
		data := make(chan int)
		rf.ResponseNOK(http.StatusOK, data, w)
	})
}

func TestGetLastPath(t *testing.T) {
	path := "first/last"
	lastPath := GetLastPath(path)
	if lastPath != "last" {
		t.Fatalf("It should be naqvi instead of %v", lastPath)
	}
}
