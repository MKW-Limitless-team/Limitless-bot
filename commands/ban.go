package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	BAN_COMMAND = "ban"
)

func BanCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(BAN_COMMAND, "Bans a player from Limitlink").
		SetDefaultMemberPermissions(discordgo.PermissionBanMembers)

	return command.ApplicationCommand
}
