package responses

import (
	"fmt"
	e "limitless-bot/components/embed"
	"limitless-bot/globals"
	"limitless-bot/response"
	"limitless-bot/utils"
	"limitless-bot/utils/ltrc"

	"github.com/bwmarrin/discordgo"
)

func LicenseResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(LicenseData(session, interaction))

	return response.InteractionResponse
}

func LicenseData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	var data *response.Data
	userID := interaction.Member.User.ID

	query := `SELECT name, friend_code, discord_id, mmr, mii 
					FROM playerdata
					WHERE discord_id = ?`
	rows, err := globals.GetConnection().Query(query, userID)

	if err != nil {
		return response.NewResponseData("Failed find license").InteractionResponseData
	}
	defer rows.Close()

	playerData := &ltrc.PlayerData{}

	// do if row.has next
	rows.Next()
	rows.Scan(&playerData.Name, &playerData.FriendCode, &playerData.DiscordID, &playerData.Mmr, &playerData.Mii)

	query = `SELECT mii 
					FROM playerdata
					WHERE discord_id = ?`

	mii, err := globals.GetConnection().Query(query, userID)

	if err != nil {
		return response.NewResponseData("Failed find license").InteractionResponseData
	}
	defer rows.Close()

	mii.Next()
	mii.Scan(&playerData.Mii)

	data = response.NewResponseData("")

	guild := utils.GetGuild(session, interaction.GuildID)
	embed := LicenseEmbed(playerData, guild)
	data.AddEmbed(embed)

	return data.InteractionResponseData
}

func LicenseEmbed(playerData *ltrc.PlayerData, guild *discordgo.Guild) *e.Embed {

	embed := e.NewRichEmbed(playerData.Name, "", 0xd70ccf)

	if playerData.Mii != "" {
		embed.SetThumbnail(fmt.Sprintf("https://mii-unsecure.ariankordi.net/miis/image.png?data=%s&expression=normal&cameraYRotate=30", playerData.Mii))
	} else {
		embed.SetThumbnail(guild.IconURL(""))
	}

	embed.AddField("", fmt.Sprintf("**Friend-Code:** %s", playerData.FriendCode), false)
	embed.AddField("", fmt.Sprintf("**MMR:** %d", playerData.Mmr), false)

	return embed
}
