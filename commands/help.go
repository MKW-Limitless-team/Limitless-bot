package commands

import (
	"limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

func HelpCommand() *discordgo.ApplicationCommand {
	command := command.NewChatApplicationCommand("help", "Lists the bot's commands").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return command.ApplicationCommand
}
