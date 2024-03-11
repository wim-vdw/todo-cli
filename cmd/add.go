package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-cli/internal/task"
	"os"
)

var addCmd = &cobra.Command{
	Use:     "add tasks...",
	Short:   "Add tasks to the To-Do list",
	Aliases: []string{"create", "new"},
	Run:     addTask,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: specify at least 1 task to add.")
		cmd.Help()
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
		tasks = append(tasks, item)
	}
	err = task.SaveTasks(filename, tasks)
	if err != nil {
		fmt.Println("Error writing datafile containing tasks.")
		fmt.Println("Error message ->", err)
	}
	fmt.Println("Task(s) written to datafile.")
}
