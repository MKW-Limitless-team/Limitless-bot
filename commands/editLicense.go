package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	EDIT_LICENSE_COMMAND = "edit-license"
)

func EditLicenseCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(EDIT_LICENSE_COMMAND, "allows user to edit their license details").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	command.AddOption(c.NewCommandOption("name", "Name", discordgo.ApplicationCommandOptionString, false).ApplicationCommandOption)
	command.AddOption(c.NewCommandOption("friend_code", "Friend-Code", discordgo.ApplicationCommandOptionString, false).ApplicationCommandOption)

	return command.ApplicationCommand
}
