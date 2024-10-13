package db_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/mahmoudalnkeeb/togo/config"
	"github.com/mahmoudalnkeeb/togo/db"
)

var (
	pwd    = os.Getenv("PWD")
	Config = config.LoadConfig(fmt.Sprintf("%s/.env", pwd))
)

func TestConnectSqlite(t *testing.T) {
	tmpdir := t.TempDir()

	dbFile, err := os.CreateTemp(tmpdir, "test*.db")
	if err != nil {
		t.Fatalf("Failed to create temp database file %s", err)
	}
	defer dbFile.Close()

	db, err := db.ConnectSqlite(dbFile.Name())
	t.Run("Attempt to connect to the SQLite database", func(t *testing.T) {
		if err != nil {
			t.Fatalf("Expected database to connect, instead got error: %v", err)
		}
	})
	defer db.Close()

	t.Run("Check if the connection is alive", func(t *testing.T) {
		if err := db.Ping(); err != nil {
			t.Errorf("Expected to ping the database successfully, but got error: %v", err)
		}
	})

	t.Run("Ensure at least one driver is loaded", func(t *testing.T) {
		got := len(sql.Drivers())
		want := 1
		if got < want {
			t.Errorf("Expected at least %d drivers, instead got %d", want, got)
		}
	})
}
