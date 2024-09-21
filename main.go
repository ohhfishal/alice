package main

import (
    "github.com/ohhfishal/alice/event"
    "github.com/ohhfishal/alice/file"
)

func main() {
    test := event.Event{
        Name: "123",
    }
    var group event.EventGroup
    group.Append(&test)
    group.Append(&test)
    group.Append(&test)
    group.Append(&test)

    err := file.ToFile(group, "test.txt")
    if err != nil {
        panic(err)
    }

}
