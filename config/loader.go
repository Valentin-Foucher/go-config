package config

import (
	"errors"
)

type configType int

const (
	YAML configType = iota
	JSON
	TOML
)

func Load(filePath string, configType configType) (Config, error) {
	switch configType {
	case YAML:
		return loadYAML(filePath)
	case JSON:
		return loadJSON(filePath)
	case TOML:
		return loadTOML(filePath)
	}

	return nil, errors.New("unsupported config type")
}
