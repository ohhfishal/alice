package cmd

import (
  "fmt"
  "strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
  alice "github.com/ohhfishal/alice/api/v1"
  "github.com/ohhfishal/alice/task"
)

func NewCreateCommand() *cobra.Command {
  var createCmd = &cobra.Command{
    Use:   "create",
    Short: "Create a new task or event",
    Run: func(cmd *cobra.Command, args []string) {
      cmd.Help()
    },
  }
  createCmd.PersistentFlags().StringP("date", "d", "", "due date")

  createCmd.AddCommand(_NewCreateTaskCmd())
  return createCmd
}

func NewTask(description string, args *pflag.FlagSet) (*task.Task, error) {
  t := task.NewTask(description)

  dueString, err := args.GetString("date")
  if err != nil {
    return nil, err
  }

  due, err := task.StringToTime(dueString)
  if err != nil {
    return nil, err
  }
  t.Date = *due
  return t, nil

}

func NewApi(flags *pflag.FlagSet)  (alice.API, error) {
  filepath, err := flags.GetString("file")
  if err != nil {
    return nil, fmt.Errorf("filepath not found: %w", err)
  }
  api := alice.NewAPI(filepath)
  return api, nil

}

func _NewCreateTaskCmd() *cobra.Command {
  var taskCmd = &cobra.Command {
    Use : "task DESCRIPTION",
    Short: "Create a new task",
    Args:  cobra.MinimumNArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
      api, err := NewApi(cmd.Flags())
      if err != nil {
        return fmt.Errorf("api creation: %w", err)
      }

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
