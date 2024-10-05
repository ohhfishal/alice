package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
  Use:   "create TASK [flags]",
  Short: "Create a new task",
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("creating...")
    fmt.Println(args)
  },
}
