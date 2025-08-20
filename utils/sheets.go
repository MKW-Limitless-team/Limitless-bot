package utils

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func GetPlayerData() []*PlayerData {
	ctx := context.Background()

	saFile := "credentials.json"

	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(saFile))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1d682WcmXa1qKOKCTJFj89hsbOALjTMQqQIWqbgmfBY8"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, "Playerdata!A2:E12").ValueRenderOption("UNFORMATTED_VALUE").ValueRenderOption("FORMULA").Do()

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	players := make([]*PlayerData, 0)
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			if row[0] != "" {
				playerdata := &PlayerData{}

				if str, ok := row[0].(string); ok {
					playerdata.Name = str
				}

				if mmr, ok := row[3].(float64); ok {
					playerdata.Mmr = mmr
				}

				if len(row) == 5 { // hardcoded for the mii lines not showing up if empty
					if mii, ok := row[4].(string); ok {
						playerdata.Mii = mii
					}
				}

				players = append(players, playerdata)

			} else {
				break
			}
		}
	}

	return players
}
