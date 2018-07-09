package env

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Load is responsible for loading environment variables from
// .env file if it exists in root folder of project
func Load() {
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
