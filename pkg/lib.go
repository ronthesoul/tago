package lib

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func CreateJsonFile() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	jsonFile := os.Getenv("JSON_FILE")
	if jsonFile == "" {
		return fmt.Errorf("JSON_FILE environment variable is not set")
	}
	file, err := os.Create(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to create JSON file: %w", err)
	}
	defer file.Close()
	return nil
}
