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

	command.AddOption(c.NewCommandOption("user", "the user you want to view", discordgo.ApplicationCommandOptionUser, false).ApplicationCommandOption)
	command.AddOption(c.NewCommandOption("pid", "the profile ID you want to view", discordgo.ApplicationCommandOptionString, false).ApplicationCommandOption)

	sort := c.NewCommandOption("sort", "sort by number or alphabetically", discordgo.ApplicationCommandOptionString, false)
	sort.AddChoice(c.NewOptionChoice("number", "number").ApplicationCommandOptionChoice)
	sort.AddChoice(c.NewOptionChoice("alphabetical", "alphabetical").ApplicationCommandOptionChoice)
	command.AddOption(sort.ApplicationCommandOption)

	return command.ApplicationCommand
}
