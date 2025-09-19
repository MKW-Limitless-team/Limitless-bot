package tests

import (
	"database/sql"
	"fmt"
	"limitless-bot/utils/ltrc"
	"os"
	"testing"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	_ "github.com/ncruces/go-sqlite3/vfs/memdb"
	r "github.com/nwoik/generate-mii/rkg"
)

func TestPlacements(t *testing.T) {

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("create", func(t *testing.T) {
		query := `CREATE TABLE
					IF NOT EXISTS placements (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						track TEXT,
						discord_id TEXT,
						flag TEXT,
						minutes INTEGER,
						seconds INTEGER,
						milliseconds INTEGER,
						character TEXT,
						vehicle TEXT,
						drift_type TEXT,
						category TEXT,
        				url TEXT,
						approved BOOLEAN
					);`

		_, err := db.Exec(query)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert", func(t *testing.T) {
		query := `INSERT INTO placements (track, discord_id, flag, minutes, seconds, milliseconds,
					character, vehicle, drift_type, category, approved)
					VALUES (?,?,?,?,?,?,?,?,?,?,?)`

		insert, err := db.Exec(query, "Wii Mushroom Gorge", "123453457890", "🇮🇪", 2, 13, 340,
			"Mario", "Standard Bike M",
			"MANUAL", "regular", false)

		if err != nil {
			t.Fatal(err)
		}

		println(insert.RowsAffected())
	})

	t.Run("select", func(t *testing.T) {
		query := `SELECT track, discord_id, flag, minutes, seconds, milliseconds,
					character, vehicle, drift_type, category, approved
					FROM placements`
		rows, err := db.Query(query)
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var placement ltrc.Placement

			rows.Scan(&placement.Track, &placement.DiscordID, &placement.Flag,
				&placement.Minutes, &placement.Seconds, &placement.Milliseconds,
				&placement.Character, &placement.Vehicle,
				&placement.DriftType, &placement.Category, &placement.Approved)

			fmt.Println(placement)
		}
	})

	t.Run("insert from rkg", func(t *testing.T) {
		file, err := os.ReadFile("./2m20s397.rkg")
		if err != nil {
			println(err.Error())
		}

		rkg := r.ParseRKG(file)
		readable := r.ConvertRkg(rkg)
		header := readable.Header

		query := `INSERT INTO placements (track, discord_id, flag, minutes, seconds, milliseconds,
					character, vehicle, drift_type, category, approved)
					VALUES (?,?,?,?,?,?,?,?,?,?,?)`

		insert, err := db.Exec(query, header.Track, "123451247890", "🇮🇪",
			header.FinishTime.Minutes, header.FinishTime.Seconds, header.FinishTime.Milliseconds,
			header.Character, header.Vehicle,
			header.DriftType, "regular", false)

		if err != nil {
			t.Fatal(err)
		}

		println(insert.RowsAffected())
	})

	t.Run("drop", func(t *testing.T) {
		query := `DROP TABLE placements`

		_, err := db.Exec(query)

		if err != nil {
			t.Fatal(err)
		}
	})
}
