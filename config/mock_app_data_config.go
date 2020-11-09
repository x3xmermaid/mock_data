package config

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// ServiceConfig stores the whole configuration for service.
type ServiceConfig struct {
	ServiceData ServiceDataConfig `yaml:"service_data"`
}

// ServiceDataConfig contains the service data configuration.
type ServiceDataConfig struct {
	Address string `yaml:"address"`
}

func getRawConfig(fileLocation string) (*ServiceConfig, error) {
	configByte, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		logrus.Errorf("Error Read File Raw Config: %v", err)
		return nil, err
	}
	config := &ServiceConfig{}
	err = yaml.Unmarshal(configByte, config)
	if err != nil {
		logrus.Errorf("Error Unmarshal Raw Config: %v", err)
		return nil, err
	}
	return config, nil
}

// GetServiceConfig parse the configuration from YAML file.
func (c *YAMLConfigLoader) GetServiceConfig() (*ServiceConfig, error) {
	config, err := getRawConfig(c.fileLocation)
	if err != nil {
		return nil, fmt.Errorf("Unable to get raw config content: %v", err)
	}
	return config, nil
}
