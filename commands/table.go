package commands

import (
	c "limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	TABLE_COMMAND = "table"
)

func TableCommand() *discordgo.ApplicationCommand {
	command := c.NewChatApplicationCommand(TABLE_COMMAND, "creates a table")

	return command.ApplicationCommand
}
