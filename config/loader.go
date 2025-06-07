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

func LoadType(filePath string, configType configType) (Config, error) {
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

func Load(filePath string) (Config, error) {
	for _, load := range []func(string) (Config, error){loadJSON, loadYAML, loadTOML} {
		if config, err := load(filePath); err == nil {
			return config, nil
		}
	}

	return nil, errors.New("unrecognized configuration file (supported: json, yaml, toml)")
}
