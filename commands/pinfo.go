package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	PINFO_COMMAND = "pinfo"
)

func PinfoCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(PINFO_COMMAND, "shows player info from Limitlink").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	command.AddOption(c.NewCommandOption("friend_code", "the friend code you want to view", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)

	return command.ApplicationCommand
}
