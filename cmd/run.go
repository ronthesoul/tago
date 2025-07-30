/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ronthesoul/tago/pkg"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a specific task",

	Run: func(cmd *cobra.Command, args []string) {
		task, err := pkg.GetTaskWithIndex(index)
		if err != nil {
			fmt.Printf("Error retrieving task: %v\n", err)
			return
		}
		if pkg.CheckIfTaskHasCommand(task) {
			err := pkg.ExecuteCommand(task.Command)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
				return
			}
			fmt.Printf("Task %d executed successfully.\n", index)
		} else {
			fmt.Printf("No command found for task %d.\n", index)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVarP(&index, "index", "i", -1, "Execute the task at the given index")
}
