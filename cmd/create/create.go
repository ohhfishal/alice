package create

import (
	"github.com/spf13/cobra"
  alice "github.com/ohhfishal/alice/api/v1"
)

func NewCreateCommand(api alice.API) *cobra.Command {
  var createCmd = &cobra.Command{
    Use:   "create",
    Short: "Create a new task or event",
    Run: func(cmd *cobra.Command, args []string) {
      cmd.Help()
    },
  }
  createCmd.PersistentFlags().StringP("date", "d", "", "due date")

  createCmd.AddCommand(_NewCreateTaskCmd(api))
  return createCmd
}


