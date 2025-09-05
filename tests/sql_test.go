package tests

import (
	"database/sql"
	"testing"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	_ "github.com/ncruces/go-sqlite3/vfs/memdb"
)

func TestSQLOperations(t *testing.T) {

	t.Run("Create db", func(t *testing.T) {
		db, err := sql.Open("sqlite3", "./test.db")
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		rows, err := db.Query(`SELECT id FROM events`)
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var (
				name string
			)

			rows.Scan(&name)
			println(name)
		}
	})
}
