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

	command.AddOption(c.NewCommandOption("user", "mention a user, enter a PID, or enter an FC", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)

	return command.ApplicationCommand
}
