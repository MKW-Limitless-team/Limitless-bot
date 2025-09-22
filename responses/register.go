package responses

import (
	"limitless-bot/response"
	"limitless-bot/utils"
	"limitless-bot/utils/db"

	"github.com/bwmarrin/discordgo"
)

var (
	// REGISTRATION_FORM    = "registration_form"
	REGISTRATON_RESPONSE = "registration_response"
)

// func RegistrationFormResponse() *discordgo.InteractionResponse {
// 	response := response.NewModalResponse().
// 		SetResponseData(RegistrationFormData())

// 	return response.InteractionResponse
// }

// func RegistrationFormData() *discordgo.InteractionResponseData {
// 	data := response.NewFormData("Registration Form", REGISTRATION_FORM)

// 	actionRow := components.NewActionRow()
// 	ign := modal.NewTextField("In-Game name", "ign", "In-game name", true)
// 	actionRow.AddComponent(ign)
// 	data.AddComponent(actionRow)

// 	actionRow = components.NewActionRow()
// 	fc := modal.NewTextField("Friend code", "fc", "Friend code", true)
// 	data.AddComponent(actionRow)
// 	actionRow.AddComponent(fc)

// 	return data.InteractionResponseData
// }

func RegistrationResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(RegistrationResponseData(session, interaction))

	return response.InteractionResponse
}

func RegistrationResponseData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	var data *response.Data

	args := interaction.ApplicationCommandData().Options

	userID := interaction.Member.User.ID
	ign := utils.GetArgument(args, "ign").StringValue()
	fc := utils.GetArgument(args, "fc").StringValue()

	err := db.RegisterPlayer(ign, fc, userID)

	if err != nil {
		return response.NewResponseData(err.Error()).InteractionResponseData
	}

	guild := utils.GetGuild(session, interaction.GuildID)
	playerData, _ := db.GetPlayer(userID)
	data = response.NewResponseData("").AddEmbed(LicenseEmbed(playerData, guild))

	return data.InteractionResponseData
}
