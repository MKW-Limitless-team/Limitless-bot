package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetGuild(session *discordgo.Session, guildID string) *discordgo.Guild {
	guild, _ := session.Guild(guildID)
	return guild
}

func GetSubmitDataValueByID(submitData discordgo.ModalSubmitInteractionData, id string) (string, error) {
	var value string

	for _, component := range submitData.Components {
		actionRow := component.(*discordgo.ActionsRow)
		for _, input := range actionRow.Components {
			switch inputType := input.Type(); inputType {
			case discordgo.TextInputComponent:
				textInput := input.(*discordgo.TextInput)
				if textInput.CustomID == id {
					return textInput.Value, nil
				}
			}
		}
	}

	return value, fmt.Errorf("No field found for ID: %s", id)
}
