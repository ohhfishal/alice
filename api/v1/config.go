package v1

import (
	"errors"
	"github.com/ohhfishal/alice/file"
	"github.com/ohhfishal/alice/task"
)

type Config struct {
	Filepath string
}

func (c Config) Add(task task.Task) error {
	return file.AppendToFile(task, c.Filepath)
}

func (c Config) Delete(query task.Query) error {
	return errors.New("DELETE not implemented")
}

func (c Config) Filter(query task.Query) error {
	return errors.New("FILTER not implemented")
}
