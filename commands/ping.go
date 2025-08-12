package commands

import (
	"limitless-bot/command"
	"limitless-bot/command/response"

	"github.com/bwmarrin/discordgo"
)

func PingCommand() *discordgo.ApplicationCommand {
	ping := command.NewChatApplicationCommand("ping", "pings the bot").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return ping.ApplicationCommand
}

func PingResponse() *discordgo.InteractionResponse {
	response := response.
		NewMessageResponse().
		SetInteractionResponseData(PingData())

	return response.InteractionResponse
}

func PingData() *discordgo.InteractionResponseData {
	data := response.NewResponseData("pong")

	return data.InteractionResponseData
}
