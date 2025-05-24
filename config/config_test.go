package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestConfig() Config {
	return Config{
		"test-map": map[any]any{
			"test-nested-map": map[any]any{
				"test-nested-bool": true,
			},
			"test-nested-slice": []any{"1", "2", "3"},
			"ultra-nested": []any{
				map[any]any{
					"leaf": "leaf",
				},
			},
			"string": "s",
		},
		"test-slice": []any{1, 2, 3},
		"int":        5,
		"float":      5.123456,
		"string":     "test",
		"bool":       false,
	}
}

func TestEmptyConfig(t *testing.T) {
	var c Config = make(Config)

	value, err := c.GetString("anyKey")

	assert.Equal(t, "", value)
	assert.EqualError(t, err, "key \"anyKey\" not found")

	value, err = c.GetString("0")

	assert.Equal(t, "", value)
	assert.EqualError(t, err, "expected slice got: map[interface {}]interface {}: map[]")
}

func TestInt(t *testing.T) {
	c := getTestConfig()

	value, err := c.GetInt("int")
	assert.Equal(t, 5, value)
	assert.Nil(t, err)

	value, err = c.GetInt("float")
	assert.Equal(t, 0, value)
	assert.EqualError(t, err, "expected an int, got float64: float64")
}

func TestFloat(t *testing.T) {
	c := getTestConfig()

	value, err := c.GetFloat("float")
	assert.Equal(t, 5.123456, value)
	assert.Nil(t, err)

	value, err = c.GetFloat("int")
	assert.Equal(t, 0.0, value)
	assert.EqualError(t, err, "expected a float, got int: int")
}

func TestString(t *testing.T) {
	c := getTestConfig()

	value, err := c.GetString("string")
	assert.Equal(t, "test", value)
	assert.Nil(t, err)

	value, err = c.GetString("bool")
	assert.Equal(t, "", value)
	assert.EqualError(t, err, "expected a string, got bool: bool")
}

func TestBool(t *testing.T) {
	c := getTestConfig()

	value, err := c.GetBool("bool")
	assert.Equal(t, false, value)
	assert.Nil(t, err)

	value, err = c.GetBool("string")
	assert.Equal(t, false, value)
	assert.EqualError(t, err, "expected a bool, got string: string")
}

func TestSlice(t *testing.T) {
	c := getTestConfig()

	value, err := c.GetInt("test-slice.2")
	assert.Equal(t, 3, value)
	assert.Nil(t, err)
}

func TestMap(t *testing.T) {
	c := getTestConfig()

	value, err := c.GetString("test-map.string")
	assert.Equal(t, "s", value)
	assert.Nil(t, err)
}

func TestNested(t *testing.T) {
	c := getTestConfig()

	value, err := c.GetString("test-map.ultra-nested.0.leaf")
	assert.Equal(t, "leaf", value)
	assert.Nil(t, err)
}
