package commands

import (
	"limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	LEADERBOARD_COMMAND = "leaderboard"
)

func LeaderBoardCommand() *discordgo.ApplicationCommand {
	command := command.NewChatApplicationCommand(LEADERBOARD_COMMAND, "shows the leaderboard").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return command.ApplicationCommand
}
