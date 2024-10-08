package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-client/task"
)

const doneExamples = `  # Mark task with ID 1 as done
  todo-cli done 1`

var doneCmd = &cobra.Command{
	Use:     "done task-id",
	Short:   "Mark a task as done in the To-Do list",
	Aliases: []string{"finish", "ok"},
	Example: doneExamples,
	Args:    checkArgsDone,
	RunE:    doneTask,
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

func doneTask(cmd *cobra.Command, args []string) error {
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("task-id should be a numeric value")
	}
	filename := viper.GetString("datafile")
	taskClient := &task.Client{
		Filename: filename,
	}
	err = taskClient.ReadTasks()
	if err != nil {
		return fmt.Errorf("could not read datafile '%s'", filename)
	}
	err = taskClient.FinishTask(taskID)
	if err != nil {
		return err
	}
	err = taskClient.SaveTasks()
	if err != nil {
		return fmt.Errorf("could not save datafile '%s'", filename)
	}
	fmt.Println("Task finished with success.")
	fmt.Println("Task(s) written to datafile.")
	return nil
}
