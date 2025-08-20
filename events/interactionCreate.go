package events

import (
	"limitless-bot/responses"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	// Handle interaction type
	if interaction.Type == discordgo.InteractionApplicationCommand && interaction.GuildID != "" {
		switch cmd := interaction.ApplicationCommandData().Name; cmd {
		case "help":
			response = responses.HelpResponse(session, interaction)
		case "ping":
			response = responses.PingResponse()
		case "leaderboard":
			response = responses.LeaderBoardResponse(session, interaction, 0)
		}
	} else if interaction.Type == discordgo.InteractionMessageComponent && interaction.GuildID != "" { // these are for button interactions
		switch customID := interaction.Interaction.MessageComponentData().CustomID; customID {
		case "previous_button":
			response = responses.IncPage(session, interaction, -1)
		case "home_button":
			response = responses.LeaderBoardResponse(session, interaction, 0)
		case "next_button":
			response = responses.IncPage(session, interaction, 1)
		}
		response.Type = discordgo.InteractionResponseUpdateMessage
	} else {
		response = &discordgo.InteractionResponse{
			Data: &discordgo.InteractionResponseData{
				Content: "This command cannot be used in direct messages.",
			},
			Type: discordgo.InteractionResponseChannelMessageWithSource,
		}
	}
	_ = session.InteractionRespond(interaction.Interaction, response)
}
