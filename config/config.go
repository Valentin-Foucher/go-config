package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Config map[any]any

func (c Config) GetString(key string) (string, error) {
	v, err := c.get(key)
	if err != nil {
		return "", err
	}

	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("expected a string, got %[1]T: %[1]T", v)
	}

	return s, nil
}

func (c Config) GetInt(key string) (int, error) {
	v, err := c.get(key)
	if err != nil {
		return 0, err
	}

	i, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("expected an int, got %[1]T: %[1]T", v)
	}

	return i, nil
}

func (c Config) GetFloat(key string) (float64, error) {
	v, err := c.get(key)
	if err != nil {
		return 0, err
	}

	f, ok := v.(float64)
	if !ok {
		return 0, fmt.Errorf("expected a float, got %[1]T: %[1]T", v)
	}

	return f, nil
}

func (c Config) GetBool(key string) (bool, error) {
	v, err := c.get(key)
	if err != nil {
		return false, err
	}

	b, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("expected a bool, got %[1]T: %[1]T", v)
	}

	return b, nil
}

func (c Config) get(key string) (any, error) {
	path := strings.Split(key, ".")

	if len(path) == 0 {
		return nil, errors.New("key cannot be empty")
	}

	return c.getValue(path, map[any]any(c))
}

func (c Config) getValue(path []string, node any) (any, error) {
	next, err := c.getNext(path, node)
	if err != nil {
		return nil, err
	}

	if len(path) == 1 {
		return next, nil
	}

	return c.getValue(path[1:], next)
}

func (c Config) getNext(path []string, node any) (any, error) {
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

func (c Config) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(b)
}
