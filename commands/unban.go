package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	UNBAN_COMMAND = "unban"
)

func UnbanCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(UNBAN_COMMAND, "Unbans a player from Limitlink").
		SetDefaultMemberPermissions(discordgo.PermissionBanMembers)

	command.AddOption(c.NewCommandOption("friend_code", "Friend-Code", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)

	return command.ApplicationCommand
}
