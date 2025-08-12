package events

import "github.com/bwmarrin/discordgo"

func ConfigureEvents(session *discordgo.Session) {
	// this is where all events are bound to the session
	session.AddHandler(Ready)
}
