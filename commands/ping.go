package commands

import (
	"limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	PING_COMMAND = "ping"
)

func PingCommand() *discordgo.ApplicationCommand {
	command := command.NewChatApplicationCommand(PING_COMMAND, "pings the bot").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return command.ApplicationCommand
}
