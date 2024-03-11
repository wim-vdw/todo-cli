package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A To-Do list application written in Go.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("datafile", "./tasks.json", "datafile containing tasks")
	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
}
