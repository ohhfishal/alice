package alice

import (
	"errors"
	"github.com/ohhfishal/alice/file"
	"github.com/ohhfishal/alice/task"
)

type Alice struct {
	Filepath string
}

func NewAPI(filepath string) API {
  return &Alice {
    Filepath: filepath,
  }
}

func (c Alice) Create(task task.Task) error {
	return file.AppendToFile(task, c.Filepath)
}

func (c Alice) Delete(query task.Query) error {
	return errors.New("DELETE not implemented")
}

func (c Alice) Filter(query task.Query) error {
	return errors.New("FILTER not implemented")
}
