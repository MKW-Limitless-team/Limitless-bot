package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	SUBMIT_TIME_COMMAND = "submit-time"
)

func SubmitTimeCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(SUBMIT_TIME_COMMAND, "submit a time trial for leaderboard approval").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	command.AddOption(c.NewCommandOption("file", "ghost file *.rkg", discordgo.ApplicationCommandOptionAttachment, true).ApplicationCommandOption)

	return command.ApplicationCommand
}
