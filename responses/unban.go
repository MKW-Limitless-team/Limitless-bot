package responses

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type UnbanRequestSpec struct {
	Secret    string `json:"secret"`
	ProfileID uint32 `json:"pid"`
}

func UnbanRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	hasRole := utils.HasRole(interaction.Member, globals.ADMIN_ROLE)
	var response *r.Response
	if hasRole {
		response = r.NewMessageResponse().SetResponseData(UnbanData(interaction))
	} else {
		response = r.NewMessageResponse().SetResponseData(r.NewResponseData("User lacks sufficient role to use the `/unban` command").InteractionResponseData)
	}

	return response.InteractionResponse
}

func UnbanData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	args := interaction.ApplicationCommandData().Options

	friendCode := utils.GetOption(args, "friend_code").StringValue()
	fc, err := strconv.Atoi(strings.ReplaceAll(friendCode, "-", ""))
	if err != nil {
		return r.NewResponseData("Invalid friend-code").InteractionResponseData
	}

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

	user := jsonResponse.User
	if !user.HasBan {
		return r.NewResponseData(fmt.Sprintf("User `{%s:%s}` already unbanned", user.LastInGameSn, friendCode)).InteractionResponseData
	}

	unbanReqSpec := &UnbanRequestSpec{
		Secret:    globals.SECRET,
		ProfileID: uint32(user.ProfileID),
	}

	marshalled, err := json.Marshal(unbanReqSpec)
	if err != nil {
		return r.NewResponseData("Failed to form unban request").InteractionResponseData
	}

	resp, err = http.Post("http://localhost/api/unban", "application/json", bytes.NewBuffer(marshalled))
	if err != nil {
		return r.NewResponseData("Failed to unban user").InteractionResponseData
	}
	defer resp.Body.Close()

	var responseJson map[string]string
	json.NewDecoder(resp.Body).Decode(&responseJson)

	if _, ok := responseJson["success"]; ok {
		return r.NewResponseData(fmt.Sprintf("Unbanned **%s** from Limitlink", friendCode)).InteractionResponseData
	}

	return r.NewResponseData("Failed to unban user").InteractionResponseData
}
