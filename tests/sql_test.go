package tests

import (
	"database/sql"
	"testing"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	_ "github.com/ncruces/go-sqlite3/vfs/memdb"
)

func TestSQLOperations(t *testing.T) {

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("create", func(t *testing.T) {
		query := `CREATE TABLE IF NOT EXISTS events (
   					id INTEGER PRIMARY KEY AUTOINCREMENT,
    				name TEXT
				);`

		_, err := db.Exec(query)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert", func(t *testing.T) {
		query := `INSERT INTO events (name)
					VALUES ('event 1')`

		insert, err := db.Exec(query)

		if err != nil {
			t.Fatal(err)
		}

		println(insert.RowsAffected())
	})

	t.Run("select", func(t *testing.T) {
		query := `SELECT id FROM events`
		rows, err := db.Query(query)
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var id int

			rows.Scan(&id)
			println(id)
		}
	})

	t.Run("update", func(t *testing.T) {
		query := `UPDATE events 
					SET name = ?
					WHERE id = ?`

		update, err := db.Exec(query, "event 2", 1)

		if err != nil {
			t.Fatal(err)
		}

		println(update.RowsAffected())
	})

	t.Run("replace", func(t *testing.T) {
		query := `REPLACE INTO events (id, name)
					VALUES (?, ?)`

		replace, err := db.Exec(query, 1, "event 1")

		if err != nil {
			t.Fatal(err)
		}

		println(replace.RowsAffected())
	})

	t.Run("delete", func(t *testing.T) {
		query := `DELETE FROM events
					WHERE id = ?`

		delete, err := db.Exec(query, 1)

		if err != nil {
			t.Fatal(err)
		}

		println(delete.RowsAffected())
	})

	t.Run("drop", func(t *testing.T) {
		query := `DROP TABLE events`

		_, err := db.Exec(query)

		if err != nil {
			t.Fatal(err)
		}
	})
}
