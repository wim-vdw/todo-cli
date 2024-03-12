package task

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

type Task struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	position    int
}

type Tasks []Task

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

func ReadTasks(filename string) (Tasks, error) {
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

func SaveTasks(filename string, tasks Tasks) error {
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

func (t Tasks) DeleteTask(taskID int) (Tasks, error) {
	if taskID > len(t) {
		return []Task{}, errors.New("task-id does not exist")
	}
	taskID -= 1
	t = append(t[:taskID], t[taskID+1:]...)
	return t, nil
}
