package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	CHARACTERS_COMMAND = "characters"
)

func CharactersCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(CHARACTERS_COMMAND, "shows user's character usage").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	command.AddOption(c.NewCommandOption("user", "mention a user, enter a PID, or enter an FC", discordgo.ApplicationCommandOptionString, false).ApplicationCommandOption)

	return command.ApplicationCommand
}
