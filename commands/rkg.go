package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	RKG_COMMAND = "rkg"
)

func RKGCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(RKG_COMMAND, "shows detail of .rkg file").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	command.AddOption(c.NewCommandOption("rkg", ".rkg file", discordgo.ApplicationCommandOptionAttachment, true).ApplicationCommandOption)

	track := c.NewCommandOption(TRACK_OPTION_NAME, "name of track", discordgo.ApplicationCommandOptionString, true).
		SetAutoComplete(true)
	command.AddOption(track.ApplicationCommandOption)

	return command.ApplicationCommand
}
