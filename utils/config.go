package utils

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
