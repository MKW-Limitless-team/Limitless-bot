package responses

import (
	"fmt"
	"limitless-bot/components"
	"limitless-bot/components/modal"
	"limitless-bot/response"
	"limitless-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	REGISTRATION_FORM    = "registration_form"
	REGISTRATON_RESPONSE = "registration_response"
)

func RegistrationFormResponse() *discordgo.InteractionResponse {
	response := response.NewModalResponse().
		SetResponseData(FormData())

	return response.InteractionResponse
}

func FormData() *discordgo.InteractionResponseData {
	data := response.NewFormData("Registration Form", REGISTRATION_FORM)

	actionRow := components.NewActionRow()
	textfield := modal.NewTextField("In-Game Name", "ign", "In-game name", true)

	data.AddComponent(actionRow)

	actionRow.AddComponent(textfield)

	return data.InteractionResponseData
}

func RegistrationResponse(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(RegistrationResponseData(interaction))

	return response.InteractionResponse
}

func RegistrationResponseData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	var data *response.Data

	fieldID := "ign"
	userID := interaction.Member.User.ID

	submitData := interaction.ModalSubmitData()

	value, err := utils.GetSubmitDataValueByID(submitData, fieldID)

	if err != nil {
		data = response.NewResponseData("User registered")
	} else {
		data = response.NewResponseData(fmt.Sprintf("<@%s> registered as %s", userID, value))
	}

	return data.InteractionResponseData
}
