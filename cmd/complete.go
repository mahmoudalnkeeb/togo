package cmd

import (
	"fmt"

	"github.com/mahmoudalnkeeb/togo/internal/todo"
	"github.com/spf13/cobra"
)

var completeID int

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a todo item as complete",
	Run: func(cmd *cobra.Command, args []string) {
		success, err := todo.MarkTodoAsComplete(db, completeID)
		if err != nil {
			logger.Error("Error marking todo as complete", "error", err)
		} else if success {
			logger.Info(fmt.Sprintf("Todo with ID %d marked as complete successfully\n", completeID))
		}
	},
}

func init() {
	completeCmd.Flags().IntVarP(&completeID, "id", "i", 0, "ID of the todo item to mark as complete")
	completeCmd.MarkFlagRequired("id")
}
