package commands

import (
	"github.com/bwmarrin/discordgo"
)

var GlobalCommands []*discordgo.ApplicationCommand = make([]*discordgo.ApplicationCommand, 0)

func RegisterCommands(session *discordgo.Session) error {

	// Add commands here
	GlobalCommands = append(GlobalCommands, HelpCommand())
	GlobalCommands = append(GlobalCommands, PingCommand())
	GlobalCommands = append(GlobalCommands, OnlineCommand())
	GlobalCommands = append(GlobalCommands, TableCommand())
	GlobalCommands = append(GlobalCommands, GenerateEventsCommand())
	GlobalCommands = append(GlobalCommands, TracklistCommand())
	GlobalCommands = append(GlobalCommands, TrackFolderCommand())
	GlobalCommands = append(GlobalCommands, RKGCommand())
	// GlobalCommands = append(GlobalCommands, LeaderBoardCommand())
	// GlobalCommands = append(GlobalCommands, RegisterCommand())
	// GlobalCommands = append(GlobalCommands, SubmitTimeCommand())
	// GlobalCommands = append(GlobalCommands, EditMiiCommand())
	// GlobalCommands = append(GlobalCommands, LicenseCommand())

	// Register commands globally
	_, err := session.ApplicationCommandBulkOverwrite(session.State.User.ID, "", GlobalCommands)
	if err != nil {
		return err
	}

	return nil
}
