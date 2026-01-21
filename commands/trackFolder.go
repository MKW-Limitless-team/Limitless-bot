package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	TRACKFOLDER_COMMAND = "track-folder"
	TRACK_OPTION_NAME   = "track-name"
)

func TrackFolderCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(TRACKFOLDER_COMMAND, "Gets the folder name for a track")

	track := c.NewCommandOption(TRACK_OPTION_NAME, "name of track", discordgo.ApplicationCommandOptionString, true).
		SetAutoComplete(true)
	command.AddOption(track.ApplicationCommandOption)

	return command.ApplicationCommand
}
