package cmd

import (
	"database/sql"
	"log"
	"log/slog"
	"os"

	"github.com/mahmoudalnkeeb/togo/config"
	database "github.com/mahmoudalnkeeb/togo/db"
	"github.com/mahmoudalnkeeb/togo/internal/todo"
	"github.com/mahmoudalnkeeb/togo/internal/utils"
	"github.com/spf13/cobra"
)

var (
	db     *sql.DB
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
)

var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "Togo CLI application",
	Long:  `A simple command-line application to manage your to-do items.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		var affectedRows int64

		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current directory: %v", err)
		}

		cfg := config.LoadConfig(pwd)
		db, err = database.ConnectSqlite(cfg.DBFile)
		if err != nil {
			logger.Error("Error connecting to the database", "error", err)
			return
		}
		utils.RunMigration(cfg.TodosTable, db, logger)

		affectedRows, err = todo.AutoCompleteTodos(db)
		if err != nil {
			logger.Error("Error auto-completing todos", "error", err)
		} else if affectedRows > 0 {
			logger.Info("Auto-completed %d todos.\n", "rows affected", affectedRows)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if db != nil {
			db.Close()
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(completeCmd)
}
