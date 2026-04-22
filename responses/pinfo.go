package responses

import (
	"bytes"
	"encoding/json"
	"fmt"
	e "limitless-bot/components/embed"
	"limitless-bot/globals"
	r "limitless-bot/response"
	"limitless-bot/utils"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

type PinfoRequestSpec struct {
	Secret    string `json:"secret"`
	ProfileID uint32 `json:"pid"`
}

type PinfoPlayer struct {
	ProfileID uint32 `json:"profile_id"`
	MiiName   string `json:"mii_name"`
	MiiData   string `json:"mii_data"`
	OpenHost  bool   `json:"open_host"`
	Banned    bool   `json:"banned"`
	DiscordID string `json:"discord_id"`
}

type PinfoAPIResponse struct {
	Player  PinfoPlayer `json:"player"`
	Success bool        `json:"success"`
	Error   string      `json:"error"`
}

func PinfoResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := r.NewMessageResponse().
		SetResponseData(PinfoData(interaction))

	return response.InteractionResponse
}

func PinfoData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	options := interaction.ApplicationCommandData().Options
	friendCode := utils.GetOption(options, "friend_code").StringValue()
	friendCode = utils.NormalizeFriendCode(friendCode)

	profileID, err := utils.FriendCodeToPID(friendCode)
	if err != nil {
		return r.NewResponseData("Invalid friend-code").InteractionResponseData
	}

	reqBody, err := json.Marshal(&PinfoRequestSpec{
		Secret:    globals.SECRET,
		ProfileID: profileID,
	})
	if err != nil {
		return r.NewResponseData("Failed to form pinfo request").InteractionResponseData
	}

	resp, err := http.Post("http://localhost/api/pinfo", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return r.NewResponseData("Unable to get player info").InteractionResponseData
	}
	defer resp.Body.Close()

	var apiResponse PinfoAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return r.NewResponseData("Unable to decode pinfo response").InteractionResponseData
	}

	if !apiResponse.Success {
		if apiResponse.Error == "" {
			apiResponse.Error = "Failed to find user in the database"
		}

		return r.NewResponseData(apiResponse.Error).InteractionResponseData
	}

	data := r.NewResponseData("")
	data.AddEmbed(PinfoEmbed(utils.FormatFriendCode(friendCode), apiResponse.Player))
	return data.InteractionResponseData
}

func PinfoEmbed(friendCode string, player PinfoPlayer) *e.Embed {
	embed := e.NewRichEmbed(fmt.Sprintf("Player info for friend code %s", friendCode), "", 0xf08aac)

	embed.AddField("Profile ID", strconv.FormatUint(uint64(player.ProfileID), 10), false)

	if player.MiiName == "" {
		embed.AddField("Mii Name", "Unknown", false)
	} else {
		embed.AddField("Mii Name", player.MiiName, false)
	}

	embed.AddField("Open Host", strconv.FormatBool(player.OpenHost), false)
	embed.AddField("Banned", strconv.FormatBool(player.Banned), false)

	discordIDValue := "Not linked"
	if player.DiscordID != "" {
		discordIDValue = fmt.Sprintf("<@%s>", player.DiscordID)
	}
	embed.AddField("Discord ID", discordIDValue, false)

	if player.MiiData != "" {
		embed.SetThumbnail("https://mii-unsecure.ariankordi.net/miis/image.png?data=" + url.QueryEscape(player.MiiData) + "&expression=normal&shaderType=switch")
	}

	return embed
}
