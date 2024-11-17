package alice

import (
	"github.com/ohhfishal/alice/event"
)

type API interface {
	Create(event.Event) error
	List() error
}
