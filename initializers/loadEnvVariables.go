package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load environment variables from .env file
	// This function is called in the main function to load the environment variables
	// before starting the server.
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
