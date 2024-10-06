package task

import (
	"bytes"
	"io"
	"testing"
	"time"
)

func NewMockWriter() io.Writer {
	var buffer bytes.Buffer
	return (io.Writer)(&buffer)
}

func NewMockReader(content *bytes.Buffer) io.Reader {
	return (io.Reader)(content)
}

func NewTestTasks() TaskGroup {
	now := time.Now()
	a := Task{
		Date:        &now,
		Description: "full event",
	}
	b := Task{
		Date:        &now,
		Description: "full event",
	}
	c := Task{}
	d := Task{
		Description: "abc",
	}
	e := Task{
		Description: "test",
	}
	f := Task{}

	var group TaskGroup
	group.Append(&a)
	group.Append(&b)
	group.Append(&c)
	group.Append(&d)
	group.Append(&e)
	group.Append(&f)
	return group

}

func TestReadAndWrite(t *testing.T) {
	var events, results TaskGroup
	writer := NewMockWriter()
	events = NewTestTasks()

	err := events.Save(writer)
	if err != nil {
		t.Error(err)
	}

	reader := NewMockReader(writer.(*bytes.Buffer))
	count, err := results.Load(reader)
	if err != nil {
		t.Error(err)
	}
	if count != len(events.events) {
		t.Errorf("%d events read. Expected %d", count, len(events.events))
	}
	// TODO: Validate that all the events are what we expect

}
