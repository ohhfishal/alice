package event

import (
    "encoding/json"
    "fmt"
    "io"
    "time"
)



type Event struct {
    Name string `json:"name.omitempty"`
    Date time.Time `json:"date,omitempty"`
    Description string `json:"description,omitempty"`
}

func (event Event) Save(writer io.Writer) error {
    encoder := json.NewEncoder(writer)
    if encoder == nil {
        return fmt.Errorf("failed to create encoder")
    }
    // Disable indentation
    encoder.SetIndent("", "",)
    return encoder.Encode(event)
}

type EventGroup struct {
    events []*Event
}

func (group *EventGroup) Append(event *Event) {
    group.events = append(group.events, event)
}

func (group EventGroup) Save(writer io.Writer) error {
    for i := 0; i < len(group.events); i++ {
        if err := group.events[i].Save(writer); err != nil {
            return err
        }
    }
    return nil
}

func (group *EventGroup) Load(reader io.Reader) (int, error) {
    initialLength := len(group.events)
    for {
        var newEvent Event
        decoder := json.NewDecoder(reader)

        if err := decoder.Decode(&newEvent); err == io.EOF {
            return len(group.events) - initialLength, nil
        } else if err != nil {
            return len(group.events) - initialLength, err
        }
        reader = decoder.Buffered()
        group.Append(&newEvent)
    }
}


