package responses

import (
	"encoding/json"
	"fmt"
	"limitless-bot/response"
	"limitless-bot/utils"
	"log"
	"net/http"
	"strings"

	"github.com/MKW-Limitless-team/limitless-types/responses"
	"github.com/bwmarrin/discordgo"
)

func EditLicenseResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(EditLicenseData(session, interaction))

	return response.InteractionResponse
}

func EditLicenseData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	args := interaction.ApplicationCommandData().Options

	userID := interaction.Member.User.ID
	var name, fc string

	nameOption := utils.GetOption(args, "name")
	if nameOption != nil {
		name = nameOption.StringValue()
	}

	fcOption := utils.GetOption(args, "friend_code")
	if fcOption != nil {
		fc = fcOption.StringValue()
		fc = strings.ReplaceAll(fc, "-", "")
	}

	log.Printf("Editing license for: [%s]", userID)

	resp, err := http.Get(fmt.Sprintf("http://localhost:5000/edit?name=%s&discord_id=%s&friend_code=%s", name, userID, fc))

	if err != nil {
		return response.NewResponseData("Unable to edit license, contact admin").InteractionResponseData
	}

	var jsonResponse *responses.PlayerInfoResponse
	json.NewDecoder(resp.Body).Decode(&jsonResponse)

	if jsonResponse.Status == responses.Failure {
		return response.NewResponseData(jsonResponse.Message).InteractionResponseData
	}

	return LicenseData(session, interaction)
}
