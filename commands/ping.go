package commands

import (
	"limitless-bot/command"
	"limitless-bot/command/response"

	"github.com/bwmarrin/discordgo"
)

func PingCommand() *discordgo.ApplicationCommand {
	ping := command.NewChatApplicationCommand("ping", "pings the bot")

	return ping.ApplicationCommand
}

func PingInteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	response := response.
		NewMessageResponse().
		SetInteractionResponseData(PingData())

	_ = session.InteractionRespond(interaction.Interaction, response.InteractionResponse)
}

func PingData() *discordgo.InteractionResponseData {
	data := response.NewResponseData("pong")

	return data.InteractionResponseData
}
