package responses

import (
	"fmt"
	"io"
	"limitless-bot/response"
	"limitless-bot/utils"
	"limitless-bot/utils/db"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	REGISTRATION_FORM    = "registration_form"
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
		SetResponseData(RegistrationResponseData(interaction))

	return response.InteractionResponse
}

func RegistrationResponseData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	var data *response.Data

	file := utils.GetAttachment(interaction)
	var mii string
	miiData := []byte{}
	if file != nil {
		resp, err := http.Get(file.URL)

		if err != nil {
			return response.NewResponseData("Failed to get file").InteractionResponseData
		}

		if !strings.HasSuffix(file.Filename, ".miigx") {
			return response.NewResponseData("The file must be a **.miigx** file").InteractionResponseData
		}

		data, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		miiData = append(miiData, data...)

		if err != nil {
			return response.NewResponseData("Failed to mii").InteractionResponseData
		}

	}
	mii = string(miiData)

	args := interaction.ApplicationCommandData().Options

	userID := interaction.Member.User.ID
	ign := utils.GetArgument(args, "ign").StringValue()
	fc := utils.GetArgument(args, "fc").StringValue()

	err := db.RegisterPlayer(ign, fc, userID, mii)

	if err != nil {
		switch err.Error() {
		case "sqlite3: constraint failed: UNIQUE constraint failed: playerdata.discord_id":
			return response.NewResponseData("User already registered").InteractionResponseData
		}
	}

	data = response.NewResponseData(fmt.Sprintf("<@%s> registered as %s \nFriend code: %s", userID, ign, fc))

	return data.InteractionResponseData
}
