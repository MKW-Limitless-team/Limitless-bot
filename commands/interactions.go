package commands

import (
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	// Handle interaction type
	if interaction.Type == discordgo.InteractionApplicationCommand && interaction.GuildID != "" {
		switch cmd := interaction.ApplicationCommandData().Name; cmd {
		case "help":
			response = HelpResponse(session, interaction)
		case "ping":
			response = PingResponse()
		}
	} else if interaction.Type == discordgo.InteractionMessageComponent && interaction.GuildID != "" { // these are for button interactions
		switch customID := interaction.Interaction.MessageComponentData().CustomID; customID {
		}
		// response.Type = discordgo.InteractionResponseUpdateMessage
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
