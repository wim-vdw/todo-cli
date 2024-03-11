package task

import (
	"encoding/json"
	"os"
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

func ReadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []Task{}, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return []Task{}, err
	}
	for i, _ := range tasks {
		tasks[i].position = i + 1
	}
	return tasks, nil
}

func SaveTasks(filename string, tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
