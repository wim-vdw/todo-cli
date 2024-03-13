package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/user"
	"path/filepath"
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
	u, err := user.Current()
	if err != nil {
		fmt.Println("Unable to detect home directory. Please set datafile using --datafile.")
	}
	datafile := filepath.Join(u.HomeDir, ".tasks.json")
	rootCmd.PersistentFlags().String("datafile", datafile, "Datafile containing tasks.")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Display this help message.")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
}
