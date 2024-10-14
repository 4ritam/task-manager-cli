/*
Copyright © 2024 Pritam Das <daspritam4231@gmail.com>
*/

package cmd

import (
	"fmt"
	"strconv"

	"task-manager/core"

	"github.com/spf13/cobra"
)

// addCmd represents the manager command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the task manager.",
	Long: `Add a new task to the task manager. For example:
	task-manager add "Buy milk"
	task-manager add "Buy eggs"`,

	Run: func(cmd *cobra.Command, args []string) {
		taskManager := core.NewTaskManager()
		description := ""
		for _, arg := range args {
			description += arg + " "
		}
		taskManager.AddTask(description)
		fmt.Println("Task added successfully. Task ID: ", taskManager.LastID)

	},
}

// listCmd represents the manager command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the task manager.",
	Long: `List all tasks in the task manager. For example:
	task-manager list`,
	Run: func(cmd *cobra.Command, args []string) {
		taskManager := core.NewTaskManager()
		tasks := taskManager.GetTasks()
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("Tasks:")
		for _, t := range tasks {
			var status string
			if t.Status == "pending" {
				status = "   "
			} else {
				status = " ❌"
			}

			fmt.Printf("[%s]  %d: %s\n", status, t.ID, t.Description)
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the task manager.",
	Long: `Delete a task from the task manager. For example:
	task-manager delete 1`,
	Run: func(cmd *cobra.Command, args []string) {
		taskManager := core.NewTaskManager()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}
		taskManager.DeleteTask(id)
		fmt.Println("Task deleted successfully.")
	},
}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed.",
	Long: `Mark a task as completed. For example:
	task-manager complete 1`,
	Run: func(cmd *cobra.Command, args []string) {
		taskManager := core.NewTaskManager()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}
		taskManager.CompleteTask(id)
		fmt.Println("Task marked as completed.")
	},
}

var incompleteCmd = &cobra.Command{
	Use:   "incomplete",
	Short: "Mark a task as incomplete.",
	Long: `Mark a task as incomplete. For example:
	task-manager incomplete 1`,
	Run: func(cmd *cobra.Command, args []string) {
		taskManager := core.NewTaskManager()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}
		taskManager.IncompleteTask(id)
		fmt.Println("Task marked as incomplete.")
	},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for tasks.",
	Long: `Search for tasks. For example:
	task-manager search milk`,
	Run: func(cmd *cobra.Command, args []string) {
		taskManager := core.NewTaskManager()
		tasks := taskManager.SearchTasks(args[0])
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("Tasks:")
		for _, t := range tasks {
			var status string
			if t.Status == "pending" {
				status = "   "
			} else {
				status = " ❌"
			}

			fmt.Printf("[%s]  %d: %s\n", status, t.ID, t.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(incompleteCmd)
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// managerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// managerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
