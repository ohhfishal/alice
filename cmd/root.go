package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
  "github.com/ohhfishal/alice/cmd/create"
  alice "github.com/ohhfishal/alice/api/v1"
)

func NewRootCommand() (*cobra.Command, error) {
  var rootCmd = &cobra.Command{
    Use:   "alice",
    Short: "alice is a simple task/event tracker",
    Run: func(cmd *cobra.Command, args []string) {
      cmd.Help()
    },
  }

	rootCmd.PersistentFlags().StringP("file", "f", "test.txt", "task file")
  viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))

  api, err := NewApi(rootCmd.PersistentFlags())
  if err != nil {
    return nil, fmt.Errorf("root creation: %w", err)
  }
  rootCmd.AddCommand(create.NewCreateCommand(api))
  return rootCmd, nil
}

func NewApi(flags *pflag.FlagSet)  (alice.API, error) {
  filepath, err := flags.GetString("file")
  if err != nil {
    return nil, fmt.Errorf("filepath not found: %w", err)
  }
  api := alice.NewAPI(filepath)
  return api, nil
}

func Execute() {
  root, err := NewRootCommand()
  if err != nil {
		fmt.Println(err)
		os.Exit(1)
  }

	if root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
