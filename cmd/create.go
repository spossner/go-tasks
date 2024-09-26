/*
Copyright © 2024 NAME HERE <EMAIL createRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"possner.de/tasks/internal/store"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [task]",
	Short: "Add a new task to the todo list",
	Args:  cobra.ExactArgs(1),
	RunE:  createTask,
}

func createTask(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("missing task name")
	}
	t := store.GetStore().Create(args[0])
	fmt.Println("createed", t)
	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
