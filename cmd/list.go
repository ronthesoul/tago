/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"tago/pkg"

	"github.com/spf13/cobra"
)

var (
	all          bool
	shellcommand bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List spesified tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			err := pkg.ReadAllTasksFromCsv(os.Getenv("CSV_FILE"))
			if err != nil {
				fmt.Printf("Error reading tasks from CSV file: %v\n", err)
				return
			}
		} else if shellcommand {
			err := pkg.ReadCommandTasksFromCsv(os.Getenv("CSV_FILE"))
			if err != nil {
				fmt.Printf("Error reading done tasks from CSV file: %v\n", err)
				return
			}
		} else {
			err := pkg.ReadPendingTasksFromCsv(os.Getenv("CSV_FILE"))
			if err != nil {
				fmt.Printf("Error reading tasks from CSV file: %v\n", err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "List all tasks")
	listCmd.Flags().BoolVarP(&shellcommand, "commands", "c", false, "List only done tasks")
}
