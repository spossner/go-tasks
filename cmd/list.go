/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"possner.de/tasks/internal/store"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks in the todo list",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks := store.GetStore().List()

		w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
		fmt.Fprintln(w, "ID\tName\tCreated\tDue to\tCompleted")
		for _, t := range tasks {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", t.ID, t.Name, t.GetCreatedAt(), t.GetDueTo(), t.GetCompleted())
		}
		w.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
