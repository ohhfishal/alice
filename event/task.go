package event


import (
    "time"
)

type Task struct {
    Description string `json:"description,omitempty"`
    Status Status `json:"status,omitempty"`
    Date time.Time `json:"date,omitempty"`
    Subtasks []Task `json:"subtasks,omitempty"`
}

func (t Task) IsDue() bool {
    return t.Date.After(time.Now())
}

func (t *Task) Complete() {
    t.Status = S_DONE 
    
    // Note this behavior may not always be wanted
    for _, task := range t.Subtasks {
        task.Complete()
    }
}
