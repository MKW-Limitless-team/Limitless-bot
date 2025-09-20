package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	LICENSE_COMMAND = "license"
)

func LicenseCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(LICENSE_COMMAND, "show's user's license").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return command.ApplicationCommand
}
