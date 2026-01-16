package commands

import (
	"limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	TRACKLIST_COMMAND = "tracklist"
)

func TracklistCommand() *discordgo.ApplicationCommand {
	command := command.NewChatApplicationCommand(TRACKLIST_COMMAND, "Generate a tracklist").
		SetDefaultMemberPermissions(discordgo.PermissionManageMessages)

	amountOption := &discordgo.ApplicationCommandOption{
		Name:        "amount",
		Description: "Number of tracks to select (default: 32)",
		Type:        discordgo.ApplicationCommandOptionInteger,
		Required:    false,
	}

	command.AddOption(amountOption)

	return command.ApplicationCommand
}
