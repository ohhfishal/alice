package create

import (
  "fmt"
  "strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
  alice "github.com/ohhfishal/alice/api/v1"
  "github.com/ohhfishal/alice/task"
)

func _NewCreateTaskCmd(api alice.API) *cobra.Command {
  var taskCmd = &cobra.Command {
    Use : "task DESCRIPTION",
    Short: "Create a new task",
    Args:  cobra.MinimumNArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
      description := strings.Join(args[:], " ")
      newTask, err := NewTask(description, cmd.Flags())
      if err != nil {
        return fmt.Errorf("task creation: %w", err)
      }
      return api.Create(*newTask)
    },
  }
  return taskCmd
}

func NewTask(description string, args *pflag.FlagSet) (*task.Task, error) {
  t := task.NewTask(description)

  dueString, err := args.GetString("date")
  if err != nil {
    return nil, err
  }

  if dueString == "" {
    return t, nil
  }

  due, err := task.StringToTime(dueString)
  if err != nil {
    return nil, err
  }
  t.Date = *due
  return t, nil

}

