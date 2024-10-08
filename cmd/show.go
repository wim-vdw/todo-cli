package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-client/task"
)

var (
	displayPriority bool
	sortPriority    bool
)

const showExamples = `  # Show tasks
  todo-cli show

  # Show tasks including priority
  todo-cli show --priority

  # Show tasks sorted by priority
  todo-cli show --sorted

  # Show tasks including priority sorted by priority
  todo-cli show --sorted --priority`

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Display tasks from the To-Do list",
	Aliases: []string{"display", "list"},
	Example: showExamples,
	RunE:    showTasks,
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVarP(&displayPriority, "priority", "p", false, "Display priority.")
	showCmd.Flags().BoolVarP(&sortPriority, "sort", "s", false, "Sort by priority (finished tasks are put at the end).")
}

func showTasks(cmd *cobra.Command, args []string) error {
	filename := viper.GetString("datafile")
	taskClient := &task.Client{
		Filename: filename,
	}
	err := taskClient.ReadTasks()
	if err != nil {
		return fmt.Errorf("could not read datafile '%s'", filename)
	}
	taskClient.DisplayTasks(sortPriority, displayPriority)
	return nil
}
