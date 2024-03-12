package task

import (
	"strconv"
)

type Task struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	position    int
}

func (t *Task) SetPriority(priority int) {
	t.Priority = priority
}

func (t *Task) PrettyPriority() string {
	switch t.Priority {
	case 1:
		return "[HIGH]"
	case 3:
		return "[LOW]"
	default:
		return "[MEDIUM]"
	}
}

func (t *Task) PrettyPosition() string {
	return "(" + strconv.Itoa(t.position) + ")"
}
