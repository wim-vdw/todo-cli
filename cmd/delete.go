package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-cli/internal/task"
	"os"
	"strconv"
	"strings"
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
	Args:    checkArgs,
	Run:     deleteTask,
}

var forceDelete bool

func init() {
	deleteCmd.Flags().BoolVar(&forceDelete, "force", false, "Immediately delete task and bypass graceful delete.")
	rootCmd.AddCommand(deleteCmd)
}

func checkArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("specify exactly 1 argument containing task-id")
	}
	return nil
}

func deleteTask(cmd *cobra.Command, args []string) {
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
			fmt.Println("Error: task-id should be a numeric value.")
			os.Exit(1)
		}
		filename := viper.GetString("datafile")
		tasks, err := task.ReadTasks(filename)
		if err != nil {
			fmt.Println("Error reading datafile containing tasks.")
			fmt.Println("Error message ->", err)
			os.Exit(1)
		}
		tasks, err = tasks.DeleteTask(taskID)
		if err != nil {
			fmt.Println("Error deleting task from the To-Do list.")
			fmt.Println("Error message ->", err)
			os.Exit(1)
		}
		fmt.Println("Task deleted with success.")
		err = task.SaveTasks(filename, tasks)
		if err != nil {
			fmt.Println("Error writing datafile containing tasks.")
			fmt.Println("Error message ->", err)
			os.Exit(1)
		}
		fmt.Println("Task(s) written to datafile.")
	}
}
