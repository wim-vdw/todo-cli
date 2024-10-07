package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-client/task"
)

const editExamples = `  # Update task description for task with ID 1
  todo-cli edit 1 "New task description"`

var editCmd = &cobra.Command{
	Use:     "edit task-id new-description",
	Short:   "Edit task description",
	Aliases: []string{"change", "update"},
	Example: editExamples,
	Args:    checkArgsUpdate,
	RunE:    editTask,
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func checkArgsUpdate(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("specify exactly 2 arguments containing task-id and the new task description")
	}
	return nil
}

func editTask(cmd *cobra.Command, args []string) error {
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("task-id should be a numeric value")
	}
	newDescription := args[1]
	filename := viper.GetString("datafile")
	taskClient := task.Client{}
	err = taskClient.ReadTasks(filename)
	if err != nil {
		return fmt.Errorf("could not read datafile '%s'", filename)
	}
	err = taskClient.UpdateTaskDescription(taskID, newDescription)
	if err != nil {
		return err
	}
	err = taskClient.SaveTasks(filename)
	if err != nil {
		return fmt.Errorf("could not save datafile '%s'", filename)
	}
	fmt.Println("Task description updated with success.")
	fmt.Println("Task(s) written to datafile.")
	return nil
}
