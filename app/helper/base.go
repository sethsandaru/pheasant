package helper

import (
	"log"
	"os"
	"strconv"
)

// GetEnv finds an env variable or the given fallback.
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}

func GetIntEnv(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("ENV KEY %s need to be an integer", key)
	}

	return intValue
}
