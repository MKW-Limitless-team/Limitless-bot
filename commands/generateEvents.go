package commands

import (
	"limitless-bot/command"

	"github.com/bwmarrin/discordgo"
)

var (
	GENERATE_EVENTS_COMMAND = "generate-events"
)

func GenerateEventsCommand() *discordgo.ApplicationCommand {
	command := command.NewChatApplicationCommand(GENERATE_EVENTS_COMMAND, "Generates events for 3 days")

	return command.ApplicationCommand
}
