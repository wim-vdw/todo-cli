package cmd

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

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
	tasks, err := task.ReadTasks(filename)
	if err != nil {
		return fmt.Errorf("could not read datafile '%s'", filename)
	}
	if len(tasks) == 0 {
		fmt.Println("Nothing on your To-Do list for the moment.")
	} else {
		if sortPriority {
			sort.Sort(task.ByPriority(tasks))
		}
		w := tabwriter.NewWriter(os.Stdout, 4, 0, 1, ' ', 0)
		printTitles(w, displayPriority)
		printTasks(w, &tasks, displayPriority)
		w.Flush()
	}
	return nil
}

func printTitles(w *tabwriter.Writer, displayPriority bool) {
	if displayPriority {
		fmt.Fprintln(w, "ID\tTask\tStatus\tPriority")
		fmt.Fprintln(w, "--\t----\t------\t--------")
	} else {
		fmt.Fprintln(w, "ID\tTask\tStatus")
		fmt.Fprintln(w, "--\t----\t------")
	}
}

func printTasks(w *tabwriter.Writer, tasks *task.Tasks, displayPriority bool) {
	for _, t := range *tasks {
		if displayPriority {
			fmt.Fprintln(w, t.PrettyPosition()+"\t"+t.Description+"\t"+t.PrettyStatus()+"\t"+t.PrettyPriority())
		} else {
			fmt.Fprintln(w, t.PrettyPosition()+"\t"+t.Description+"\t"+t.PrettyStatus())
		}
	}
}
