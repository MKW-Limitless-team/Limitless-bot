package responses

import (
	"limitless-bot/components"
	"limitless-bot/components/modal"
	"limitless-bot/response"
	"limitless-bot/utils"
	"net/url"

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

	data.SetContent("https://gb.hlorenzi.com/table.png?data=" + url.QueryEscape(tableData))

	return data.InteractionResponseData
}
