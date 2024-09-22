package main

import (
    "time"
    . "github.com/ohhfishal/alice/event"
    "github.com/ohhfishal/alice/file"
)


func NewTestEvents() EventGroup {
    now := time.Now()
    a := Event{
        Name: "abc",
        Date: &now,
        Description: "full event",
    }
    b := Event{
        Name: "abc",
        Date: &now,
        Description: "full event",
    }
    c := Event{}
    var task Task
    d := Event{
        Name: "abc",
        Head: &task,
    }
    task2 := Task {
        Description: "Marshaling",
    }
    e := Event{
        Name: "abc",
        Head: &task2,
    }

    var group EventGroup
    group.Append(&a)
    group.Append(&b)
    group.Append(&c)
    group.Append(&d)
    group.Append(&e)
    return group
}

func main() {
    group := NewTestEvents()

    err := file.ToFile(group, "test.txt")
    if err != nil {
        panic(err)
    }

}
