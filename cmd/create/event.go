package create

import (
  "fmt"
  "strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
  alice "github.com/ohhfishal/alice/api/v1"
  "github.com/ohhfishal/alice/task"
)

func NewEventCmd(api alice.API) *cobra.Command {
  var taskCmd = &cobra.Command {
    Use : "event DATE DESCRIPTION",
    Short: "Create a new event",
    Args:  cobra.MinimumNArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
      newTask, err := NewEvent(args, cmd.Flags())
      if err != nil {
        return fmt.Errorf("event creation: %w", err)
      }
      return api.Create(*newTask)
    },
  }
  return taskCmd
}

func NewEvent(args []string, flags *pflag.FlagSet) (*task.Task, error) {
  dueDateRaw := args[0]
  dueDate, err := task.StringToTime(dueDateRaw)
  if err != nil {
    return nil, err
  }
  description := strings.Join(args[1:], " ")
  return task.NewEvent(*dueDate, description), nil
}

