package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type EnvTypes interface {
	string | []string | int | time.Duration
}

// LookupEnv is a generic type implementation to search env keys
func LookupEnv[T EnvTypes](name string, defaultValue T) T {
	value, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}

	var result any
	switch any(defaultValue).(type) {
	case string:
		result = value
	case []string:
		// Should be a comma separated list
		strs := strings.Split(value, ",")
		result = strs
	case int:
		i, _ := strconv.ParseInt(value, 10, 64)
		result = int(i)
	case time.Duration:
		var err error
		if result, err = time.ParseDuration(value); err != nil {
			return defaultValue
		}
	}

	return result.(T)
}
