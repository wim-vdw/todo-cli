package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-cli/internal/task"
)

const doneExamples = `  # Mark task with ID 1 as done
  todo-cli done 1`

var doneCmd = &cobra.Command{
	Use:     "done task-id",
	Short:   "Mark a task as done in the To-Do list",
	Aliases: []string{"finish", "ok"},
	Example: doneExamples,
	Args:    checkArgsDone,
	Run:     doneTask,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func checkArgsDone(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("specify exactly 1 argument containing task-id")
	}
	return nil
}

func doneTask(cmd *cobra.Command, args []string) {
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
	tasks, err = tasks.FinishTask(taskID)
	if err != nil {
		fmt.Println("Error setting task to done.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	fmt.Println("Task finished with success.")
	err = task.SaveTasks(filename, tasks)
	if err != nil {
		fmt.Println("Error writing datafile containing tasks.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	fmt.Println("Task(s) written to datafile.")
}
