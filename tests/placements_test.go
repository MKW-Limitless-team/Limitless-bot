package tests

import (
	"database/sql"
	"fmt"
	"limitless-bot/utils/ltrc"
	"testing"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	_ "github.com/ncruces/go-sqlite3/vfs/memdb"
)

func TestPlacements(t *testing.T) {

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("create", func(t *testing.T) {
		query := `CREATE TABLE IF NOT EXISTS placements (
					track TEXT,
   					discord_id TEXT PRIMARY KEY,
					flag TEXT,
					time TEXT,
					character TEXT,
					vehicle TEXT,
					drift_type TEXT,
					category TEXT
				);`

		_, err := db.Exec(query)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert", func(t *testing.T) {
		query := `INSERT INTO placements (track, discord_id, flag, time, character, vehicle, drift_type, category)
					VALUES (?,?,?,?,?,?,?,?)`

		insert, err := db.Exec(query, "Wii Mushroom Gorge", "1234567890", "🇮🇪", "2:13.340", "Mario", "Standard Bike M", "manual", "regular")

		if err != nil {
			t.Fatal(err)
		}

		println(insert.RowsAffected())
	})

	t.Run("select", func(t *testing.T) {
		query := `SELECT track, discord_id, flag, time, character, vehicle, drift_type, category 
					FROM placements`
		rows, err := db.Query(query)
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var placement ltrc.Placement

			rows.Scan(&placement.Track, &placement.DiscordID, &placement.Flag,
				&placement.Time, &placement.Character, &placement.Vehicle,
				&placement.DriftType, &placement.Category)

			fmt.Println(placement)
		}
	})

	t.Run("drop", func(t *testing.T) {
		query := `DROP TABLE placements`

		_, err := db.Exec(query)

		if err != nil {
			t.Fatal(err)
		}
	})
}
