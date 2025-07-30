package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ronthesoul/tago/pkg"

	"github.com/spf13/cobra"
)

var id int = 3
var desc string
var name string
var command string
var complete bool = false
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := pkg.SetIndex()
		if err != nil {
			log.Fatalf("Error getting index: %v", err)
		}
		task := pkg.Task{
			ID:       id,
			Name:     name,
			Command:  command,
			Time:     time.Now().Format(time.RFC1123),
			Desc:     desc,
			Complete: complete}
		record := []string{
			strconv.Itoa(task.ID),
			task.Name,
			task.Command,
			task.Time,
			task.Desc,
			strconv.FormatBool(task.Complete),
		}
		err = pkg.WriteToCSVFile(record)
		if err != nil {
			log.Fatalf("Error writing to CSV file: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&desc, "desc", "d", " ", "Description of the task")
	addCmd.Flags().StringVarP(&name, "name", "n", "task"+fmt.Sprintf("%d", id), "Name of the task")
	addCmd.Flags().StringVarP(&command, "command", "c", " ", "Command to run")
}
