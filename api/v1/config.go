package alice

import (
  "fmt"
	"os"

	"github.com/ohhfishal/alice/event"
)

type OutputFormat string

const (
  FORMAT_STRING OutputFormat = "string"
  FORMAT_MARKDOWN = "markdown"
  FORMAT_JSON = "json"

)

var SUPPORTED_FORMATS = []OutputFormat {
  FORMAT_STRING,
  FORMAT_MARKDOWN,
  FORMAT_JSON,
}

type Config struct {
	Filename      string
  Output        OutputFormat
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
  events, err := c.loadFile()
  if err != nil {
    return err
  }

  for i, cur := range events {
    str, err := c.format(i, cur)
    if err != nil {
      return err
    }

    fmt.Printf("%s\n", str)
  }
	return nil
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
