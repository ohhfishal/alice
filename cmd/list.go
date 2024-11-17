package cmd

import (
  "fmt"
	alice "github.com/ohhfishal/alice/api/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)


func ListFlags() *pflag.FlagSet {
	flags := pflag.NewFlagSet("list", pflag.ContinueOnError)
  msg := fmt.Sprintf("output format (%v)", alice.SUPPORTED_FORMATS)
  flags.StringP("output", "o", "string", msg)
	return flags
}

func NewList(args []string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:          "list",
    Aliases:      []string{"ls"},
		Short:        "List all tasks",
		RunE:         List,
		Args:         cobra.NoArgs,
		SilenceUsage: true,
	}
	cmd.SetArgs(args)
	cmd.Flags().AddFlagSet(ListFlags())
  
	return cmd
}

func List(cmd *cobra.Command, args []string) error {
	// Use Viper to unmarshal everything
	vConfig := viper.New()
	vConfig.BindPFlags(cmd.Flags())

	config, err := NewConfig(vConfig)
	if err != nil {
		return err
	}
  return config.List()
}
