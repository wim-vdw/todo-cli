package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-cli/internal/task"
)

var priority int

const addExamples = `  # Add one task
  todo-cli add "Go to store"

  # Add multiple tasks
  todo-cli add "Go to store" "Visit family"

  # Add multiple tasks all with the same priority
  todo-cli add "Go to store" "Visit family" --priority 1`

var addCmd = &cobra.Command{
	Use:     "add tasks...",
	Short:   "Add tasks to the To-Do list",
	Aliases: []string{"create", "new"},
	Example: addExamples,
	Run:     addTask,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority.")
}

func addTask(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: specify at least 1 task to add.")
		os.Exit(1)
	}
	filename := viper.GetString("datafile")
	tasks, err := task.ReadTasks(filename)
	if err != nil {
		fmt.Println("Error reading datafile containing tasks.")
		fmt.Println("Error message ->", err)
		fmt.Println("Tasks will be written to a new file.")
	}
	for _, x := range args {
		item := task.Task{Description: x}
		item.SetPriority(priority)
		tasks = append(tasks, item)
	}
	fmt.Println("Task(s) added with success.")
	err = task.SaveTasks(filename, tasks)
	if err != nil {
		fmt.Println("Error writing datafile containing tasks.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	fmt.Println("Task(s) written to datafile.")
}
