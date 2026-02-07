package db

import (
	"errors"
	"limitless-bot/globals"
	"limitless-bot/utils/ltrc"
)

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
