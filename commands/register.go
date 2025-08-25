package commands

import (
	"limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	REGISTER_COMMAND = "register"
)

func RegisterCommand() *discordgo.ApplicationCommand {
	command := command.NewChatApplicationCommand(REGISTER_COMMAND, "registers a user with the bot").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return command.ApplicationCommand
}
