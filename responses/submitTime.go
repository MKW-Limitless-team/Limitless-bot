package responses

import (
	"fmt"
	"io"
	"limitless-bot/response"
	"limitless-bot/utils"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	r "github.com/nwoik/generate-mii/rkg"
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

	rkg := r.ParseRKG(rkgData)
	readable := r.ConvertRkg(rkg)

	fmt.Println(readable.Header)

	return data.InteractionResponseData
}
