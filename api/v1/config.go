package alice

import (
	"errors"
	"os"

	"github.com/ohhfishal/alice/event"
)

type Config struct {
	Filename      string
	CanCreateFile bool `json:"can-create-file,omitempty"`
}

func (c Config) Create(newEvent event.Event) error {
	flags := os.O_WRONLY | os.O_APPEND
	if c.CanCreateFile {
		flags = flags | os.O_CREATE
	}

	file, err := os.OpenFile(c.Filename, flags, 0777)
	if err != nil {
		return err
	}
	return newEvent.To(file)
}

func (c Config) List() error {
	return errors.New("(c Config) List() not implemented")
}
