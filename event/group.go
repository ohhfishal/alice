package event

import (
    "encoding/json"
    "io"
    "os"
)

type EventGroup struct {
    events []*Event
}

func NewFromFile(path string) (*EventGroup, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    var group EventGroup 
    _, err = group.Load(file)
    if err != nil {
        return nil, err
    }
    return &group, nil
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

