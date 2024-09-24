package event

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

func NewTestEvents() EventGroup {
	now := time.Now()
	a := Event{
		Date:        &now,
		Description: "full event",
	}
	b := Event{
		Date:        &now,
		Description: "full event",
	}
	c := Event{}
	d := Event{
		Description: "abc",
	}
	e := Event{
		Description: "test",
	}
	f := Event{}

	var group EventGroup
	group.Append(&a)
	group.Append(&b)
	group.Append(&c)
	group.Append(&d)
	group.Append(&e)
	group.Append(&f)
	return group

}

func TestReadAndWrite(t *testing.T) {
	var events, results EventGroup
	writer := NewMockWriter()
	events = NewTestEvents()

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
