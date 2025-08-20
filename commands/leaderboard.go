package commands

import (
	"limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

func LeaderBoardCommand() *discordgo.ApplicationCommand {
	command := command.NewChatApplicationCommand("leaderboard", "shows the leaderboard").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return command.ApplicationCommand
}
