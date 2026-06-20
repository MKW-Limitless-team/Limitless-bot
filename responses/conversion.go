package responses

import (
	"fmt"
	"limitless-bot/response"
	"limitless-bot/utils"
	"regexp"
	"strconv"
	"strings"

	"github.com/MKW-Limitless-team/limitless-types/wwfc"
	"github.com/bwmarrin/discordgo"
)

var friendCodeFormatRegex = regexp.MustCompile(`^[0-9]{12}$`)

func FCToPIDResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(FCToPIDData(interaction))

	return response.InteractionResponse
}

func PIDToFCResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(PIDToFCData(interaction))

	return response.InteractionResponse
}

func FCToPIDData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	options := interaction.ApplicationCommandData().Options
	friendCode := utils.GetOption(options, "friend_code").StringValue()
	fcNum, err := ParseFriendCode(friendCode)
	if err != nil {
		return response.NewResponseData("Invalid friend-code").InteractionResponseData
	}

	pid := wwfc.FCToPid(fcNum)

	return response.NewResponseData(fmt.Sprintf("**FC:** %s\n**PID:** %d", FormatFC(fcNum), pid)).InteractionResponseData
}

func PIDToFCData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	options := interaction.ApplicationCommandData().Options
	pidStr := utils.GetOption(options, "pid").StringValue()

	pid, err := strconv.ParseUint(pidStr, 10, 64)
	if err != nil || pid == 0 {
		return response.NewResponseData("Invalid PID").InteractionResponseData
	}

	fc := wwfc.PidToFC(pid)

	return response.NewResponseData(fmt.Sprintf("**PID:** %d\n**FC:** %s", pid, FormatFC(fc))).InteractionResponseData
}

func FormatFC(fc uint64) string {
	fcStr := strconv.FormatUint(fc, 10)
	for len(fcStr) < 12 {
		fcStr = "0" + fcStr
	}

	return fmt.Sprintf("%s-%s-%s", fcStr[:4], fcStr[4:8], fcStr[8:12])
}

func ParseFriendCode(friendCode string) (uint64, error) {
	fc := strings.ReplaceAll(friendCode, "-", "")
	fc = strings.TrimSpace(fc)

	if !friendCodeFormatRegex.MatchString(fc) {
		return 0, fmt.Errorf("invalid friend-code")
	}

	fcNum, err := strconv.ParseUint(fc, 10, 64)
	if err != nil || fcNum == 0 {
		return 0, fmt.Errorf("invalid friend-code")
	}

	return fcNum, nil
}
