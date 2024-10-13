package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectSqlite(dbpath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbpath)
	db.Exec("")
	if err != nil {
		return nil, err
	}
	return db, nil
}
