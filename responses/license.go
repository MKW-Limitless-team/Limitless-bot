package responses

import (
	"encoding/json"
	"fmt"
	e "limitless-bot/components/embed"
	"limitless-bot/response"
	"limitless-bot/utils"
	"net/http"

	"github.com/MKW-Limitless-team/limitless-types/ltrc"
	"github.com/MKW-Limitless-team/limitless-types/responses"
	"github.com/MKW-Limitless-team/limitless-types/wwfc"
	"github.com/bwmarrin/discordgo"
)

func LicenseResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(LicenseData(session, interaction))

	return response.InteractionResponse
}

func LicenseData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	var data *response.Data
	var userID string
	options := interaction.ApplicationCommandData().Options

	if len(options) == 0 {
		userID = interaction.Member.User.ID
	} else {
		user := utils.GetOption(options, "user").UserValue(session)
		userID = user.ID
	}

	resp, err := http.Get("http://localhost:8080/player?discord_id=" + userID)

	if err != nil {
		return response.NewResponseData("Unable to get license data").InteractionResponseData
	}

	data = response.NewResponseData("")

	var jsonResponse *responses.PlayerInfoResponse

	json.NewDecoder(resp.Body).Decode(&jsonResponse)

	if jsonResponse.Status == responses.Failure {
		return data.SetContent(jsonResponse.Message + "\nUse `/register` if you haven't registered already").InteractionResponseData
	}

	embed := LicenseEmbed(jsonResponse.PlayerData, jsonResponse.User)
	data.AddEmbed(embed)

	return data.InteractionResponseData
}

func LicenseEmbed(playerData *ltrc.PlayerData, user *wwfc.User) *e.Embed {
	embed := e.NewRichEmbed(user.LastInGameSn, fmt.Sprintf("<@%s>", playerData.DiscordID), 0xd70ccf)

	embed.AddField("", fmt.Sprintf("**Friend-Code:** %s", playerData.GetFC()), false)
	embed.AddField("", fmt.Sprintf("**MMR:** %d", playerData.Mmr), false)

	if user.FriendInfo != "" {
		embed.SetThumbnail(wwfc.ShowMii(user.GetMii()))
	} else {
		embed.AddField("", "`No mii found, use /edit-mii to set license icon`", false)
	}

	return embed
}
