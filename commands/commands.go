package commands

import (
	"github.com/bwmarrin/discordgo"
)

var globalCommands []*discordgo.ApplicationCommand = make([]*discordgo.ApplicationCommand, 0)

func RegisterCommands(session *discordgo.Session) error {

	// Add commands here
	globalCommands = append(globalCommands, PingCommand())

	// Register commands globally
	_, err := session.ApplicationCommandBulkOverwrite(session.State.User.ID, "", globalCommands)
	if err != nil {
		return err
	}

	return nil
}

func RegisterInteractions(session *discordgo.Session) {
	session.AddHandler(PingInteractionCreate)
}
