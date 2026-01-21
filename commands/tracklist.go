package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	TRACKLIST_COMMAND = "tracklist"
)

func TracklistCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(TRACKLIST_COMMAND, "Generate a tracklist").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	amountOption := c.NewCommandOption("amount", "Number of tracks to select (default: 32)", discordgo.ApplicationCommandOptionInteger, false)

	command.AddOption(amountOption.ApplicationCommandOption)

	return command.ApplicationCommand
}
