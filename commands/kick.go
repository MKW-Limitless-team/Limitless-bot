package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	KICK_COMMAND = "kick"
)

func KickCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(KICK_COMMAND, "Kicks a player from Limitlink").
		SetDefaultMemberPermissions(discordgo.PermissionBanMembers)

	return command.ApplicationCommand
}
