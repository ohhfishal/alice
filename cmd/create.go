package cmd

import (
	"fmt"

  "github.com/ohhfishal/alice/task"
	"github.com/spf13/cobra"
)

var newTask task.Task
var due date

func init() {
  createCmd.Flags().VarP(&due, "date", "d", "due date")
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create task DESCRIPTION [flags]",
	Short: "Create a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("creating...")
		fmt.Println(args)
		fmt.Println("using...")
		fmt.Println(c)

	},
}
