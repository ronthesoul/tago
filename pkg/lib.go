package pkg

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func CreateCSVFile() error {

	file, err := os.Create(CSVFile)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()
	return nil
}

func WriteToCSVFile(records []string) error {
	file, err := os.OpenFile(CSVFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func ReadAllTasksFromCsv(CSVFile string) error {

	file, err := os.Open(CSVFile)
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

func ReadPendingTasksFromCsv(CSVFile string) error {
	file, err := os.Open(CSVFile)
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
		complete, err := strconv.ParseBool(row[5])
		if err != nil {
			return fmt.Errorf("failed to parse complete status: %w", err)
		}
		if !complete {
			err = table.Append(row)
			if err != nil {
				return fmt.Errorf("failed to append row to table: %w", err)
			}

		}
	}
	err = table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}
	return nil
}

func ReadCommandTasksFromCsv(CSVFile string) error {
	file, err := os.Open(CSVFile)
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
		command := row[2]
		if command != "" && command != " " {
			err = table.Append(row)
			if err != nil {
				return fmt.Errorf("failed to append row to table: %w", err)
			}

		}
	}
	err = table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}
	return nil
}

func RemoveTaskFromCSV(indexToRemove int) error {

	file, err := os.Open(CSVFile)
	if err != nil {
		return fmt.Errorf("failed to to open csv file: %w", err)

	}

	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}
	if indexToRemove < 1 || indexToRemove > len(tasks) {
		return fmt.Errorf("index out of range: %d", indexToRemove)
	}
	defer file.Close()

	tasks = append(tasks[:indexToRemove-1], tasks[indexToRemove+1:]...)
	for i := range tasks {
		tasks[i][0] = strconv.Itoa(i + 1)
	}

	file, _ = os.Create(CSVFile)
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
	file, err := os.Open(CSVFile)
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

func ExecuteCommand(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error executing command '%s': %v\nOutput: %s", command, err, output)
	}
	fmt.Printf("Command '%s' executed successfully. \noutput: %s\n", command, output)
	return nil
}

func CheckIfTaskHasCommand(task Task) bool {
	if task.Command == "" || task.Command == " " {
		fmt.Printf("Task %d has no command to execute.\n", task.ID)
		return false
	}
	return true
}

func GetTaskWithIndex(index int) (Task, error) {
	file, err := os.Open(CSVFile)
	if err != nil {
		return Task{}, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()
	if err != nil {
		return Task{}, fmt.Errorf("failed to read CSV file: %w", err)
	}
	if index < 1 || index > len(tasks) {
		return Task{}, fmt.Errorf("index out of range: %d", index)
	}
	taskData := tasks[index-1]
	if len(taskData) < 6 {
		return Task{}, fmt.Errorf("task data is incomplete for index %d", index)
	}
	id, err := strconv.Atoi(taskData[0])
	if err != nil {
		return Task{}, fmt.Errorf("invalid ID for task at index %d: %w", index, err)
	}
	complete, err := strconv.ParseBool(taskData[5])
	if err != nil {
		return Task{}, fmt.Errorf("invalid complete status for task at index %d: %w", index, err)
	}
	return Task{
		ID:       id,
		Name:     taskData[1],
		Command:  taskData[2],
		Time:     taskData[3],
		Desc:     taskData[4],
		Complete: complete,
	}, nil
}

func CompleteTask(index int) error {
	file, err := os.Open(CSVFile)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	if index < 1 || index > len(tasks) {
		return fmt.Errorf("index out of range: %d", index)
	}

	tasks[index-1][5] = "true"

	file, err = os.Create(CSVFile)
	if err != nil {
		return fmt.Errorf("failed to open CSV file for writing: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(tasks)
	if err != nil {
		return fmt.Errorf("failed to write tasks to CSV file: %w", err)
	}

	return nil
}
