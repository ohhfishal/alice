package event

import (
	"fmt"
)

type Status uint

const (
	IN_PROGRESS = iota
	DONE
)

func (s Status) Valid() error {
	switch {
	case s < IN_PROGRESS || s > DONE:
		return fmt.Errorf("%d is not a valid Status", s)
	default:
		return nil
	}
}
