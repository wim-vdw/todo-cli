package cmd

import (
	"fmt"
	"github.com/wim-vdw/todo-cli/internal/task"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	filename := viper.GetString("datafile")
	tasks, err := task.ReadTasks(filename)
	if err != nil {
		fmt.Println("Error reading datafile containing tasks.")
		fmt.Println("Error message ->", err)
		os.Exit(1)
	}
	if len(tasks) == 0 {
		fmt.Println("Nothing on your To-Do list for the moment.")
	}
	for _, t := range tasks {
		fmt.Println(t.Description)
	}
}
