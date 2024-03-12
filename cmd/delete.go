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

const deleteExamples = `  # Delete task with ID 1
  todo-cli delete 1`

var deleteCmd = &cobra.Command{
	Use:     "delete task-id",
	Short:   "Delete task from the To-Do list",
	Aliases: []string{"del", "rm", "remove"},
	Example: deleteExamples,
	Args:    checkArgs,
	Run:     deleteTask,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func checkArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("specify exactly 1 argument containing task-id")
	}
	return nil
}

func deleteTask(cmd *cobra.Command, args []string) {
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: task-id should be a numeric value.")
		os.Exit(1)
	}
	filename := viper.GetString("datafile")
	tasks, err := task.ReadTasks(filename)
	if err != nil {
		fmt.Println("Error reading datafile containing tasks.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	tasks, err = tasks.DeleteTask(taskID)
	if err != nil {
		fmt.Println("Error deleting task from the To-Do list.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	fmt.Println("Task deleted with success.")
	err = task.SaveTasks(filename, tasks)
	if err != nil {
		fmt.Println("Error writing datafile containing tasks.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	fmt.Println("Task(s) written to datafile.")
}
