package cmd

import (
	"strings"

	"github.com/ohhfishal/alice/event"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func CreateFlags() *pflag.FlagSet {
	flags := pflag.NewFlagSet("create", pflag.ContinueOnError)
	flags.StringArray("tags", []string{}, "tags to add to the new task")
	flags.BoolP("can-create-file", "c", false, "enable creating of missing files")
	return flags
}

func Create(cmd *cobra.Command, args []string) error {
	// Use Viper to unmarshal everything
	vConfig := viper.New()
	vConfig.BindPFlags(cmd.Flags())

	config, err := NewConfig(vConfig)
	if err != nil {
		return err
	}

	options, err := CreateOptions(vConfig)
	if err != nil {
		return err
	}

	newEvent, err := event.New(strings.Join(args, " "), options...)
	if err != nil {
		return err
	}
	return config.Create(*newEvent)
}

func CreateOptions(config *viper.Viper) ([]event.Option, error) {
	var options []event.Option
	// TODO: Translate viper info to options
	//       - [ ] Due date
	return options, nil
}
