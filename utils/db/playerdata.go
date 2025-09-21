package db

import (
	"errors"
	"fmt"
	"limitless-bot/globals"
	"limitless-bot/utils/ltrc"
)

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

func GetPlayer(userID string) (*ltrc.PlayerData, error) {
	query := `SELECT name, friend_code, mii, discord_id, mmr 
					FROM playerdata
					WHERE discord_id = ?`
	rows, err := globals.GetConnection().Query(query, userID)

	if err != nil {
		return nil, errors.New("failed to fetch license, ping admin")
	}
	defer rows.Close()

	playerData := &ltrc.PlayerData{}

	if rows.Next() {
		rows.Scan(&playerData.Name, &playerData.FriendCode, &playerData.Mii, &playerData.DiscordID, &playerData.Mmr)
	} else {
		return nil, errors.New("license not found, please /register to create one")
	}

	return playerData, nil
}

func RegisterPlayer(name string, friend_code string, discord_id string) error {
	query := `INSERT INTO playerdata (name, friend_code, discord_id)
					VALUES (?, ?, ?)`

	_, err := globals.GetConnection().Exec(query, name, friend_code, discord_id)

	if err != nil {
		return err
	}

	return nil
}
