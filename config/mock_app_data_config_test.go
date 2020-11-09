package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nconfig "ketitik/netmonk/mock-app-data/config"
)

func TestGetServiceConfigOK(t *testing.T) {
	yamlFileLocation := "_fixtures_netmonk_bec_ga/config_01.yaml"
	configLoader := nconfig.NewYamlConfigLoader(yamlFileLocation)
	config, err := configLoader.GetServiceConfig()
	if err != nil {
		t.Fatalf("It should be OK: %v", err)
	}
	if config == nil {
		t.Fatalf("It should be not nil: %v", err)
	}
	// Check the source data.
	assert.Equal(t, "localhost:8080", config.ServiceData.Address, "they should be equal")
}

func TestGetServiceConfigNOK(t *testing.T) {
	t.Run("Wrong Configuration Structure", func(t *testing.T) {
		yamlFileLocation := "_fixtures_netmonk_bec_ga/config_02.yaml"
		configLoader := nconfig.NewYamlConfigLoader(yamlFileLocation)
		_, err := configLoader.GetServiceConfig()
		if err == nil {
			t.Fatalf("It should be NOK.")
		}
	})
	t.Run("No File Found", func(t *testing.T) {
		yamlFileLocation := "_fixtures_netmonk_bec_ga/config_03.yaml"
		configLoader := nconfig.NewYamlConfigLoader(yamlFileLocation)
		_, err := configLoader.GetServiceConfig()
		if err == nil {
			t.Fatalf("It should be NOK.")
		}
	})
}
