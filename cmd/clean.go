package cmd

import (
	"fmt"

	"github.com/ronthesoul/tago/pkg"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans the task list",
	Run: func(cmd *cobra.Command, args []string) {
		err := pkg.CreateCSVFile()
		if err != nil {
			fmt.Printf("Error creating CSV file: %v\n", err)
			return
		}
		fmt.Println("Task list cleaned successfully.")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
