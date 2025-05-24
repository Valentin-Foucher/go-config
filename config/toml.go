package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

func loadTOML(filePath string) (Config, error) {
	var configMap map[any]any

	if _, err := toml.DecodeFile(filePath, &configMap); err != nil {
		log.Fatalf("Error loading TOML file: %v", err)
	}

	return configMap, nil
}
