package event

import (
	"encoding/json"
	"io"
	"time"
)

type Option func(*Event) error
type Event struct {
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Due         time.Time `json:"due,omitempty"`
}

func (e *Event) To(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "")
	return encoder.Encode(e)
}

func NewFrom(reader io.Reader) ([]Event, error) {
	var events []Event
	decoder := json.NewDecoder(reader)
	for {
		if !decoder.More() {
			return events, nil
		}
		var newEvent Event
		err := decoder.Decode(&newEvent)
		if err != nil {
			return []Event{}, err
		}
		events = append(events, newEvent)
	}
}

func New(description string, options ...Option) (*Event, error) {
	newEvent := &Event{
		Description: description,
	}

	for _, option := range options {
		err := option(newEvent)
		if err != nil {
			return nil, err
		}
	}
	return newEvent, nil
}

func Due(date time.Time) Option {
	return func(e *Event) error {
		e.Due = date
		return nil
	}
}
