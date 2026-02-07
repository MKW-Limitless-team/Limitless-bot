package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/MKW-Limitless-team/limitless-types/ltrc"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func GetPlayerData() []*ltrc.PlayerData {
	ctx := context.Background()

	saFile := "credentials.json"

	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(saFile))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1d682WcmXa1qKOKCTJFj89hsbOALjTMQqQIWqbgmfBY8"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, "Playerdata!A2:G").ValueRenderOption("UNFORMATTED_VALUE").ValueRenderOption("FORMULA").Do()

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	players := make([]*ltrc.PlayerData, 0)
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			if row[0] != "" {
				playerdata := &ltrc.PlayerData{}

				// if str, ok := row[0].(string); ok {
				// 	playerdata.Name = str
				// }

				if mmr, ok := row[3].(float64); ok {
					MMR := int64(mmr)
					playerdata.Mmr = &MMR
				}

				if len(row) > 6 {
					// if mii, ok := row[4].(string); ok {
					// 	playerdata.Mii = mii
					// }

					if discordID, ok := row[5].(string); ok {
						playerdata.DiscordID = discordID
					}

					if profileID, ok := row[6].(float64); ok {
						playerdata.ProfileID = uint64(profileID)
					}
					players = append(players, playerdata)

				}

			} else {
				break
			}
		}
	}

	b, err := json.Marshal(players)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	return players
}
