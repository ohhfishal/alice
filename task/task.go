package task

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Task struct {
	Description  string     `json:"description,omitempty"`
	Date         time.Time `json:"date,omitempty"`
	Status       Status     `json:"status,omitempty"`
	Children     []Task     `json:"children,omitempty"`
	ExpiresOnDue bool       `json:"expires_on_due"`
}

func NewTask(description string) *Task {
	return &Task{
		Description: description,
		Status:      IN_PROGRESS,
	}
}

func NewEvent(description string) *Task {
	return &Task{
		Description: description,
	}
}

func (event Task) Save(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	if encoder == nil {
		return fmt.Errorf("failed to create encoder")
	}
	// Disable indentation
	encoder.SetIndent("", "")
	return encoder.Encode(event)
}

func (e Task) IsDue() bool {
	return e.Date.After(time.Now())
}

func (e Task) Complete() error {
	e.Status = DONE

	// Note this behavior may not always be wanted
	for _, task := range e.Children {
		task.Complete()
	}
	return nil
}
