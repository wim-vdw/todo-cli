package cmd

import (
	"fmt"
	"github.com/wim-vdw/todo-cli/internal/task"
	"os"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Display tasks from the To-Do list",
	Aliases: []string{"display", "list"},
	Run:     showTasks,
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func showTasks(cmd *cobra.Command, args []string) {
	tasks, err := task.ReadTasks("./tasks.json")
	if err != nil {
		fmt.Println("Error reading datafile containing tasks.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	for _, t := range tasks {
		fmt.Println(t.Description)
	}
}
