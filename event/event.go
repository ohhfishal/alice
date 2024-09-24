package event

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Event struct {
	Description string     `json:"description,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	Status      Status     `json:"status,omitempty"`
	Children    []Event    `json:"children,omitempty"`
}

func NewTask(description string) *Event {
	return &Event{
		Description: description,
		Status:      IN_PROGRESS,
	}
}

func NewEvent(description string) *Event {
	return &Event{
		Description: description,
	}
}

func (event Event) Save(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	if encoder == nil {
		return fmt.Errorf("failed to create encoder")
	}
	// Disable indentation
	encoder.SetIndent("", "")
	return encoder.Encode(event)
}

func (e Event) IsDue() bool {
	return e.Date.After(time.Now())
}

func (e Event) Complete() error {
	e.Status = DONE

	// Note this behavior may not always be wanted
	for _, task := range e.Children {
		task.Complete()
	}
	return nil
}
