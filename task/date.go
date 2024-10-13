package task

import (
  "errors"
  "fmt"
	"time"

  naturaldate "github.com/tj/go-naturaldate"
)

func StringToTime(humanReadable string) (*time.Time, error) {
  if humanReadable == "" {
    return nil, errors.New("empty string")
  }

  reference := time.Now()
  date, err := naturaldate.Parse(humanReadable, reference)
  if err != nil {
    return nil, err
  }
  if date == reference {
    return nil, fmt.Errorf("time formatted poorly: %s", humanReadable)
  }
  return &date, nil
}

