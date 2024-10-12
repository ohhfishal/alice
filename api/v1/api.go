package alice

import (
	"github.com/ohhfishal/alice/task"
)

type API interface {
	Create(task task.Task) error
	Delete(query task.Query) error
	Filter(query task.Query) error
}
