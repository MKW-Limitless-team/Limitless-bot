package events

import (
	"limitless-bot/commands"
	"limitless-bot/responses"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	// Handle interaction type
	if interaction.Type == discordgo.InteractionApplicationCommand && interaction.GuildID != "" {
		switch cmd := interaction.ApplicationCommandData().Name; cmd {
		case commands.HELP_COMMAND:
			response = responses.HelpResponse(session, interaction)
		case commands.PING_COMMAND:
			response = responses.PingResponse()
		case commands.LEADERBOARD_COMMAND:
			response = responses.LeaderBoardResponse(session, interaction, 0)
		case commands.REGISTER_COMMAND:
			response = responses.RegistrationFormResponse()
		}
	} else if interaction.Type == discordgo.InteractionMessageComponent && interaction.GuildID != "" { // these are for button interactions
		switch customID := interaction.Interaction.MessageComponentData().CustomID; customID {
		case responses.PREVIOUS_BUTTON:
			response = responses.IncPage(session, interaction, -1)
		case responses.HOME_BUTTON:
			response = responses.LeaderBoardResponse(session, interaction, 0)
		case responses.NEXT_BUTTON:
			response = responses.IncPage(session, interaction, 1)
		}
		response.Type = discordgo.InteractionResponseUpdateMessage
	} else if interaction.Type == discordgo.InteractionModalSubmit && interaction.GuildID != "" {
		switch customID := interaction.ModalSubmitData().CustomID; customID {
		case responses.REGISTRATION_FORM:
			response = responses.RegistrationResponse(interaction)
		}
	} else {
		response = &discordgo.InteractionResponse{
			Data: &discordgo.InteractionResponseData{
				Content: "No response for this interaction type is registered",
			},
			Type: discordgo.InteractionResponseChannelMessageWithSource,
		}
	}
	_ = session.InteractionRespond(interaction.Interaction, response)
}
