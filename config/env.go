package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables(filename string) {
	godotenv.Load(filename)
	if os.Getenv("HOST_ADDRESS") == "" {
		os.Setenv("HOST_ADDRESS", "localhost:8001")
	}
}
