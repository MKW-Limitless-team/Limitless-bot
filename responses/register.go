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

var (
	REGISTER = "register"
)

func Register(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(RegisterData(session, interaction))

	return response.InteractionResponse
}

func RegisterData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	args := interaction.ApplicationCommandData().Options

	userID := interaction.Member.User.ID
	fc := utils.GetOption(args, "friend_code").StringValue()
	fc = strings.ReplaceAll(fc, "-", "")

	log.Printf("Registering: %s - %s", interaction.Member.User.Username, fc)

	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/register?discord_id=%s&friend_code=%s", userID, fc))

	if err != nil {
		return response.NewResponseData("Unable to register player, contact admin").InteractionResponseData
	}

	var jsonResponse *responses.PlayerInfoResponse
	json.NewDecoder(resp.Body).Decode(&jsonResponse)

	if jsonResponse.Status == responses.Failure {
		return response.NewResponseData(jsonResponse.Message).InteractionResponseData
	}

	return response.NewResponseData(jsonResponse.Message + ". Use /license to see your profile").InteractionResponseData
}
