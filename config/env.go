package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DotEnvVariable -> get .env
func Getenv(key string) string {

	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
