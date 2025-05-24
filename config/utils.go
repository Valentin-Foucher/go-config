package config

import "strconv"

func asInteger(s string) (int, error) {
	return strconv.Atoi(s)
}
