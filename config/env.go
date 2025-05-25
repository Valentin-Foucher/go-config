package config

import (
	"os"
	"strings"
)

const envPrefix = "---ENV "

func isEnvVariableReference(value string) bool {
	return strings.HasPrefix(value, envPrefix)
}

func loadEnvVariable(value string) string {
	return os.Getenv(strings.Replace(value, envPrefix, "", 1))
}
