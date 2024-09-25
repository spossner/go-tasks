/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"possner.de/tasks/internal/store"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task to the todo list",
	Args:  cobra.ExactArgs(1),
	RunE:  addTask,
}

func addTask(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("missing task name")
	}
	t := store.GetStore().Add(args[0])
	fmt.Println("added", t)
	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
