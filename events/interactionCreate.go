package events

import (
	"limitless-bot/responses"
	"limitless-bot/utils"
	"strings"

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

		if interactionResp != nil && utils.HasPermission(member, interactionResp.Permission) {
			if strings.HasPrefix(customID, responses.USAGE_BUTTON) {
				err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{Type: discordgo.InteractionResponseDeferredMessageUpdate})
				if err != nil {
					println(err.Error())
					return
				}

				responseData := responses.UsagePageData(session, interaction)
				_, err = session.InteractionResponseEdit(interaction.Interaction, &discordgo.WebhookEdit{
					Content:    &responseData.Content,
					Embeds:     &responseData.Embeds,
					Components: &responseData.Components,
				})

				if err != nil {
					println(err.Error())
				}
				return
			}

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
		err := session.InteractionRespond(interaction.Interaction, response)

		if err != nil {
			println(err.Error())
		}
	}
}
