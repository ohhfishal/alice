package event

type Status uint8

const (
	S_EMPTY = iota
	S_DONE
	S_OVERDUE
)
