package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/todo-cli/internal/task"
)

const cleanExamples = `  # Clean the complete To-Do list with confirmation message
  todo-cli clean

  # Clean the complete To-Do list without confirmation message
  todo-cli clean --force`

var cleanCmd = &cobra.Command{
	Use:     "clean",
	Short:   "Clean the complete To-Do list",
	Aliases: []string{"init", "initialize"},
	Example: cleanExamples,
	Run:     cleanTasks,
}

var forceClean bool

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().BoolVar(&forceClean, "force", false, "Immediately clean To-Do list and bypass graceful cleanup.")
}

func cleanTasks(cmd *cobra.Command, args []string) {
	if !forceClean {
		fmt.Print("Are you sure? (Y)es/(N)o): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToUpper(input))
		if input == "Y" || input == "YES" {
			forceClean = true
		}
	}
	if forceClean {
		filename := viper.GetString("datafile")
		tasks := task.Tasks{}
		err := task.SaveTasks(filename, tasks)
		if err != nil {
			fmt.Println("Error writing datafile containing tasks.")
			fmt.Println("Error message ->", err)
			os.Exit(1)
		}
		fmt.Println("Datafile containing tasks has been cleaned.")
	}
}
