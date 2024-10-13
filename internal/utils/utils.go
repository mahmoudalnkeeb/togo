package utils

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/mahmoudalnkeeb/togo/dtos"
)

func ParseTodoRows(rows *sql.Rows) ([]dtos.Todo, error) {
	var todos []dtos.Todo
	for rows.Next() {
		var t dtos.Todo
		var completeInt int
		var completeBy sql.NullTime // Use sql.NullTime to handle nullable time

		err := rows.Scan(&t.ID, &t.Title, &completeInt, &completeBy)
		if err != nil {
			log.Fatal(err)
			continue // Add continue to skip this iteration on error
		}

		t.Complete = completeInt == 1

		if completeBy.Valid {
			t.CompleteBy = completeBy.Time // Assign only if valid
		} else {
			t.CompleteBy = time.Time{} // Assign zero value for time if NULL
		}

		todos = append(todos, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func ParseTodoRow(row *sql.Row) (*dtos.Todo, error) {
	var todo dtos.Todo
	var completeInt int
	var completeByStr string
	err := row.Scan(&todo.ID, &todo.Title, &completeInt, &completeByStr)
	if err != nil {
		return nil, err
	}

	todo.Complete = completeInt == 1

	if completeByStr != "" {
		todo.CompleteBy, err = StringToTime(completeByStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing complete_by date: %v", err)
		}
	}

	return &todo, nil
}

func StringToTime(timestr string) (time.Time, error) {
	return time.Parse(time.RFC3339, timestr)
}

func RunMigration(query string, db *sql.DB, logger *slog.Logger) error {
	_, err := db.Exec(query)
	if err != nil {
		logger.Error("failed to execute migration", "error", err)
		return err
	}
	return nil
}

func FormatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
