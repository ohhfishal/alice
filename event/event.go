package event

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Event struct {
	Name        string     `json:"name,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	Description string     `json:"description,omitempty"`
	Head        *Task      `json:"head,omitempty"`
	Tasks       []Task     `json:"tasks,omitempty"`
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
