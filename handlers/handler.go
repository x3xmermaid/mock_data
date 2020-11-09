package handlers

import (
	nconfig "ketitik/netmonk/mock-app-data/config"
)

// Handler holds the API endpoint's function handler.
type Handler struct {
	Config *nconfig.ServiceDataConfig
}

// NewHandler function to make connection database into handler
func NewHandler(config *nconfig.ServiceDataConfig) *Handler {
	return &Handler{
		Config: config,
	}
}
