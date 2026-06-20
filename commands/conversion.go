package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	FC_TO_PID_COMMAND = "fc-to-pid"
	PID_TO_FC_COMMAND = "pid-to-fc"
)

func FCToPIDCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(FC_TO_PID_COMMAND, "converts a friend-code to a PID").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	command.AddOption(c.NewCommandOption("friend_code", "Friend-Code", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)

	return command.ApplicationCommand
}

func PIDToFCCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(PID_TO_FC_COMMAND, "converts a PID to a friend-code").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	command.AddOption(c.NewCommandOption("pid", "PID", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)

	return command.ApplicationCommand
}
