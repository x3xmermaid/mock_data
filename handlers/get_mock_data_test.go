package handlers_test

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	nhandler "ketitik/netmonk/mock-app-data/handlers"
	nrespwriter "ketitik/netmonk/mock-app-data/lib/responsewriter"
	ntesting "ketitik/netmonk/mock-app-data/testing"

	"github.com/gorilla/mux"
)

func TestGetXMLMockData(t *testing.T) {
	// Create Handler Object
	handler := nhandler.NewHandler(nil)

	// Setup HTTP Server
	r := mux.NewRouter()
	r.HandleFunc("/mock_data", handler.GetXMLMockData).Methods("POST")
	httpServer := ntesting.Setup(r)
	defer httpServer.Close()
	serverURL, _ := url.Parse(httpServer.URL)

	// Hit API Endpoint
	targetPath := fmt.Sprintf("%v/%v", serverURL, "mock_data?nd=1000000000001&realm=jakarta")
	req, _ := http.NewRequest("POST", targetPath, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Unable to get worker status: %v", err)
	}
	defer resp.Body.Close()

	t.Run("GetXMLMockData_OK", func(t *testing.T) {
		// Hit API Endpoint
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to get worker status: %v", err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Response code %v", resp.StatusCode)
		}

		resp.Body.Close()
	})

}

func extractContent(encodedBody []byte, content interface{}) error {
	// Unmarshal XML Resp
	var body nrespwriter.ResponseFormat
	err := xml.Unmarshal(encodedBody, &body)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal xml response: %v", err)
	}

	encodedContent, _ := xml.Marshal(body.Data)
	err = xml.Unmarshal(encodedContent, &content)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal content: %v", err)
	}
	return nil
}
