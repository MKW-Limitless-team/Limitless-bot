package responses

import (
	"io"
	"limitless-bot/response"
	"limitless-bot/utils"
	"limitless-bot/utils/db"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func SubmitTimeResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {

	response := response.
		NewMessageResponse().
		SetResponseData(SubmitTimeData(interaction))

	return response.InteractionResponse
}

func SubmitTimeData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	data := response.NewResponseData("Submitted")

	file := utils.GetAttachment(interaction)
	resp, err := http.Get(file.URL)

	if err != nil {
		return response.NewResponseData("Failed to get file").InteractionResponseData
	}

	if !strings.HasSuffix(file.Filename, ".rkg") {
		return response.NewResponseData("The file must be a **.rkg** file").InteractionResponseData
	}

	rkgData, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return response.NewResponseData("Error reading file").InteractionResponseData
	}

	args := interaction.ApplicationCommandData().Options
	category := utils.GetArgument(args, "category").StringValue()
	userID := interaction.Member.User.ID
	err = db.SubmitTime(rkgData, userID, category)

	if err != nil {
		return response.NewResponseData(err.Error()).InteractionResponseData
	}

	return data.InteractionResponseData
}
