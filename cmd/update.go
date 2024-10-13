package cmd

import (
	"fmt"

	"github.com/mahmoudalnkeeb/togo/internal/todo"
	"github.com/spf13/cobra"
)

var updateID int

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a todo item title",
	Run: func(cmd *cobra.Command, args []string) {
		success, err := todo.UpdateTodoTitle(db, updateID, title)
		if err != nil {
			logger.Info("Error updating todo", "error", err)
		} else if success {
			logger.Info(fmt.Sprintf("Todo with ID %d updated to title '%s' successfully\n", updateID, title))
		}
	},
}

func init() {
	updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "ID of the todo item to update")
	updateCmd.Flags().StringVarP(&title, "title", "t", "", "New title for the todo item")
	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("title")
}
