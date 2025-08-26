package responses

import (
	"limitless-bot/components"
	"limitless-bot/components/modal"
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

	actionRow := components.NewActionRow()
	format := modal.NewStringSelectMenu("format", 1, 1)
	format.
		AddMenuOption(modal.NewSelectMenuOption("FFA", "FFA")).
		AddMenuOption(modal.NewSelectMenuOption("2vs2", "2vs2")).
		AddMenuOption(modal.NewSelectMenuOption("3vs3", "3vs3")).
		AddMenuOption(modal.NewSelectMenuOption("4vs4", "4vs4")).
		AddMenuOption(modal.NewSelectMenuOption("5vs5", "5vs5")).
		AddMenuOption(modal.NewSelectMenuOption("6vs6", "6vs6"))

	actionRow.AddComponent(format)
	data.AddComponent(actionRow)

	return data.InteractionResponseData
}
