package main

import (
	"github.com/ohhfishal/alice/event"
	"github.com/ohhfishal/alice/file"
)

func main() {
	var manager event.EventGroup

	task := event.NewTask("new task")
	e := event.NewEvent("new event")

	manager.Append(task)
	manager.Append(e)

	err := file.ToFile(manager, "test.txt")
	if err != nil {
		panic(err)
	}
}
