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
	ign := modal.NewTextField("In-Game name", "ign", "In-game name", true)
	actionRow.AddComponent(ign)
	data.AddComponent(actionRow)

	actionRow = components.NewActionRow()
	fc := modal.NewTextField("Friend code", "fc", "Friend code", true)
	data.AddComponent(actionRow)
	actionRow.AddComponent(fc)

	return data.InteractionResponseData
}

func RegistrationResponse(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(RegistrationResponseData(interaction))

	return response.InteractionResponse
}

func RegistrationResponseData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	var data *response.Data

	ignID := "ign"
	fcID := "fc"
	userID := interaction.Member.User.ID

	submitData := interaction.ModalSubmitData()

	ign, _ := utils.GetSubmitDataValueByID(submitData, ignID)
	fc, err := utils.GetSubmitDataValueByID(submitData, fcID)

	if err != nil {
		data = response.NewResponseData("User not registered")
	} else {
		data = response.NewResponseData(fmt.Sprintf("<@%s> registered as %s \n%s", userID, ign, fc))
	}

	return data.InteractionResponseData
}
