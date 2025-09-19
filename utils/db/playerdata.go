package db

import (
	"database/sql"
	"log"
)

func RegisterPlayer(name string, friend_code string, discord_id string, mii string) error {
	query := `INSERT INTO playerdata (name, friend_code, discord_id, mii)
					VALUES (?, ?, ?, ?)`

	db, err := sql.Open("sqlite3", "./ltrc.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec(query, name, friend_code, discord_id, mii)

	if err != nil {
		return err
	}

	return nil
}
