package cmd

import (
	"fmt"
	alice "github.com/ohhfishal/alice/api/v1"
	"github.com/ohhfishal/alice/event"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strconv"
)

func DoFlags() *pflag.FlagSet {
	flags := pflag.NewFlagSet("do", pflag.ContinueOnError)
	return flags
}

func NewDo(args []string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:          "do ID",
		Short:        "Mark a task as done",
		RunE:         Do,
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
	}
	cmd.SetArgs(args)
	cmd.Flags().AddFlagSet(DoFlags())

	return cmd
}

func Do(cmd *cobra.Command, args []string) error {
	// Use Viper to unmarshal everything
	vConfig := viper.New()
	vConfig.BindPFlags(cmd.Flags())

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	config, err := NewConfig(vConfig)
	if err != nil {
		return err
	}
	return config.Select(id, DoHook)
}

func DoHook(e *event.Event) alice.HookAction {
	var noop alice.HookAction
	e.Status = event.DONE
	return noop
}
