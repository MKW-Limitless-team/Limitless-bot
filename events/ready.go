package events

import (
	"encoding/json"
	"fmt"
	"limitless-bot/responses"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	go getMKWIIStats(session)
	fmt.Println("Ready to serve... ")
}

func getMKWIIStats(session *discordgo.Session) {
	ticker := time.NewTicker(2 * time.Second)

	for {
		_, ok := <-ticker.C
		if ok {
			resp, err := http.Get("http://localhost/api/stats?game=mariokartwii")

			if err != nil {
				UpdateStatus(session, "🎮 Limitlink is down 😵")
			} else {
				stats := map[string]*responses.Stats{}
				err = json.NewDecoder(resp.Body).Decode(&stats)

				global := stats["global"]
				if global.OnlinePlayerCount == 0 {
					UpdateStatus(session, fmt.Sprintf("🎮 %d Players Online 😴", global.OnlinePlayerCount))
				} else if global.OnlinePlayerCount == 1 {
					UpdateStatus(session, fmt.Sprintf("🎮 %d Player Online 😔", global.OnlinePlayerCount))
				} else if global.OnlinePlayerCount < 5 {
					UpdateStatus(session, fmt.Sprintf("🎮 %d Players Online 😃", global.OnlinePlayerCount))
				} else {
					UpdateStatus(session, fmt.Sprintf("🎮 %d Players Online 😎", global.OnlinePlayerCount))
				}
			}
		}
	}
}

func UpdateStatus(session *discordgo.Session, status string) {
	session.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Activities: []*discordgo.Activity{
			{
				Name:     status,
				Type:     discordgo.ActivityTypeGame,
				URL:      "https://wiki.tockdom.com/wiki/Mario_Kart_Wii:_Limitless",
				Instance: true,
			},
		},
		Status: "online",
		AFK:    false,
	})
}
