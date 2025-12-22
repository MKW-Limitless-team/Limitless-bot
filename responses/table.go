package responses

import (
	"fmt"
	"limitless-bot/components"
	e "limitless-bot/components/embed"
	"limitless-bot/components/modal"
	"limitless-bot/response"
	"limitless-bot/utils"
	"limitless-bot/utils/api"
	"strings"

	"github.com/MKW-Limitless-team/limitless-types/table"
	"github.com/bwmarrin/discordgo"
)

var (
	TABLE_SUBMIT = "table_submit"
)

func TableRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewModalResponse().
		SetResponseData(TableFormData())

	return response.InteractionResponse
}

func TableFormData() *discordgo.InteractionResponseData {
	data := response.NewFormData("Table data", TABLE_SUBMIT)

	actionRow := components.NewActionRow()
	textfield := modal.NewTextArea("Table Data", "data", true)

	actionRow.AddComponent(textfield)
	data.AddComponent(actionRow)

	return data.InteractionResponseData
}

func TableResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	guild := utils.GetGuild(session, interaction.GuildID)
	response := response.NewMessageResponse().
		SetResponseData(TableData(interaction, guild))

	return response.InteractionResponse
}

func TableData(interaction *discordgo.InteractionCreate, guild *discordgo.Guild) *discordgo.InteractionResponseData {
	data := response.NewResponseData("")
	submitData := interaction.ModalSubmitData()
	tableData, err := utils.GetSubmitDataValueByID(submitData, "data")

	if err != nil {
		data = response.NewResponseData("Couldn't get form data")
	}

	table, err := api.GetTable(tableData)
	data.SetContent(fmt.Sprintf("# %s • %s", table.Title, table.Date))

	if err != nil {
		data = response.NewResponseData("Couldn't get form data")
	} else {
		for _, group := range table.Groups {
			data.AddEmbed(TableEmbed(group))
		}
	}

	return data.InteractionResponseData
}

func TableEmbed(group *table.Group) *e.Embed {
	title := fmt.Sprintf("**%s**", group.Name)
	embed := e.NewRichEmbed(title, group.Desc, utils.HexToInt(group.Color))

	names := []string{}
	flags := []string{}
	scores := make([]string, len(group.Players))

	for i, player := range group.Players {
		names = append(names, player.Name)
		flags = append(flags, utils.FlagEmoji(player.Flag))

		scoreStr := ""

		for index, score := range player.Scores {
			if index == 0 {
				pScore := fmt.Sprintf("%d", score)
				for len(pScore) < 5 {
					pScore += " "
				}
				scoreStr += fmt.Sprintf("%s", pScore)
			} else {
				pScore := fmt.Sprintf("%d", score)
				for len(pScore) < 5 {
					pScore += " "
				}
				scoreStr += fmt.Sprintf("| %s ", pScore)
			}
		}

		if player.Penalty != 0 {
			scoreStr += fmt.Sprintf("| %d", player.Penalty)
		}

		scores[i] = fmt.Sprintf("**%s**", scoreStr)
	}

	embed.AddField("Players", strings.Join(names, "\n"), true)
	embed.AddField("Flags", strings.Join(flags, "\n"), true)
	embed.AddField("Scores", strings.Join(scores, "\n"), true)

	return embed
}
