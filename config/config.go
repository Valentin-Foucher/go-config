package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Config map[any]any

func (c Config) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(b)
}

func (c Config) MustGetString(key string) (string, error) {
	value, err := getForType[string](c, key)
	if err != nil {
		return "", err
	}

	if isEnvVariableReference(value) {
		value = loadEnvVariable(value)
	}

	return value, nil
}

func (c Config) MustGetInt(key string) (int, error) {
	return getForType[int](c, key)
}

func (c Config) MustGetFloat(key string) (float64, error) {
	return getForType[float64](c, key)
}

func (c Config) MustGetBool(key string) (bool, error) {
	return getForType[bool](c, key)
}

func (c Config) MustGetMap(key string) (map[any]any, error) {
	return getForType[map[any]any](c, key)
}

func (c Config) MustGetSlice(key string) ([]any, error) {
	return getForType[[]any](c, key)
}

func (c Config) GetStringOrDefault(key string, defaultValue string) string {
	return getOrDefault(c, key, defaultValue)
}

func (c Config) GetIntOrDefault(key string, defaultValue int) int {
	return getOrDefault(c, key, defaultValue)
}

func (c Config) GetFloatOrDefault(key string, defaultValue float64) float64 {
	return getOrDefault(c, key, defaultValue)
}

func (c Config) GetBoolOrDefault(key string, defaultValue bool) bool {
	return getOrDefault(c, key, defaultValue)
}

func (c Config) GetMapOrDefault(key string, defaultValue map[any]any) map[any]any {
	return getOrDefault(c, key, defaultValue)
}

func (c Config) GetSliceOrDefault(key string, defaultValue []any) []any {
	return getOrDefault(c, key, defaultValue)
}

func MustGetType[T any](c Config, key string) (T, error) {
	var result T

	value, err := c.MustGetMap(key)
	if err != nil {
		return result, err
	}

	err = mapstructure.Decode(value, &result)
	return result, err
}

func (c Config) ListKeys(key string) ([]any, error) {
	value, err := c.MustGetMap(key)
	if err != nil {
		return nil, err
	}

	var keys []any
	for k := range value {
		keys = append(keys, k)
	}

	return keys, nil
}

func getForType[T any](c Config, key string) (T, error) {
	var dummy T

	v, err := c.get(key)
	if err != nil {
		return dummy, err
	}

	b, ok := v.(T)
	if !ok {
		return dummy, fmt.Errorf("expected a %[1]T, got %[2]T: %[2]T", dummy, v)
	}

	return b, nil
}

func getOrDefault[T any](c Config, key string, defaultValue T) T {
	value, err := getForType[T](c, key)
	if err != nil {
		return defaultValue
	}

	return value
}

func (c Config) get(key string) (any, error) {
	path := strings.Split(key, ".")

	if len(path) == 0 {
		return nil, errors.New("key cannot be empty")
	}

	return c.getValue(path, map[any]any(c))
}

func (c Config) getValue(path []string, node any) (any, error) {
	next, err := c.getChildElement(path, node)
	if err != nil {
		return nil, err
	}

	if len(path) == 1 {
		return next, nil
	}

	return c.getValue(path[1:], next)
}

func (c Config) getChildElement(path []string, node any) (any, error) {
	key := path[0]
	index, err := asInteger(key)
	if err == nil {
		return c.getSliceElement(index, node)
	}

	return c.getMapElement(key, node)
}

func (c Config) getSliceElement(index int, node any) (any, error) {
	nodeSlice, isSlice := node.([]any)
	if !isSlice {
		return nil, fmt.Errorf("expected slice got: %[1]T: %[1]v", node)
	}

	if len(nodeSlice) < index-1 || index < 0 {
		return nil, fmt.Errorf("invalid index %d", index)
	}

	return nodeSlice[index], nil
}

func (c Config) getMapElement(key string, node any) (any, error) {
	nodeMap, isMap := node.(map[any]any)
	if !isMap {
		return nil, fmt.Errorf("expected map got %[1]T: %[1]v", node)
	}

	value, exists := nodeMap[key]
	if !exists {
		return nil, fmt.Errorf("key \"%s\" not found", key)
	}

	return value, nil
}
