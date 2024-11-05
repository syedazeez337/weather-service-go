package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func GetAPIKey() (string, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("API key not set in environment variable WEATHER_API_KEY")
	}
	return apiKey, nil
}