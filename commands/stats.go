package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	STATS_COMMAND = "stats"
)

func StatsCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(STATS_COMMAND, "Shows current users").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return command.ApplicationCommand
}
