package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	REGISTER_COMMAND = "register"
)

func RegisterCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(REGISTER_COMMAND, "registers a user with the bot").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	command.AddOption(c.NewCommandOption("mii", ".miigx file", discordgo.ApplicationCommandOptionAttachment, true).ApplicationCommandOption)
	command.AddOption(c.NewCommandOption("ign", "In-game name", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)
	command.AddOption(c.NewCommandOption("fc", "Friend-Code", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)

	return command.ApplicationCommand
}
