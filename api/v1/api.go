package v1

import (
  "github.com/ohhfishal/alice/event"
)

type API interface {
  Add(event event.Event) error
  Delete(query event.Event) error
  Filter(query event.Event) error
}

