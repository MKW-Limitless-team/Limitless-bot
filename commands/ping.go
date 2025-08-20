package commands

import (
	"limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

func PingCommand() *discordgo.ApplicationCommand {
	command := command.NewChatApplicationCommand("ping", "pings the bot").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return command.ApplicationCommand
}
