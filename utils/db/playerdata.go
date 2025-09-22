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
		return errors.New("failed to edit mii, ping admin/dev")
	}

	return nil
}

func GetFlags() ([]*ltrc.Flag, error) {
	flags := []*ltrc.Flag{}

	query := `SELECT * FROM flags`

	rows, err := globals.GetConnection().Query(query)

	if err != nil {
		return nil, errors.New("failed to get flags")
	}

	for rows.Next() {
		flag := &ltrc.Flag{}

		rows.Scan(&flag.Emoji, &flag.Name)

		flags = append(flags, flag)
	}

	return flags, nil
}

func GetPlayer(userID string) (*ltrc.PlayerData, error) {
	query := `SELECT name, friend_code, mii, mmr 
					FROM playerdata
					WHERE discord_id = ?`
	rows, err := globals.GetConnection().Query(query, userID)

	if err != nil {
		return nil, errors.New("failed to fetch license, ping admin/dev")
	}
	defer rows.Close()

	playerData := &ltrc.PlayerData{
		DiscordID: userID,
	}

	if rows.Next() {
		rows.Scan(&playerData.Name, &playerData.FriendCode, &playerData.Mii, &playerData.Mmr)
	} else {
		return nil, errors.New("license not found, please /register to create one")
	}

	return playerData, nil
}

func RegisterPlayer(name string, friend_code string, discord_id string, flag string) error {
	query := `INSERT INTO playerdata (name, friend_code, discord_id, flag)
					VALUES (?, ?, ?, ?)`

	_, err := globals.GetConnection().Exec(query, name, friend_code, discord_id, flag)

	if err.Error() == "sqlite3: constraint failed: UNIQUE constraint failed: playerdata.discord_id" {
		return errors.New("user is already registered")
	} else if err != nil {
		return err
	}

	return nil
}
