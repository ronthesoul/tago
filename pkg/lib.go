package pkg

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/olekukonko/tablewriter"
)

func CreateJsonFile() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	csvFile := os.Getenv("CSV_FILE")
	if csvFile == "" {
		return fmt.Errorf("CSV_FILE environment variable is not set")
	}
	file, err := os.Create(csvFile)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
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

func ReadTasksFromCsv(csvFile string) error {

	file, err := os.Open(csvFile)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Name", "Command", "Time", "Description", "Complete"})
	for _, row := range records {
		table.Append(row)
	}
	table.Render()
	return nil
}
