package responses

import (
	"fmt"
	"limitless-bot/command"
	"limitless-bot/commands"
	"limitless-bot/response"
	"limitless-bot/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func TrackFolderResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().SetResponseData(TrackFolderData(interaction))

	return response.InteractionResponse
}

func TrackFolderData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	track := utils.GetOption(interaction.ApplicationCommandData().Options, commands.TRACK_OPTION_NAME).StringValue()

	folder, ok := utils.FolderNames[track]

	if ok {
		return response.NewResponseData(fmt.Sprintf("**%s**: `%s`", track, folder)).InteractionResponseData
	} else {
		return response.NewResponseData(fmt.Sprintf("No track found for `%s`", track)).InteractionResponseData
	}
}

func TrackFolderAutoComplete(session *discordgo.Session, interaction *discordgo.InteractionCreate, focusedOption *discordgo.ApplicationCommandInteractionDataOption) *discordgo.InteractionResponse {
	response := response.NewAutoCompleteResponse().
		SetResponseData(TrackFolderAutoCompleteData(focusedOption))

	return response.InteractionResponse
}

func TrackFolderAutoCompleteData(focusedOption *discordgo.ApplicationCommandInteractionDataOption) *discordgo.InteractionResponseData {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0)
	value := focusedOption.StringValue()

	for track := range utils.FolderNames {
		if strings.Contains(strings.ToLower(track), strings.ToLower(value)) {
			choice := command.NewOptionChoice(track, track)
			choices = append(choices, choice.ApplicationCommandOptionChoice)
		}

		if len(choices) > 5 {
			break
		}
	}

	return response.NewAutoCompleteData(choices).InteractionResponseData
}
