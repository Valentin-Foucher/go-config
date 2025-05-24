package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func loadYAML(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	configMap := make(map[any]any)

	if err := yaml.Unmarshal(data, &configMap); err != nil {
		return nil, err
	}

	return configMap, nil
}
