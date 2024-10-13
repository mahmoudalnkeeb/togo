package cmd

import (
	"fmt"

	"github.com/mahmoudalnkeeb/togo/internal/todo"
	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo item",
	Run: func(cmd *cobra.Command, args []string) {
		success, err := todo.DeleteTodo(db, deleteID)
		if err != nil {
			logger.Error("Error deleting todo", "error", err)
		} else if success {
			logger.Info(fmt.Sprintf("Todo with ID %d deleted successfully\n", deleteID))
		}
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "ID of the todo item to delete")
	deleteCmd.MarkFlagRequired("id")
}
