package main

import (
	"context"
	"fmt"
	"limitless-bot/playerdata"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	ctx := context.Background()

	// Path to your service account key file
	saFile := "credentials.json"

	// Create the Sheets service with service account credentials
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(saFile))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1d682WcmXa1qKOKCTJFj89hsbOALjTMQqQIWqbgmfBY8"
	readRange := "Playerdata!A2:M"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).ValueRenderOption("FORMULA").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			if len(row) != 0 {
				playerdata := playerdata.PlayerData{}

				if str, ok := row[0].(string); ok {
					playerdata.Name = str
				}

				if num, ok := row[3].(int); ok {
					playerdata.Mmr = num
				}

				if str, ok := row[4].(string); ok {
					playerdata.Mii = str
				}

				fmt.Printf("%+v", playerdata)
				println()
			}

		}
	}
}
