package responses

import (
	"limitless-bot/response"

	"github.com/bwmarrin/discordgo"
)

var (
	START_TOURNAMENT_FORM     = "start_tournament_form"
	START_TOURNAMENT_RESPONSE = "start_tournament_response"
)

func StartTournamentFormResponse() *discordgo.InteractionResponse {
	response := response.NewModalResponse().
		SetResponseData(StartTournamentFormData())

	return response.InteractionResponse
}

func StartTournamentFormData() *discordgo.InteractionResponseData {
	data := response.NewFormData("Start Tournament Form", START_TOURNAMENT_FORM)

	return data.InteractionResponseData
}
