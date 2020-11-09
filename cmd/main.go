package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	nconfig "ketitik/netmonk/mock-app-data/config"
	nhandler "ketitik/netmonk/mock-app-data/handlers"
)

const (
	// VERSION of this service.
	VERSION = "1.0.0"
	// DEVELOPER of this service.
	DEVELOPER = "Amoeba Ketitik"
	// ConfigFileLocation is the file configuration of ths service.
	ConfigFileLocation = "conf/netmonk_config.yaml"
	ketitikASCIIImage  = `
	888  888 88888888 88888888 88888888 88888888 88888888 888  888
	888  888 88          88       88       88       88    888  888
	888 888  88          88       88       88       88    888 888
	88888    88888888    88       88       88       88    88888
	888 888  88          88       88       88       88    888 888
	888  888 88          88       88       88       88    888  888
	888  888 88888888    88    88888888    88    88888888 888  888
	`
	// timeout
	timeout = 10
)

// Handler hold the function handler for API's endpoint.
type Handler interface {
	GetXMLMockData(w http.ResponseWriter, r *http.Request)
}

// NewRouter returns router.
func NewRouter(handler Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/mock_data", handler.GetXMLMockData).Methods("POST")
	return r
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	if len(arr) >= 2 {
		return fmt.Sprintf("%s/%s", arr[len(arr)-2], arr[len(arr)-1])
	}

	return arr[len(arr)-1]
}

func main() {
	customFormatter := &logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", formatFilePath(f.File), f.Line)
		},
	}

	logrus.SetFormatter(customFormatter)
	logrus.SetReportCaller(true)
	// Pre-printed text at startup.
	logrus.Printf("Netmonk Mock-App-Data v%v", VERSION)
	logrus.Printf("Developed by %v.\n%v", DEVELOPER, ketitikASCIIImage)
	logrus.Println("Start service...")

	// Get Config
	configLoader := nconfig.NewYamlConfigLoader(ConfigFileLocation)
	config, err := configLoader.GetServiceConfig()
	if err != nil {
		logrus.Fatalf("Unable to load configuration: %v", err)
	}

	handler := nhandler.NewHandler(&config.ServiceData)
	r := NewRouter(handler)

	// Run Web Server
	logrus.Printf("Starting http server at %v", config.ServiceData.Address)
	err = http.ListenAndServe(config.ServiceData.Address, r)
	if err != nil {
		logrus.Fatalf("Unable to run http server: %v", err)
	}
	logrus.Println("Stopping API Service...")
}
