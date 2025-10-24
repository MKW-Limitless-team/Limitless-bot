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

	command.AddOption(c.NewCommandOption("rkg", ".rkg file", discordgo.ApplicationCommandOptionAttachment, true).ApplicationCommandOption)

	category := c.NewCommandOption("category", "category", discordgo.ApplicationCommandOptionString, true)

	regular := c.NewOptionChoice("Regular", "regular")
	shortcut := c.NewOptionChoice("Shortcut", "shortcut")
	glitch := c.NewOptionChoice("Glitch", "glitch")

	category.AddChoice(regular.ApplicationCommandOptionChoice)
	category.AddChoice(shortcut.ApplicationCommandOptionChoice)
	category.AddChoice(glitch.ApplicationCommandOptionChoice)
	command.AddOption(category.ApplicationCommandOption)

	return command.ApplicationCommand
}
