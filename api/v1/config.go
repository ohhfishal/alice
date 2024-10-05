package v1

import (
  "errors"
  "github.com/ohhfishal/alice/event"
  "github.com/ohhfishal/alice/file"
)

type Config struct {
  Filepath string
}


func (c Config) Add(event event.Event) error {
  return file.AppendToFile(event, c.Filepath)
}

func (c Config) Delete(query event.Event) error {
  return errors.New("DELETE not implemented")
}

func (c Config) Filter(query event.Event) error {
  return errors.New("FILTER not implemented")
}
