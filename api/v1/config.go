package alice

import (
	"fmt"
	"os"

	"github.com/ohhfishal/alice/event"
)

type OutputFormat string

const (
	FORMAT_STRING   OutputFormat = "string"
	FORMAT_MARKDOWN              = "markdown"
	FORMAT_JSON                  = "json"
)

var SUPPORTED_FORMATS = []OutputFormat{
	FORMAT_STRING,
	FORMAT_MARKDOWN,
	FORMAT_JSON,
}

type Config struct {
	Filename      string
	Output        OutputFormat
	CanCreateFile bool `json:"can-create-file,omitempty"`
	ListAll       bool `json:"all,omitempty"`
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
	defer file.Close()
	return newEvent.To(file)
}

type HookAction uint
type Hook func(*event.Event) HookAction

func (c Config) Select(id int, hook Hook) error {
	events, err := c.loadFile()
	if err != nil {
		return err
	}

	if id >= len(events) {
		return fmt.Errorf("invalid id: %d", id)
	}

	// Any error from this point corrupts the file
	file, err := os.Create(c.Filename)
	if err != nil {
		return err
	}

	defer file.Close()
	for i, event := range events {
		// TODO: Use the actions
		// Should be able to do edit with this as well!
		var _ HookAction
		if id == i {
			_ = hook(&event)
		}

		err := event.To(file)
		if err != nil {
			return fmt.Errorf("writing to file: %w", err)
		}
	}
	return nil
}

func (c Config) List() ([]string, error) {
	events, err := c.loadFile()
	if err != nil {
		return []string{}, err
	}

	var lines []string
	for i, cur := range events {
		if !c.ListAll && cur.Status == event.DONE {
			continue
		}
		str, err := c.format(i, cur)
		if err != nil {
			return []string{}, err
		}
		lines = append(lines, str)
	}
	return lines, nil
}

func filter(events []event.Event) []event.Event {
	var retVal []event.Event
	for _, e := range events {
		if e.Status == event.IN_PROGRESS {
			retVal = append(retVal, e)
		}
	}
	return retVal

}

func (c Config) format(id int, e event.Event) (string, error) {
	switch c.Output {
	case FORMAT_STRING:
		// TODO: Make better
		return fmt.Sprintf("[%d] %s", id, e), nil
	case FORMAT_MARKDOWN:
		return "", fmt.Errorf("%s not implemented", c.Output)
	case FORMAT_JSON:
		return "", fmt.Errorf("%s not implemented", c.Output)
	default:
		return "", fmt.Errorf("invalid format: %s", c.Output)
	}
}

func (c Config) loadFile() ([]event.Event, error) {
	file, err := os.Open(c.Filename)
	if err != nil {
		return []event.Event{}, err
	}

	return event.NewFrom(file)
}
