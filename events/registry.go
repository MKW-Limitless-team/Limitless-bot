package events

import "github.com/bwmarrin/discordgo"

func RegisterEvents(session *discordgo.Session) {
	// this is where all events are bound to the session
	session.AddHandler(Ready)
	session.AddHandler(InteractionCreate)
}
