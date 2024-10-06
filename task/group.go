package task

import (
	"encoding/json"
	"io"
	"os"
)

type TaskGroup struct {
	events []*Task
}

func NewFromFile(path string) (*TaskGroup, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var group TaskGroup
	_, err = group.Load(file)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (group *TaskGroup) Append(event *Task) {
	group.events = append(group.events, event)
}

func (group TaskGroup) Save(writer io.Writer) error {
	for i := 0; i < len(group.events); i++ {
		if err := group.events[i].Save(writer); err != nil {
			return err
		}
	}
	return nil
}

func (group *TaskGroup) Load(reader io.Reader) (int, error) {
	initialLength := len(group.events)
	for {
		var newTask Task
		decoder := json.NewDecoder(reader)

		if err := decoder.Decode(&newTask); err == io.EOF {
			return len(group.events) - initialLength, nil
		} else if err != nil {
			return len(group.events) - initialLength, err
		}
		reader = decoder.Buffered()
		group.Append(&newTask)
	}
}
