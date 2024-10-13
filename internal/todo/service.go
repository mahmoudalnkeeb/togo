package todo

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/mahmoudalnkeeb/togo/dtos"
	utils "github.com/mahmoudalnkeeb/togo/internal/utils"
)

func AddTodo(title string, completeBy *time.Time, db *sql.DB) (dtos.Todo, error) {
	query := `
		INSERT INTO togo (title, complete_by, complete)
		VALUES (?, ?, ?)
	`

	var completeByValue interface{}
	if completeBy != nil {
		completeByValue = completeBy.Format(time.RFC3339)
	} else {
		completeByValue = nil
	}

	result, err := db.Exec(query, title, completeByValue, false)
	if err != nil {
		log.Fatalf("Failed to add todo: %v", err)
		return dtos.Todo{}, err
	}

	todoID, err := result.LastInsertId()
	if err != nil {
		return dtos.Todo{}, fmt.Errorf("could not retrieve last insert ID: %v", err)
	}

	var completeByParsed time.Time
	if completeBy != nil {
		completeByParsed = *completeBy
	} else {
		completeByParsed = time.Time{}
	}

	return dtos.Todo{
		ID:         int(todoID),
		Title:      title,
		Complete:   false,
		CompleteBy: completeByParsed,
	}, nil
}

// just for testing
func GetTodo(id int, db *sql.DB) (*dtos.Todo, error) {
	query := "SELECT ID , title , complete , complete_by FROM togo WHERE ID=?"
	row := db.QueryRow(query, id)

	todo, err := utils.ParseTodoRow(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no todo found with ID %d", id)
		}
		log.Fatal(err)
		return nil, err
	}
	return todo, nil
}

func ListTodos(db *sql.DB) ([]dtos.Todo, error) {
	query := "SELECT ID , title , complete , complete_by FROM togo"
	rows, qErr := db.Query(query)
	if qErr != nil {
		log.Fatal(qErr)
		return nil, qErr
	}
	todos, err := utils.ParseTodoRows(rows)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return todos, nil
}

func DeleteTodo(db *sql.DB, id int) (bool, error) {
	query := "DELETE FROM togo WHERE ID = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete todo with ID %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("could not check rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return false, fmt.Errorf("no todo found with ID %d", id)
	}

	return true, nil
}

func UpdateTodoTitle(db *sql.DB, id int, title string) (bool, error) {
	query := "UPDATE togo SET title = ? WHERE ID = ?"

	result, err := db.Exec(query, title, id)
	if err != nil {
		return false, fmt.Errorf("failed to update title for todo with ID %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("could not check rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return false, fmt.Errorf("no todo found with ID %d", id)
	}

	return true, nil
}

func MarkTodoAsComplete(db *sql.DB, id int) (bool, error) {
	query := "UPDATE togo SET complete = 1 WHERE ID = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		return false, fmt.Errorf("failed to mark todo as complete with ID %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("could not check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return false, fmt.Errorf("no todo found with ID %d", id)
	}

	return true, nil
}

func AutoCompleteTodos(db *sql.DB) (int64, error) {
	query := `
		UPDATE togo
		SET complete = 1
		WHERE complete_by IS NOT NULL
		AND complete_by < CURRENT_TIMESTAMP
		AND complete = 0;
	`
	result, err := db.Exec(query)
	if err != nil {
		return 0, fmt.Errorf("failed to auto-complete todos: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not check rows affected: %v", err)
	}

	return rowsAffected, nil
}
