package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Execute() {
	if err := NewRoot(os.Args[1:]).Execute(); err != nil {
		os.Exit(1)
	}
}

func NewRoot(args []string) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:          "alice DESCRIPTION",
		Short:        "Alice is a task todo list tool",
		RunE:         Create,
		Args:         cobra.MinimumNArgs(1),
		SilenceUsage: true,
	}
	rootCmd.PersistentFlags().AddFlagSet(RootFlags())
	rootCmd.Flags().AddFlagSet(CreateFlags())
	rootCmd.SetArgs(args)

  rootCmd.AddCommand(NewList(args))
	return rootCmd
}

func RootFlags() *pflag.FlagSet {
	flags := pflag.NewFlagSet("alice", pflag.ContinueOnError)
	flags.StringP("filename", "f", "test.json", "file to contain tasks")
	return flags
}
