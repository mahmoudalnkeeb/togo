package cmd

import (
	"fmt"
	"time"

	"github.com/mahmoudalnkeeb/togo/internal/todo"
	"github.com/spf13/cobra"
)

var (
	title      string
	completeBy string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item",
	Run: func(cmd *cobra.Command, args []string) {
		var parsedTime *time.Time
		if completeBy != "" {
			parsed, timeErr := time.Parse("2006-01-02 15:04:05", completeBy)
			if timeErr != nil {
				logger.Error("Error parsing time, should be in '2006-01-02 15:04:05' format", "error", timeErr)
				return
			}
			parsedTime = &parsed
		}

		if _, err := todo.AddTodo(title, parsedTime, db); err != nil {
			logger.Error("Error adding todo", "error", err)
		} else {
			logger.Info(fmt.Sprintf("Todo '%s' added successfully\n", title))
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the todo item")
	addCmd.Flags().StringVarP(&completeBy, "complete_by", "c", "", "Completion deadline (format: 'YYYY-MM-DD HH:MM:SS')")
	addCmd.MarkFlagRequired("title")
}
