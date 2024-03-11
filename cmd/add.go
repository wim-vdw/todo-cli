package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add tasks to the To-Do list",
	Aliases: []string{"create", "new"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called", args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
