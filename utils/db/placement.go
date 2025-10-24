package db

import (
	"errors"
	"limitless-bot/globals"
	"limitless-bot/utils/crc"
	"limitless-bot/utils/ltrc"

	r "github.com/nwoik/generate-mii/rkg"
)

func GetTimeByCRC(crc uint32) (*ltrc.Placement, *ltrc.PlayerData, error) {
	query := `SELECT id,
				track,
				discord_id,
				minutes,
				seconds,
				milliseconds,
				character,
				vehicle,
				drift_type,
				category,
				url,
				approved,
				name,
				friend_code,
				mii
			FROM placements
			INNER JOIN playerdata ON placements.discord_id = playerdata.discord_id
			WHERE crc = ?`

	rows, err := globals.GetConnection().Query(query, crc)

	if err != nil {
		return nil, nil, errors.New("failed to fetch time, ping admin/dev")
	}
	defer rows.Close()

	if rows.Next() {
		placement := &ltrc.Placement{CRC: crc}
		playerdata := &ltrc.PlayerData{}

		rows.Scan(&placement.ID, &placement.Track, &placement.DiscordID,
			&placement.Minutes, &placement.Seconds, &placement.Milliseconds,
			&placement.Character, &placement.Vehicle, &placement.DriftType,
			&placement.Category, &placement.Url, &placement.Approved,
			&playerdata.Name, &playerdata.FriendCode, playerdata.Mii)

		return placement, playerdata, nil
	}

	return nil, nil, nil
}

func SubmitTime(bytes []byte, discordID string, category string, url string) error {
	rkg := r.ParseRKG(bytes)
	crc := crc.CRC(bytes)
	readable := r.ConvertRkg(rkg)
	header := readable.Header

	query := `INSERT INTO placements (track, discord_id, minutes, seconds, milliseconds,
					character, vehicle, drift_type, category, crc, url, approved)
					VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`

	_, err := globals.GetConnection().Exec(query, header.Track, discordID,
		header.FinishTime.Minutes, header.FinishTime.Seconds, header.FinishTime.Milliseconds,
		header.Character, header.Vehicle,
		header.DriftType, category, crc, url, false)

	if err != nil && err.Error() == "sqlite3: constraint failed: UNIQUE constraint failed: placements.crc" {
		return errors.New("can't upload duplicate times")
	} else if err != nil {
		return errors.New("failed to submit")
	}

	return nil
}
