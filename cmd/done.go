package cmd

import (
	"fmt"

	"github.com/ronthesoul/tago/pkg"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		task, error := pkg.GetTaskWithIndex(index)
		if error != nil {
			fmt.Printf("Error retrieving task: %v\n", error)
			return
		}
		if task.Complete {

			fmt.Printf("Task %d is already marked as done.\n", index)
			return
		}
		if err := pkg.CompleteTask(index); err == nil {
			fmt.Printf("Task %d marked as done successfully.\n", index)
		} else {
			fmt.Printf("Failed to mark task %d as done: %v\n", index, err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
	doneCmd.Flags().IntVarP(&index, "index", "i", -1, "Mark a task as done at the given index")
}
