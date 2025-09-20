package responses

import (
	"encoding/base64"
	"fmt"
	"io"
	"limitless-bot/response"
	"limitless-bot/utils"
	"limitless-bot/utils/db"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func EditMiiResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(EditMiiData(interaction))

	return response.InteractionResponse
}

func EditMiiData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
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
			return response.NewResponseData("Failed to read mii").InteractionResponseData
		}
	}
	mii = base64.StdEncoding.EncodeToString(miiData)
	userID := interaction.Member.User.ID

	err := db.EditMii(mii, userID)

	if err != nil {
		return response.NewResponseData("Failed to edit mii").InteractionResponseData
	}

	data = response.NewResponseData(fmt.Sprintf("<@%s>'s mii has been changed", userID))

	return data.InteractionResponseData
}
