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
	KICK_SUBMIT = "kick_submit"
)

type KickRequestSpec struct {
	Secret    string `json:"secret"`
	ProfileID uint32 `json:"pid"`
	Reason    string `json:"reason"`
}

func KickRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	hasRole := utils.HasRole(interaction.Member, globals.ADMIN_ROLE)
	var response *r.Response
	if hasRole {
		response = r.NewModalResponse().SetResponseData(KickForm())
	} else {
		response = r.NewMessageResponse().SetResponseData(r.NewResponseData("User lacks sufficient role to use the `/kick` command").InteractionResponseData)
	}

	return response.InteractionResponse
}

func KickForm() *discordgo.InteractionResponseData {
	data := r.NewFormData("Kick Form", KICK_SUBMIT)

	actionRow := components.NewActionRow()
	fc := modal.NewTextField("Friend-Code", "friend-code", "", true)
	actionRow.AddComponent(fc)
	data.AddComponent(actionRow)

	actionRow = components.NewActionRow()
	reason := modal.NewTextArea("Reason", "reason", true)
	actionRow.AddComponent(reason)
	data.AddComponent(actionRow)

	return data.InteractionResponseData
}

func KickResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := r.NewMessageResponse().SetResponseData(KickData(interaction))

	return response.InteractionResponse
}

func KickData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	submitData := interaction.ModalSubmitData()

	friendCode, _ := utils.GetSubmitDataValueByID(submitData, "friend-code")
	fc, err := strconv.Atoi(strings.ReplaceAll(friendCode, "-", ""))
	if err != nil {
		return r.NewResponseData("Invalid friend-code").InteractionResponseData
	}

	reason, _ := utils.GetSubmitDataValueByID(submitData, "reason")

	profileID := uint32(wwfc.FCToPid(uint64(fc)))
	resp, err := http.Get(fmt.Sprintf("http://localhost:5000/user?profile_id=%d", profileID))
	if err != nil {
		return r.NewResponseData("Can't verify if player exists. (contact admin)").InteractionResponseData
	}
	defer resp.Body.Close()

	var jsonResponse *responses.PlayerInfoResponse
	json.NewDecoder(resp.Body).Decode(&jsonResponse)

	if jsonResponse.Status == responses.Failure {
		return r.NewResponseData("User doesn't exist").InteractionResponseData
	}

	kickReqSpec := &KickRequestSpec{
		Secret:    globals.SECRET,
		ProfileID: uint32(jsonResponse.User.ProfileID),
		Reason:    reason,
	}

	marshalled, err := json.Marshal(kickReqSpec)
	if err != nil {
		return r.NewResponseData("Failed to form kick request").InteractionResponseData
	}

	resp, err = http.Post("http://localhost/api/kick", "application/json", bytes.NewBuffer(marshalled))
	if err != nil {
		return r.NewResponseData("Failed to kick user").InteractionResponseData
	}
	defer resp.Body.Close()

	var responseJson map[string]string
	json.NewDecoder(resp.Body).Decode(&responseJson)

	if _, ok := responseJson["success"]; ok {
		return r.NewResponseData(fmt.Sprintf("Kicked **%s** from Limitlink\nReason: `%s`", friendCode, reason)).InteractionResponseData
	}

	return r.NewResponseData("Failed to kick user").InteractionResponseData
}
