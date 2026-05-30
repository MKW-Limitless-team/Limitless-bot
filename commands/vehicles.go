package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	VEHICLES_COMMAND = "vehicles"
)

func VehiclesCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(VEHICLES_COMMAND, "shows user's vehicle usage").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	command.AddOption(c.NewCommandOption("user", "mention a user or enter a PID", discordgo.ApplicationCommandOptionString, false).ApplicationCommandOption)

	return command.ApplicationCommand
}
