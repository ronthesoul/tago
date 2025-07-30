package pkg

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/olekukonko/tablewriter"
)

func CreateCSVFile() error {
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
		err = table.Append(row)
		if err != nil {
			return fmt.Errorf("failed to append row to table: %w", err)
		}

	}
	err = table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}
	return nil
}

func RemoveTaskFromCSV(indexToRemove int) error {

	file, err := os.Open(os.Getenv("CSV_FILE"))
	if err != nil {
		return fmt.Errorf("failed to to open csv file: %w", err)

	}

	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()
	defer file.Close()

	tasks = append(tasks[:indexToRemove-1], tasks[indexToRemove+1:]...)
	for i := range tasks {
		tasks[i][0] = strconv.Itoa(i + 1)
	}

	file, _ = os.Create(os.Getenv("CSV_FILE"))
	writer := csv.NewWriter(file)
	err = writer.WriteAll(tasks)
	if err != nil {
		return fmt.Errorf("failed to write updated tasks to CSV file: %w", err)
	}
	defer file.Close()
	if err != nil {
		return fmt.Errorf("failed to to open csv file: %w", err)
	}
	return nil

}

func SetIndex() (int, error) {
	file, err := os.Open(os.Getenv("CSV_FILE"))
	if err != nil {
		return 0, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	currentTasks, err := reader.ReadAll()
	if err != nil {
		return 0, fmt.Errorf("failed to read CSV file: %w", err)
	}
	if len(currentTasks) == 0 {
		return 1, nil
	}
	return len(currentTasks) + 1, nil
}
