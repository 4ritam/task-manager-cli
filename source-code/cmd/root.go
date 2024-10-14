/*
Copyright © 2024 Pritam Das <daspritam4231@gmail.com>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-manager",
	Short: "A very simple task manager.",
	Long: `
A CLI based task manager application built using cobra and golang.
You can use this application to manage your tasks. You can add, list, do, undo and remove tasks.
You can also mark a task as completed or incomplete.

Example:
  task-manager add Buy milk
  task-manager list
  task-manager complete 1
  task-manager incomplete 1
  task-manager delete 1
  task-manager search milk`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task-manager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
