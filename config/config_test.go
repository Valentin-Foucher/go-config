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

	value, err := c.MustGetString("anyKey")

	assert.Equal(t, "", value)
	assert.EqualError(t, err, "key \"anyKey\" not found")

	value, err = c.MustGetString("0")

	assert.Equal(t, "", value)
	assert.EqualError(t, err, "expected slice got: map[interface {}]interface {}: map[]")
}

func TestMissingKey(t *testing.T) {
	c := getTestConfig()

	value, err := c.MustGetInt("doesnotexist")
	assert.Equal(t, 0, value)
	assert.EqualError(t, err, "key \"doesnotexist\" not found")

}
func TestInt(t *testing.T) {
	c := getTestConfig()

	value, err := c.MustGetInt("int")
	assert.Equal(t, 5, value)
	assert.Nil(t, err)

	value = c.GetIntOrDefault("doesnoexist", 17)
	assert.Equal(t, 17, value)

	value, err = c.MustGetInt("float")
	assert.Equal(t, 0, value)
	assert.EqualError(t, err, "expected a int, got float64: float64")
}

func TestFloat(t *testing.T) {
	c := getTestConfig()

	value, err := c.MustGetFloat("float")
	assert.Equal(t, 5.123456, value)
	assert.Nil(t, err)

	value = c.GetFloatOrDefault("doesnoexist", 17.123456)
	assert.Equal(t, 17.123456, value)

	value, err = c.MustGetFloat("int")
	assert.Equal(t, 0.0, value)
	assert.EqualError(t, err, "expected a float64, got int: int")
}

func TestString(t *testing.T) {
	c := getTestConfig()

	value, err := c.MustGetString("string")
	assert.Equal(t, "test", value)
	assert.Nil(t, err)

	value = c.GetStringOrDefault("doesnoexist", "fallback")
	assert.Equal(t, "fallback", value)

	value, err = c.MustGetString("bool")
	assert.Equal(t, "", value)
	assert.EqualError(t, err, "expected a string, got bool: bool")
}

func TestBool(t *testing.T) {
	c := getTestConfig()

	value, err := c.MustGetBool("bool")
	assert.Equal(t, false, value)
	assert.Nil(t, err)

	value = c.GetBoolOrDefault("doesnoexist", true)
	assert.Equal(t, true, value)

	value, err = c.MustGetBool("string")
	assert.Equal(t, false, value)
	assert.EqualError(t, err, "expected a bool, got string: string")
}

func TestSlice(t *testing.T) {
	c := getTestConfig()

	value, err := c.MustGetSlice("test-slice")
	assert.Equal(t, []any{1, 2, 3}, value)
	assert.Nil(t, err)

	value = c.GetSliceOrDefault("doesnoexist", []any{123, 456})
	assert.Equal(t, []any{123, 456}, value)

	_, err = c.MustGetSlice("test-map")
	assert.EqualError(t, err, "expected a []interface {}, got map[interface {}]interface {}: map[interface {}]interface {}")
}

func TestMap(t *testing.T) {
	c := getTestConfig()

	value, err := c.MustGetMap("test-map")
	assert.Equal(t, map[any]any{
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
	}, value)
	assert.Nil(t, err)

	value = c.GetMapOrDefault("doesnoexist", map[any]any{123: 456})
	assert.Equal(t, map[any]any{123: 456}, value)

	_, err = c.MustGetMap("test-slice")
	assert.EqualError(t, err, "expected a map[interface {}]interface {}, got []interface {}: []interface {}")
}

func TestInSlice(t *testing.T) {
	c := getTestConfig()

	value, err := c.MustGetInt("test-slice.2")
	assert.Equal(t, 3, value)
	assert.Nil(t, err)
}

func TestInMap(t *testing.T) {
	c := getTestConfig()

	value, err := c.MustGetString("test-map.string")
	assert.Equal(t, "s", value)
	assert.Nil(t, err)
}

func TestNested(t *testing.T) {
	c := getTestConfig()

	value, err := c.MustGetString("test-map.ultra-nested.0.leaf")
	assert.Equal(t, "leaf", value)
	assert.Nil(t, err)
}

func TestListKeys(t *testing.T) {
	c := getTestConfig()

	keys, err := c.ListKeys("test-map")
	assert.Contains(t, keys, "string", "test-nested-map")
	assert.Contains(t, keys, "test-nested-map")
	assert.Contains(t, keys, "test-nested-slice")
	assert.Contains(t, keys, "ultra-nested")
	assert.Nil(t, err)

	_, err = c.ListKeys("test-slice")
	assert.EqualError(t, err, "expected a map[interface {}]interface {}, got []interface {}: []interface {}")

	_, err = c.ListKeys("doesnotexist")
	assert.EqualError(t, err, "key \"doesnotexist\" not found")
}

func TestMustGetType(t *testing.T) {
	c := getTestConfig()

	type a struct {
		Leaf string
	}

	value, err := MustGetType[a](c, "test-map.ultra-nested.0")
	assert.Equal(t, "leaf", value.Leaf)
	assert.Nil(t, err)
}
