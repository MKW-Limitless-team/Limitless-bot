package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	ONLINE_COMMAND = "online"
)

func OnlineCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(ONLINE_COMMAND, "Shows current users").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return command.ApplicationCommand
}
