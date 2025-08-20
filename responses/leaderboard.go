package responses

import (
	"fmt"
	"limitless-bot/components/button"
	e "limitless-bot/components/embed"
	"limitless-bot/response"
	"limitless-bot/utils"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	PREVIOUS_BUTTON = "previous_button"
	NEXT_BUTTON     = "next_button"
	HOME_BUTTON     = "home_button"
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
	embed := e.NewRichEmbed("**Leaderboard**", "Personal Eval", 0xffd700)
	embed.SetThumbnail(guild.IconURL(""))

	pageNum := page + 1
	embed.SetFooter(fmt.Sprintf("Page : %d", pageNum), "")

	playerData := utils.SortByMMR(utils.GetPlayerData())
	playersPerPage := 10
	start := (page) * playersPerPage
	end := false

	for i := range playersPerPage {
		index := start + i
		if index < len(playerData) {
			player := playerData[index]

			embed.AddField("", fmt.Sprintf("**%d.** %s: %5.f", index+1, player.Name, player.Mmr), false)
			println(fmt.Sprintf("**%d.** %s:%5.f", index+1, player.Name, player.Mmr))
		} else {
			end = true
			embed.AddField("End", ":rewind: :regional_indicator_b: :regional_indicator_a: :regional_indicator_c: :regional_indicator_k: ", false)
			break
		}
	}

	data.AddEmbed(embed)

	actionRow := e.NewActionRow()

	previousButton := button.NewBasicButton("Previous", PREVIOUS_BUTTON, discordgo.PrimaryButton, (page == 0))
	homeButton := button.NewBasicButton("Home", HOME_BUTTON, discordgo.SecondaryButton, false)
	nextButton := button.NewBasicButton("Next", NEXT_BUTTON, discordgo.PrimaryButton, end)

	actionRow.Components = append(actionRow.Components, previousButton)
	actionRow.Components = append(actionRow.Components, homeButton)
	actionRow.Components = append(actionRow.Components, nextButton)

	data.AddActionRow(actionRow)

	return data.InteractionResponseData
}

func IncPage(session *discordgo.Session, interaction *discordgo.InteractionCreate, inc int) *discordgo.InteractionResponse {
	message := interaction.Message
	embed := message.Embeds[0]
	pageStr := strings.ReplaceAll(strings.Split(embed.Footer.Text, ":")[1], " ", "")
	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
	page := int(pageNum + int64(inc) - 1)

	if err != nil {
		return response.NewMessageResponse().
			SetInteractionResponseData(response.NewResponseData("Error changing page").InteractionResponseData).InteractionResponse
	}

	return LeaderBoardResponse(session, interaction, page)
}
