// Package helpers contains global functions
package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env function to get environment variable value
func Env(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
