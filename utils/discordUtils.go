package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func GetArgument(options []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, option := range options {
		if option.Name == name {
			return option
		}
	}
	return nil
}

func GetAttachment(interaction *discordgo.InteractionCreate) *discordgo.MessageAttachment {
	var file *discordgo.MessageAttachment

	for _, attachment := range interaction.ApplicationCommandData().Resolved.Attachments {
		file = attachment
	}

	return file
}

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

func HexToInt(hex string) int {
	hex = strings.ReplaceAll(hex, "#", "")
	num, err := strconv.ParseInt(hex, 16, 64)

	if err != nil {
		num = 0xffffff
	}

	return int(num)
}

func FlagEmoji(flag string) string {
	countryCode := strings.Replace(flag, "[", "", -1)
	countryCode = strings.Replace(countryCode, "]", "", -1)
	countryCode = strings.ToUpper(countryCode)
	if len(countryCode) != 2 {
		return ""
	}

	runes := []rune(countryCode)
	return string([]rune{
		runes[0] + 127397,
		runes[1] + 127397,
	})
}
