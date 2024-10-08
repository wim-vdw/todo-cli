package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-client/task"
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
	RunE:    addTask,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority.")
}

func addTask(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("specify at least 1 task to add")
	}
	filename := viper.GetString("datafile")
	taskClient := &task.Client{
		Filename: filename,
	}
	err := taskClient.ReadTasks()
	if err != nil {
		fmt.Println("Error reading datafile containing tasks.")
		fmt.Println("Error message ->", err)
		fmt.Println("Tasks will be written to a new file.")
	}
	for _, x := range args {
		item := task.Task{Description: x}
		item.SetPriority(priority)
		taskClient.AddTask(item)
	}
	err = taskClient.SaveTasks()
	if err != nil {
		return fmt.Errorf("could not save datafile '%s'", filename)
	}
	fmt.Println("Task(s) added with success.")
	fmt.Println("Task(s) written to datafile.")
	return nil
}
