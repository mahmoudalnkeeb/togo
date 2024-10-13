package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/mahmoudalnkeeb/togo/internal/todo"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := todo.ListTodos(db)
		if err != nil {
			fmt.Println("Error listing todos:", err)
			return
		}

		// Display the data in table format
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Complete", "Complete By"})

		for _, t := range todos {
			var completeByStr string
			if t.CompleteBy.IsZero() {
				completeByStr = ""
			} else {
				completeByStr = t.CompleteBy.Format(time.DateTime)
			}

			table.Append([]string{
				fmt.Sprintf("%d", t.ID),
				t.Title,
				fmt.Sprintf("%v", t.Complete),
				completeByStr,
			})
		}

		table.Render()
	},
}
