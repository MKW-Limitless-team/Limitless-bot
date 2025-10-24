package responses

import (
	"limitless-bot/response"

	"github.com/bwmarrin/discordgo"
)

func PingResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.
		NewMessageResponse().
		SetResponseData(PingData())

	return response.InteractionResponse
}

func PingData() *discordgo.InteractionResponseData {
	data := response.NewResponseData("pong")

	return data.InteractionResponseData
}
