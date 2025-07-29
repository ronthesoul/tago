package pkg

import (
	"encoding/csv"
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

func WriteToCSVFile(records []string) error {
	csvFile := os.Getenv("CSV_FILE")
	file, err := os.OpenFile(csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(records)
	if err != nil {
		return fmt.Errorf("failed to write to CSV file: %w", err)
	}
	return nil
}
