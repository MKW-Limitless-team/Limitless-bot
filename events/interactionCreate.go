package events

import (
	"limitless-bot/responses"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	// Handle interaction type
	if interaction.Type == discordgo.InteractionApplicationCommand && interaction.GuildID != "" {
		cmd := interaction.ApplicationCommandData().Name
		response = responses.CommandResponses[cmd](session, interaction)

	} else if interaction.Type == discordgo.InteractionMessageComponent && interaction.GuildID != "" { // these are for button interactions
		switch customID := interaction.Interaction.MessageComponentData().CustomID; customID {
		case responses.PREVIOUS_BUTTON:
			response = responses.IncPage(session, interaction)
		case responses.HOME_BUTTON:
			response = responses.LeaderBoardResponse(session, interaction)
		case responses.NEXT_BUTTON:
			response = responses.IncPage(session, interaction)
		}
		response.Type = discordgo.InteractionResponseUpdateMessage
	} else if interaction.Type == discordgo.InteractionModalSubmit && interaction.GuildID != "" {
		switch customID := interaction.ModalSubmitData().CustomID; customID {
		}
	}

	if response == nil {
		response = &discordgo.InteractionResponse{
			Data: &discordgo.InteractionResponseData{
				Content: "No response for this interaction type is registered",
			},
			Type: discordgo.InteractionResponseChannelMessageWithSource,
		}
	}
	_ = session.InteractionRespond(interaction.Interaction, response)
}
