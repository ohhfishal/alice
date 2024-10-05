package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Alice",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("alice version 0.1.0")
  },
}
