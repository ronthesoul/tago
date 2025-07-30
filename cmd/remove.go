/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tago/pkg"

	"github.com/spf13/cobra"
)

var index int
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Number of the task to remove",

	Run: func(cmd *cobra.Command, args []string) {

		if index < 0 {
			fmt.Println("Please provide a valid task number to remove.")
			return
		}
		if index == 0 {
			fmt.Println("Cannot remove the first task. Please provide a valid task number.")
			return
		}

		err := pkg.RemoveTaskFromCSV(index)
		if err != nil {
			fmt.Printf("Error removing task: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().IntVarP(&index, "index", "i", -1, "Number of the task to remove")
}
