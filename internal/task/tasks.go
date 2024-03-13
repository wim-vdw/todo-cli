package task

import (
	"encoding/json"
	"errors"
	"os"
)

type Tasks []Task

func (t Tasks) DeleteTask(taskID int) (Tasks, error) {
	if taskID <= 0 {
		return []Task{}, errors.New("task-id can not be negative")
	}
	if taskID > len(t) {
		return []Task{}, errors.New("task-id does not exist")
	}
	taskID -= 1
	t = append(t[:taskID], t[taskID+1:]...)
	return t, nil
}

func (t Tasks) FinishTask(taskID int) (Tasks, error) {
	if taskID <= 0 {
		return []Task{}, errors.New("task-id can not be negative")
	}
	if taskID > len(t) {
		return []Task{}, errors.New("task-id does not exist")
	}
	taskID -= 1
	t[taskID].SetDone()
	return t, nil
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
