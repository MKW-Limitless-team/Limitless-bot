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
		customID := interaction.Interaction.MessageComponentData().CustomID
		response = responses.InteractionResponses[customID](session, interaction)
		response.Type = discordgo.InteractionResponseUpdateMessage

	} else if interaction.Type == discordgo.InteractionModalSubmit && interaction.GuildID != "" {
		customID := interaction.ModalSubmitData().CustomID
		response = responses.ModalResponses[customID](session, interaction)
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
