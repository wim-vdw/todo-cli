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
