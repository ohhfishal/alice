package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCommand() *cobra.Command {
  var rootCmd = &cobra.Command{
    Use:   "alice",
    Short: "alice is a simple task/event tracker",
    Run: func(cmd *cobra.Command, args []string) {
      cmd.Help()
    },
  }

	rootCmd.PersistentFlags().StringP("file", "f", "test.txt", "task file")
  viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))

  rootCmd.AddCommand(NewCreateCommand())
  return rootCmd
}


func Execute() {
	if err := NewRootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
