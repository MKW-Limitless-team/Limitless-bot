package responses

import (
	"limitless-bot/components"
	"limitless-bot/components/modal"
	"limitless-bot/response"
	"limitless-bot/utils"
	"strconv"
	"strings"

	"github.com/MKW-Limitless-team/limitless-types/wwfc"
	"github.com/bwmarrin/discordgo"
)

var (
	BAN_SUBMIT = "ban_submit"
)

type BanRequestSpec struct {
	Secret       string `json:"secret"`
	ProfileID    uint32 `json:"pid"`
	Days         uint64 `json:"days"`
	Hours        uint64 `json:"hours"`
	Minutes      uint64 `json:"minutes"`
	Tos          bool   `json:"tos"`
	Reason       string `json:"reason"`
	ReasonHidden string `json:"reason_hidden"`
	Moderator    string `json:"moderator"`
}

func BanRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewModalResponse().SetResponseData(BanForm())

	return response.InteractionResponse
}

func BanForm() *discordgo.InteractionResponseData {
	data := response.NewFormData("Ban Form", BAN_SUBMIT)

	actionRow := components.NewActionRow()
	fc := modal.NewTextField("Friend-Code", "friend-code", "", true)
	actionRow.AddComponent(fc)
	data.AddComponent(actionRow)

	actionRow = components.NewActionRow()
	days := modal.NewTextField("Days", "days", "", false)
	actionRow.AddComponent(days)
	data.AddComponent(actionRow)

	actionRow = components.NewActionRow()
	tos := modal.NewStringSelectMenu("Tos", 1, 1)
	tos.AddMenuOption(modal.NewSelectMenuOption("False", "false"))
	tos.AddMenuOption(modal.NewSelectMenuOption("True", "true"))
	actionRow.AddComponent(tos.SelectMenu)
	data.AddComponent(actionRow)

	actionRow = components.NewActionRow()
	reason := modal.NewTextArea("Reason", "reason", false)
	actionRow.AddComponent(reason)
	data.AddComponent(actionRow)

	return data.InteractionResponseData
}

func BanResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewModalResponse().SetResponseData(BanData(session, interaction))

	return response.InteractionResponse
}

func BanData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	data := response.NewResponseData("")
	submitData := interaction.ModalSubmitData()

	friendCode, _ := utils.GetSubmitDataValueByID(submitData, "friend-code")
	fc, err := strconv.Atoi(strings.ReplaceAll(friendCode, "-", ""))

	if err != nil {
		return response.NewResponseData("Invalid friend-code").InteractionResponseData
	}

	daysStr, _ := utils.GetSubmitDataValueByID(submitData, "days")
	days, err := strconv.Atoi(daysStr)

	if err != nil {
		days = 365
	}

	// tosStr, _ := utils.GetSubmitDataValueByID(submitData, "tos")
	// tos := strconv.ParseBool()

	banReqSpec := &BanRequestSpec{}
	banReqSpec.ProfileID = uint32(wwfc.FCToPid(uint64(fc)))
	banReqSpec.Days = uint64(days)

	return data.InteractionResponseData
}
