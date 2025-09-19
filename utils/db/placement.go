package db

import (
	"limitless-bot/globals"
	"limitless-bot/utils/crc"

	r "github.com/nwoik/generate-mii/rkg"
)

func SubmitTime(bytes []byte, discordID string, category string) error {
	rkg := r.ParseRKG(bytes)
	crc := crc.CRC(bytes)
	readable := r.ConvertRkg(rkg)
	header := readable.Header

	query := `INSERT INTO placements (track, discord_id, flag, minutes, seconds, milliseconds,
					character, vehicle, drift_type, category, crc, approved)
					VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`

	globals.GetConnection().Exec(query, header.Track, "123451247890", "🇮🇪",
		header.FinishTime.Minutes, header.FinishTime.Seconds, header.FinishTime.Milliseconds,
		header.Character, header.Vehicle,
		header.DriftType, "regular", crc, false)

	return nil
}
