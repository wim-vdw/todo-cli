package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-client/task"
)

const deleteExamples = `  # Delete task with ID 1 with confirmation message
  todo-cli delete 1

  # Delete task with ID 1 without confirmation message
  todo-cli delete 1 --force`

var deleteCmd = &cobra.Command{
	Use:     "delete task-id",
	Short:   "Delete task from the To-Do list",
	Aliases: []string{"del", "rm", "remove"},
	Example: deleteExamples,
	Args:    checkArgsDelete,
	RunE:    deleteTask,
}

var forceDelete bool

func init() {
	deleteCmd.Flags().BoolVar(&forceDelete, "force", false, "Immediately delete task and bypass graceful delete.")
	rootCmd.AddCommand(deleteCmd)
}

func checkArgsDelete(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("specify exactly 1 argument containing task-id")
	}
	return nil
}

func deleteTask(cmd *cobra.Command, args []string) error {
	if !forceDelete {
		fmt.Print("Are you sure? (Y)es/(N)o): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToUpper(input))
		if input == "Y" || input == "YES" {
			forceDelete = true
		}
	}
	if forceDelete {
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
		err = taskClient.DeleteTask(taskID)
		if err != nil {
			return err
		}
		err = taskClient.SaveTasks()
		if err != nil {
			return fmt.Errorf("could not save datafile '%s'", filename)
		}
		fmt.Println("Task deleted with success.")
		fmt.Println("Task(s) written to datafile.")
	}
	return nil
}
