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
	VISIT_SITE_BUTTON = "visit_site_button"
)

func TableRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewModalResponse().
		SetResponseData(TableFormData(TABLE_SUBMIT, ""))

	return response.InteractionResponse
}

func EditTableRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	data := ""
	params, err := utils.GetURLParams(interaction.Message.Content)

	if err == nil {
		data = params["data"]
	}

	response := response.NewModalResponse().
		SetResponseData(TableFormData(EDIT_TABLE_SUBMIT, data))

	return response.InteractionResponse
}

func TableFormData(id string, value string) *discordgo.InteractionResponseData {
	data := response.NewFormData("Table data", id)

	actionRow := components.NewActionRow()
	textfield := modal.NewTextArea("Table Data", "data", true)
	textfield.Value = value

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
	edit_button := button.NewBasicButton("Edit", TABLE_EDIT_BUTTON, discordgo.PrimaryButton, false)
	url_button := button.NewLinkButton("Visit Site", "https://gb.hlorenzi.com/table?data="+url.QueryEscape(tableData), "🔗")

	data.AddComponent(actionRow)
	actionRow.AddComponent(edit_button)
	actionRow.AddComponent(url_button)

	return data.InteractionResponseData
}
