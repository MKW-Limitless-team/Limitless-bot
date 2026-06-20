package responses

import (
	"bytes"
	"encoding/json"
	"fmt"
	"limitless-bot/components"
	"limitless-bot/components/button"
	e "limitless-bot/components/embed"
	"limitless-bot/globals"
	r "limitless-bot/response"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/MKW-Limitless-team/limitless-types/wwfc"
	"github.com/bwmarrin/discordgo"
)

var (
	PINFO_MII_BUTTON = "pinfo_mii:"
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
		SetResponseData(PinfoData(session, interaction))

	return response.InteractionResponse
}

func PinfoData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	pid, errMessage := GetUsageRequest(session, interaction)
	if errMessage != "" {
		return r.NewResponseData(errMessage).InteractionResponseData
	}

	player, errMessage := GetPinfoPlayer(uint32(pid))
	if errMessage != "" {
		return r.NewResponseData(errMessage).InteractionResponseData
	}

	data := r.NewResponseData("")
	data.AddEmbed(PinfoEmbed(*player))

	if player.MiiData != "" {
		actionRow := components.NewActionRow()
		actionRow.AddComponent(button.NewBasicButton("Mii Data", PinfoMiiButtonID(player.ProfileID), discordgo.SecondaryButton, false))
		data.AddComponent(actionRow)
	}

	return data.InteractionResponseData
}

func PinfoMiiResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := r.NewMessageResponse().
		SetResponseData(PinfoMiiData(interaction))

	response.Data.Flags = discordgo.MessageFlagsEphemeral

	return response.InteractionResponse
}

func PinfoMiiData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	customID := interaction.MessageComponentData().CustomID
	pidStr := strings.TrimPrefix(customID, PINFO_MII_BUTTON)
	pid, err := strconv.ParseUint(pidStr, 10, 32)
	if err != nil || pid == 0 {
		return r.NewResponseData("Unable to get Mii data").InteractionResponseData
	}

	player, errMessage := GetPinfoPlayer(uint32(pid))
	if errMessage != "" {
		return r.NewResponseData(errMessage).InteractionResponseData
	}

	if player.MiiData == "" {
		return r.NewResponseData("No Mii data found").InteractionResponseData
	}

	return r.NewResponseData(fmt.Sprintf("**Mii Data**\n```text\n%s\n```", player.MiiData)).InteractionResponseData
}

func GetPinfoPlayer(pid uint32) (*PinfoPlayer, string) {
	reqBody, err := json.Marshal(&PinfoRequestSpec{
		Secret:    globals.SECRET,
		ProfileID: pid,
	})
	if err != nil {
		return nil, "Failed to form pinfo request"
	}

	resp, err := http.Post("http://localhost/api/pinfo", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, "Unable to get player info"
	}
	defer resp.Body.Close()

	var apiResponse PinfoAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, "Unable to decode pinfo response"
	}

	if !apiResponse.Success {
		if apiResponse.Error == "" {
			apiResponse.Error = "Failed to find user in the database"
		}

		return nil, apiResponse.Error
	}

	return &apiResponse.Player, ""
}

func PinfoMiiButtonID(pid uint32) string {
	return fmt.Sprintf("%s%d", PINFO_MII_BUTTON, pid)
}

func PinfoEmbed(player PinfoPlayer) *e.Embed {
	embed := e.NewRichEmbed(fmt.Sprintf("Player info for %s", FormatFC(wwfc.PidToFC(uint64(player.ProfileID)))), "", 0xf08aac)

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
