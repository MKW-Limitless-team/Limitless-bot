package responses

import (
	"bytes"
	"encoding/json"
	"fmt"
	"limitless-bot/components"
	"limitless-bot/components/modal"
	"limitless-bot/globals"
	r "limitless-bot/response"
	"limitless-bot/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/MKW-Limitless-team/limitless-types/responses"
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
	hasRole := utils.HasRole(interaction.Member, globals.ADMIN_ROLE)
	var response *r.Response
	if hasRole {
		response = r.NewModalResponse().SetResponseData(BanForm())
	} else {
		response = r.NewMessageResponse().SetResponseData(r.NewResponseData("User lacks sufficient role to use the `/ban` command").InteractionResponseData)
	}

	return response.InteractionResponse
}

func BanForm() *discordgo.InteractionResponseData {
	data := r.NewFormData("Ban Form", BAN_SUBMIT)

	actionRow := components.NewActionRow()
	fc := modal.NewTextField("Friend-Code", "friend-code", "", true)
	actionRow.AddComponent(fc)
	data.AddComponent(actionRow)

	actionRow = components.NewActionRow()
	days := modal.NewTextField("Days", "days", "", false)
	actionRow.AddComponent(days)
	data.AddComponent(actionRow)

	actionRow = components.NewActionRow()
	reason := modal.NewTextArea("Reason", "reason", true)
	actionRow.AddComponent(reason)
	data.AddComponent(actionRow)

	return data.InteractionResponseData
}

func BanResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := r.NewMessageResponse().SetResponseData(BanData(interaction))

	return response.InteractionResponse
}

func BanData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	data := r.NewResponseData("")
	submitData := interaction.ModalSubmitData()

	friendCode, _ := utils.GetSubmitDataValueByID(submitData, "friend-code")
	fc, err := strconv.Atoi(strings.ReplaceAll(friendCode, "-", ""))

	if err != nil {
		return r.NewResponseData("Invalid friend-code").InteractionResponseData
	}

	daysStr, _ := utils.GetSubmitDataValueByID(submitData, "days")
	days, err := strconv.Atoi(daysStr)

	if err != nil {
		days = 365
	}

	reason, _ := utils.GetSubmitDataValueByID(submitData, "reason")

	profileID := uint32(wwfc.FCToPid(uint64(fc)))
	resp, err := http.Get(fmt.Sprintf("http://localhost:5000/user?profile_id=%d", profileID))

	if err != nil {
		return r.NewResponseData("Can't verify if player exists. (contact admin)").InteractionResponseData
	}

	var jsonResponse *responses.PlayerInfoResponse
	json.NewDecoder(resp.Body).Decode(&jsonResponse)

	if jsonResponse.Status == responses.Failure {
		return r.NewResponseData("User doesn't exist").InteractionResponseData
	}

	user := jsonResponse.User
	if user.HasBan {
		return r.NewResponseData(fmt.Sprintf("User `{%s:%s}` already banned", user.LastInGameSn, friendCode)).InteractionResponseData
	}

	banReqSpec := &BanRequestSpec{
		Secret:    globals.SECRET,
		ProfileID: uint32(user.ProfileID),
		Days:      uint64(days),
		Tos:       true,
		Reason:    reason,
		Moderator: interaction.Member.DisplayName(),
	}

	marshalled, err := json.Marshal(banReqSpec)

	if err != nil {
		return r.NewResponseData("Failed to form ban request").InteractionResponseData
	}

	resp, err = http.Post("http://localhost/api/ban", "application/json", bytes.NewBuffer(marshalled))

	var responseJson map[string]string
	json.NewDecoder(resp.Body).Decode(&responseJson)

	_, ok := responseJson["success"]
	if ok {
		data.SetContent(fmt.Sprintf("Banned **%s** from Limitlink\nDuration: **%d**\nReason: `%s`", friendCode, days, reason))
	} else {
		data.SetContent("Failed to ban user")
	}

	return data.InteractionResponseData
}
