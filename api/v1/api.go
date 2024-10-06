package v1

import (
	"github.com/ohhfishal/alice/task"
)

type API interface {
	Add(task task.Task) error
	Delete(query task.Query) error
	Filter(query task.Query) error
}
