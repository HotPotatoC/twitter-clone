package config

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Load reads .env file configurations into ENV
func Load(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	if err := godotenv.Load(path); err != nil {
		return err
	}

	return nil
}

func lookup(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetString(key string, fallback string) string {
	return lookup(key, fallback)
}

func GetInt(key string, fallback int) int {
	value := lookup(key, "")
	if value, err := strconv.Atoi(value); err == nil {
		return value
	}
	return fallback
}

func GetBool(key string, fallback bool) bool {
	value := lookup(key, "")
	if value, err := strconv.ParseBool(value); err == nil {
		return value
	}
	return fallback
}

func GetDuration(key string, fallback time.Duration) time.Duration {
	value := lookup(key, "")
	if value, err := time.ParseDuration(value); err == nil {
		return value
	}
	return fallback
}
