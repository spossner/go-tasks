/*
Copyright © 2024 NAME HERE <EMAIL createRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/tj/go-naturaldate"
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
	due, err := cmd.Flags().GetString("due")
	if err != nil {
		return errors.New("error in due date flag")
	}
	dueDate := time.Time{}
	if due != "" {
		now := time.Now()
		dueDate, err = naturaldate.Parse(due, now)
		if err != nil || now.Equal(dueDate) {
			return errors.New("error parsing due date " + due)
		}
		if dueDate.Before(now) {
			return errors.New("can not set due date in the past: " + due)
		}
	}
	t := store.GetStore().Create(args[0], dueDate)
	fmt.Println("createed", t)
	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("due", "", "set an optional due date im human readable format (e.g. 'tomorrow' or 'in 2 days')")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
