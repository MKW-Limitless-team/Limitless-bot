package responses

import (
	"fmt"
	"limitless-bot/components/embed"
	"limitless-bot/response"
	"limitless-bot/utils"

	"github.com/bwmarrin/discordgo"
)

func LeaderBoardResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, page int) *discordgo.InteractionResponse {
	guild := utils.GetGuild(session, interaction.GuildID)

	response := response.
		NewMessageResponse().
		SetInteractionResponseData(LeaderBoardData(guild, page))

	return response.InteractionResponse
}

func LeaderBoardData(guild *discordgo.Guild, page int) *discordgo.InteractionResponseData {
	data := response.NewResponseData("")
	embed := embed.NewRichEmbed("**Leaderboard**", "Personal Eval", 0xffd700)
	embed.SetThumbnail(guild.IconURL(""))
	playerData := utils.SortByMMR(utils.GetPlayerData())

	playersPerPage := 10

	for i := range playersPerPage * page {
		index := playersPerPage + i
		player := playerData[index]

		embed.AddField("", fmt.Sprintf("**%d.** %s: %5.f", index+1, player.Name, player.Mmr), false)
		println(fmt.Sprintf("**%d.** %s:%5.f", index+1, player.Name, player.Mmr))
	}

	data.AddEmbed(embed)

	return data.InteractionResponseData
}
