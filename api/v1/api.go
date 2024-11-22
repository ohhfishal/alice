package alice

import (
	"github.com/ohhfishal/alice/event"
)

type API interface {
	Create(event.Event) error
	List() ([]string, error)
	Mark(int, event.Status) error
}
