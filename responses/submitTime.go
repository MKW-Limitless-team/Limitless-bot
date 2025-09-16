package responses

import (
	"fmt"
	"io"
	"limitless-bot/response"
	"limitless-bot/utils"
	"net/http"

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
		return response.NewResponseData(fmt.Sprint(err.Error())).InteractionResponseData
	}

	rkg, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return response.NewResponseData(fmt.Sprint(err.Error())).InteractionResponseData
	}

	fmt.Printf("%b\n", rkg)

	return data.InteractionResponseData
}
