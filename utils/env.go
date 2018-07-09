package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadENV is responsible for loading environment variables from
// .env file if it exists in root folder of project
func LoadENV() {
	f := filepath.Join(".", ".env")
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading environment variables")
		os.Exit(1)
	}
}
