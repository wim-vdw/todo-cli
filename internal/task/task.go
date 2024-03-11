package task

import (
	"encoding/json"
	"os"
)

type Task struct {
	Description string `json:"description"`
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
	return tasks, nil
}
