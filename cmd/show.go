package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-cli/internal/task"
	"os"
)

var displayPriority bool

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Display tasks from the To-Do list",
	Aliases: []string{"display", "list"},
	Run:     showTasks,
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVarP(&displayPriority, "priority", "p", false, "display priority")
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
		if displayPriority {
			fmt.Println(t.PrettyPosition(), t.Description, t.PrettyPriority())
		} else {
			fmt.Println(t.PrettyPosition(), t.Description)
		}
	}
}
