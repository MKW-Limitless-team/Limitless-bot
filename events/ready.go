package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Activities: []*discordgo.Activity{
			{
				Name:     "🎮 Mario Kart Wii: Limitless",
				Type:     discordgo.ActivityTypeGame,
				URL:      "https://wiki.tockdom.com/wiki/Mario_Kart_Wii:_Limitless",
				Instance: true,
			},
		},
		Status: "online",
		AFK:    false,
	})
	fmt.Println("Ready to serve... ")
}
