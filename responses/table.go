package responses

import (
	"limitless-bot/components"
	"limitless-bot/components/button"
	"limitless-bot/components/modal"
	"limitless-bot/response"
	"limitless-bot/utils"
	"net/url"

	"github.com/bwmarrin/discordgo"
)

var (
	TABLE_SUBMIT      = "table_submit"
	EDIT_TABLE_SUBMIT = "edit_table_submit"
	TABLE_EDIT_BUTTON = "table_edit_button"
)

func TableRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewModalResponse().
		SetResponseData(TableFormData(TABLE_SUBMIT))

	return response.InteractionResponse
}

func EditTableRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewModalResponse().
		SetResponseData(TableFormData(EDIT_TABLE_SUBMIT))

	return response.InteractionResponse
}

func TableFormData(id string) *discordgo.InteractionResponseData {
	data := response.NewFormData("Table data", id)

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

func EditTableResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	guild := utils.GetGuild(session, interaction.GuildID)
	response := response.NewMessageResponse().
		SetResponseData(TableData(interaction, guild))

	response.Type = discordgo.InteractionResponseUpdateMessage
	return response.InteractionResponse
}

func TableData(interaction *discordgo.InteractionCreate, guild *discordgo.Guild) *discordgo.InteractionResponseData {
	data := response.NewResponseData("")
	submitData := interaction.ModalSubmitData()
	tableData, err := utils.GetSubmitDataValueByID(submitData, "data")

	if err != nil {
		data = response.NewResponseData("Couldn't get form data")
	}

	data.SetContent("https://gb2.hlorenzi.com/table.png?data=" + url.QueryEscape(tableData))

	actionRow := components.NewActionRow()
	button := button.NewBasicButton("Edit", TABLE_EDIT_BUTTON, discordgo.PrimaryButton, false)

	data.AddComponent(actionRow)
	actionRow.AddComponent(button)

	return data.InteractionResponseData
}
