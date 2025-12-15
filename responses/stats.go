package responses

import (
	"encoding/json"
	"fmt"
	e "limitless-bot/components/embed"
	"limitless-bot/response"
	"limitless-bot/utils"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type Stats struct {
	OnlinePlayerCount int `json:"online"`
	ActivePlayerCount int `json:"active"`
	GroupCount        int `json:"groups"`
}

func OnlineResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	guild := utils.GetGuild(session, interaction.GuildID)

	response := response.
		NewMessageResponse().
		SetResponseData(OnlineData(guild))

	return response.InteractionResponse
}

func OnlineData(guild *discordgo.Guild) *discordgo.InteractionResponseData {
	data := response.NewResponseData("")
	embed := e.NewRichEmbed("**Online Status**", "", 0xd70ccf)
	embed.SetThumbnail(guild.IconURL(""))

	resp, err := http.Get("http://localhost/api/stats?game=mariokartwii")

	if err != nil {
		embed.AddField("Status:", "Could not fetch data from server", false)
	}

	stats := map[string]Stats{}
	err = json.NewDecoder(resp.Body).Decode(&stats)

	if err != nil {
		embed.AddField("Status:", "Failed to parse response from server", false)
	}

	global := stats["global"]

	embed.AddField("Online:", fmt.Sprintf("%d", global.OnlinePlayerCount), false)
	embed.AddField("Active:", fmt.Sprintf("%d", global.ActivePlayerCount), false)
	embed.AddField("Groups:", fmt.Sprintf("%d", global.GroupCount), false)

	data.AddEmbed(embed)
	return data.InteractionResponseData
}
