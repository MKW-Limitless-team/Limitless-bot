package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	EDIT_MII_COMMAND = "edit-mii"
)

func EditMiiCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(EDIT_MII_COMMAND, "edits the user's mii avatar").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	command.AddOption(c.NewCommandOption("mii", ".miigx file", discordgo.ApplicationCommandOptionAttachment, true).ApplicationCommandOption)

	return command.ApplicationCommand
}
