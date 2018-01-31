package cmd

import (
	"fmt"
	"strings"

	"github.com/eareese/todo/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		Add(strings.Join(args, " "))
	},
}

// Add task to the todo list by creating it in the database.
func Add(task string) {
	_, err := db.CreateTask(task)
	if err != nil {
		fmt.Println("Failed to create task:", err.Error())
		return
	}
	fmt.Printf("Added %q to the list.\n", task)
}

func init() {
	RootCmd.AddCommand(addCmd)
}
