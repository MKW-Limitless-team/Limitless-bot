package responses

import (
	"limitless-bot/response"

	"github.com/bwmarrin/discordgo"
)

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
