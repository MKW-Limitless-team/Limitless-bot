package commands

import (
	"github.com/bwmarrin/discordgo"
)

var GlobalCommands []*discordgo.ApplicationCommand = make([]*discordgo.ApplicationCommand, 0)

func RegisterCommands(session *discordgo.Session) error {

	// Add commands here
	GlobalCommands = append(GlobalCommands, HelpCommand())
	GlobalCommands = append(GlobalCommands, PingCommand())
	GlobalCommands = append(GlobalCommands, LeaderBoardCommand())
	GlobalCommands = append(GlobalCommands, RegisterCommand())

	// Register commands globally
	_, err := session.ApplicationCommandBulkOverwrite(session.State.User.ID, "", GlobalCommands)
	if err != nil {
		return err
	}

	return nil
}
