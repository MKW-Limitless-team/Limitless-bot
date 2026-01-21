package responses

import (
	"limitless-bot/command"
	"limitless-bot/response"
	"limitless-bot/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func TrackFolderAutoComplete(session *discordgo.Session, interaction *discordgo.InteractionCreate, focusedOption *discordgo.ApplicationCommandInteractionDataOption) *discordgo.InteractionResponse {
	response := response.NewAutoCompleteResponse().
		SetResponseData(TrackFolderAutoCompleteData(focusedOption))

	return response.InteractionResponse
}

func TrackFolderAutoCompleteData(focusedOption *discordgo.ApplicationCommandInteractionDataOption) *discordgo.InteractionResponseData {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0)
	value := focusedOption.StringValue()

	for track, folderName := range utils.FolderNames {
		if strings.Contains(strings.ToLower(track), strings.ToLower(value)) {
			choice := command.NewOptionChoice(track, folderName)
			choices = append(choices, choice.ApplicationCommandOptionChoice)
		}

		if len(choices) > 5 {
			break
		}
	}

	return response.NewAutoCompleteData(choices).InteractionResponseData
}
