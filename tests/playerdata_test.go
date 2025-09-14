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

func TestPlayerData(t *testing.T) {

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("create", func(t *testing.T) {
		query := `CREATE TABLE
					IF NOT EXISTS playerdata (
						name TEXT,
						friend_code TEXT,
						discord_id TEXT,
						mmr INTEGER,
						mii TEXT
					);`

		_, err := db.Exec(query)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert", func(t *testing.T) {
		query := `INSERT INTO playerdata (name, friend_code, discord_id, mmr, mii)
					VALUES (?, ?, ?, ?, ?)`

		insert, err := db.Exec(query, "player_1", "3142-5362-6490", "12325534562", 0, "")

		if err != nil {
			t.Fatal(err)
		}

		println(insert.RowsAffected())
	})

	t.Run("select", func(t *testing.T) {
		query := `SELECT name, friend_code, discord_id, mmr, mii 
					FROM playerdata`
		rows, err := db.Query(query)
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var playerData ltrc.PlayerData

			rows.Scan(&playerData.Name, &playerData.FriendCode, &playerData.DiscordID, &playerData.Mmr, &playerData.Mii)
			fmt.Println(playerData)
		}
	})

	t.Run("drop", func(t *testing.T) {
		query := `DROP TABLE playerdata`

		_, err := db.Exec(query)

		if err != nil {
			t.Fatal(err)
		}
	})
}
