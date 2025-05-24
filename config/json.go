package config

import (
	"encoding/json"
	"os"
)

func loadJSON(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	configMap := make(map[any]any)

	if err := json.Unmarshal(data, &configMap); err != nil {
		return nil, err
	}

	return configMap, nil
}
