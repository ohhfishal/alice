package task

import (
	"fmt"
)

type Status uint8

const (
	EVENT = iota
	IN_PROGRESS
	DONE
)

var status_messages = []string{"event", "in-progress", "done"}

func (s Status) MarshalText() (text []byte, err error) {
	if int(s) >= len(status_messages) {
		return []byte{}, fmt.Errorf("unknown status %d", int(s))
	}
	message := status_messages[int(s)]
	return []byte(message), nil
}

func (s *Status) UnmarshalText(text []byte) error {
	message := string(text)
	for i, target := range status_messages {
		if message == target {
			*s = (Status)(i)
			return nil
		}
	}
	return fmt.Errorf("unknown status: %s", message)
}
