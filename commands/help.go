package commands

import (
	"limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	HELP_COMMAND = "help"
)

func HelpCommand() *discordgo.ApplicationCommand {
	command := command.NewChatApplicationCommand(HELP_COMMAND, "Lists the bot's commands").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return command.ApplicationCommand
}
