package db

import (
	"fmt"
	"limitless-bot/globals"
)

func RegisterPlayer(name string, friend_code string, discord_id string) error {
	query := `INSERT INTO playerdata (name, friend_code, discord_id)
					VALUES (?, ?, ?)`

	_, err := globals.GetConnection().Exec(query, name, friend_code, discord_id)

	if err != nil {
		return err
	}

	return nil
}

func EditMii(mii string, userID string) error {
	query := fmt.Sprintf(`UPDATE playerdata
				SET mii = '%s'
				WHERE discord_id = '%s'`, mii, userID)

	_, err := globals.GetConnection().Exec(query)

	if err != nil {
		return err
	}

	return nil
}
