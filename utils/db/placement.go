package db

import (
	"errors"
	"limitless-bot/globals"
	"limitless-bot/utils/crc"

	r "github.com/nwoik/generate-mii/rkg"
)

func SubmitTime(bytes []byte, discordID string, category string, url string) error {
	rkg := r.ParseRKG(bytes)
	crc := crc.CRC(bytes)
	readable := r.ConvertRkg(rkg)
	header := readable.Header

	query := `INSERT INTO placements (track, discord_id, flag, minutes, seconds, milliseconds,
					character, vehicle, drift_type, category, crc, url, approved)
					VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`

	_, err := globals.GetConnection().Exec(query, header.Track, discordID, header.Country,
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
