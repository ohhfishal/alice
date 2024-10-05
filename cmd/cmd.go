package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

var file string

func init() {
  rootCmd.PersistentFlags().StringVar(&file, "file", "test.txt", "task file")
}

var rootCmd = &cobra.Command {
  Use: "alice",
  Short: "alice is a simple task/event tracker",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello world")
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
