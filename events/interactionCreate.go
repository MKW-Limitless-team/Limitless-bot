package events

import (
	"limitless-bot/responses"
	"limitless-bot/utils"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	// Handle interaction type
	switch interaction.Type {
	case discordgo.InteractionApplicationCommand:
		cmd := interaction.ApplicationCommandData().Name
		responseFunc, ok := responses.CommandResponses[cmd]

		if ok {
			response = responseFunc(session, interaction)
		}

	case discordgo.InteractionMessageComponent:
		member := interaction.Member
		customID := interaction.Interaction.MessageComponentData().CustomID
		interactionResp := responses.GetInteraction(customID, responses.InteractionResps)

		if utils.HasPermission(member, interactionResp.Permission) {
			response = interactionResp.Respond(session, interaction)
		}

	case discordgo.InteractionModalSubmit:
		customID := interaction.ModalSubmitData().CustomID
		responseFunc, ok := responses.ModalResponses[customID]

		if ok {
			response = responseFunc(session, interaction)
		}

	case discordgo.InteractionApplicationCommandAutocomplete:
		data := interaction.ApplicationCommandData()
		focusedOption := utils.GetFocusedOption(data.Options)

		responseFunc, ok := responses.AutoCompleteResponses[focusedOption.Name]
		if ok && focusedOption.StringValue() != "" {
			response = responseFunc(session, interaction, focusedOption)
		}
	}

	if response != nil {
		_ = session.InteractionRespond(interaction.Interaction, response)
	}
}
