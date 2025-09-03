package tests

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSQLOperations(t *testing.T) {

	t.Run("Create db", func(t *testing.T) {
		db, err := sql.Open("sqlite3", "./test.db")
		if err != nil {
			t.Fatal(err)
		}

		sqlStmt := `
			CREATE TABLE IF NOT EXISTS users (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				name TEXT,
				age INTEGER
			);
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			t.Fatal(err)
		}

	})
}
