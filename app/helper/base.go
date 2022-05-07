package helper

import (
	"github.com/google/uuid"
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

// GetIntEnv finds an env variable and parse it to integer (int32)
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

// GenerateUUID generates an UUID
func GenerateUUID() string {
	return uuid.New().String()
}
