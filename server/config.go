package server

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

// LoadConfig loads the configuration from a specified YAML file and returns a Config struct.
func LoadConfig(filename string) (*Config, error) {
	// Read the file content
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", filename, err)
	}

	// Unmarshal the YAML content into the Config struct
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML from file %s: %w", filename, err)
	}

	return &config, nil
}
