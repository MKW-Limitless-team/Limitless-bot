package commands

import (
	"limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	START_TOURNAMENT_COMMAND = "start-tournament"
)

func StartTournamentCommand() *discordgo.ApplicationCommand {
	command := command.NewChatApplicationCommand(START_TOURNAMENT_COMMAND, "Starts a tournament").
		SetDefaultMemberPermissions(discordgo.PermissionManageMessages)

	return command.ApplicationCommand
}
