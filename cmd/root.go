package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "todo-cli",
	Short:   "A To-Do list application written in Go.",
	Version: "v1.4.0",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Unable to detect home directory. Please set datafile using --datafile.")
	}
	datafile := filepath.Join(homeDir, ".tasks.json")
	rootCmd.PersistentFlags().String("datafile", datafile, "Datafile containing tasks.")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Display this help message.")
	rootCmd.Flags().BoolP("version", "v", false, "Display version info.")
	rootCmd.SetVersionTemplate("{{ .Version }}\n")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceUsage = true
	_ = viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
}
