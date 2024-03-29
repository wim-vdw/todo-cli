package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-cli/internal/task"
	"os"
	"strconv"
)

const editExamples = `  # Update task description for task with ID 1
  todo-cli edit 1 "New task description"`

var editCmd = &cobra.Command{
	Use:     "edit task-id new-description",
	Short:   "Edit task description",
	Aliases: []string{"change", "update"},
	Example: editExamples,
	Args:    checkArgsUpdate,
	Run:     editTask,
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

func editTask(cmd *cobra.Command, args []string) {
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: task-id should be a numeric value.")
		os.Exit(1)
	}
	newDescription := args[1]
	filename := viper.GetString("datafile")
	tasks, err := task.ReadTasks(filename)
	if err != nil {
		fmt.Println("Error reading datafile containing tasks.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	tasks, err = tasks.UpdateTaskDescription(taskID, newDescription)
	if err != nil {
		fmt.Println("Error updating task description.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	fmt.Println("Task description updated with success.")
	err = task.SaveTasks(filename, tasks)
	if err != nil {
		fmt.Println("Error writing datafile containing tasks.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	fmt.Println("Task(s) written to datafile.")
}
