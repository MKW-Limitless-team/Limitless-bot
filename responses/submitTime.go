package responses

import (
	"io"
	"limitless-bot/response"
	"limitless-bot/utils"
	"limitless-bot/utils/crc"
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
	url := file.URL
	resp, err := http.Get(url)

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
	category := utils.GetOption(args, "category").StringValue()
	userID := interaction.Member.User.ID
	err = db.SubmitTime(rkgData, userID, category, url)
	viewSubmission(rkgData)

	if err != nil {
		return response.NewResponseData(err.Error()).InteractionResponseData
	}

	return data.InteractionResponseData
}

func viewSubmission(bytes []byte) {
	crc := crc.CRC(bytes)

	placement, playerData, err := db.GetTimeByCRC(crc)

	if err != nil {
		println(err)
		return
	}

	println(placement, playerData)
}
