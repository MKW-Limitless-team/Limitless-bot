package responses

import (
	"fmt"
	"io"
	"limitless-bot/commands"
	e "limitless-bot/components/embed"
	"limitless-bot/response"
	"limitless-bot/utils"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nwoik/generate-mii/rkg"
)

func RKGResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().SetResponseData(Detail(interaction))

	return response.InteractionResponse
}

func Detail(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	file := utils.GetAttachment(interaction)
	url := file.URL
	resp, err := http.Get(url)

	if err != nil {
		return response.NewResponseData("Failed to get file").InteractionResponseData
	}

	if !strings.HasSuffix(file.Filename, ".rkg") {
		return response.NewResponseData("The file must be a **.rkg** file").InteractionResponseData
	}

	rkgData, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return response.NewResponseData("Error reading file").InteractionResponseData
	}
	data := response.NewResponseData("")
	options := interaction.ApplicationCommandData().Options
	track := utils.GetOption(options, commands.TRACK_OPTION_NAME)
	readable := rkg.ConvertRkg(rkg.ParseRKG(rkgData))

	rkgEmbed := RkgEmbed(interaction, track.StringValue(), readable, url)

	data.AddEmbed(rkgEmbed)

	return data.InteractionResponseData
}

func RkgEmbed(interaction *discordgo.InteractionCreate, track string, readable *rkg.ReadbleRKG, fileUrl string) *e.Embed {
	header := readable.Header
	embed := e.NewRichEmbed(track, "", 0xcb2b83)

	folderName, ok := utils.FolderNames[track]
	if ok {
		embed.AddField("", fmt.Sprintf("**Folder Name:** `%s`", folderName), false)
	}

	var courseDetails strings.Builder

	fmt.Fprint(&courseDetails, "**`Course Details:`**\n")
	fmt.Fprintf(&courseDetails, "**Character:** `%s`\n", header.Character)
	fmt.Fprintf(&courseDetails, "**Vehicle:** `%s`\n", header.Vehicle)
	fmt.Fprintf(&courseDetails, "**DriftType:** `%s`\n", header.DriftType)
	fmt.Fprintf(&courseDetails, "**Controller:** `%s`\n", header.Controller)
	fmt.Fprintf(&courseDetails, "**GhostType:** `%s`\n", header.GhostType)

	embed.AddField("", courseDetails.String(), false)

	var timeDetails strings.Builder

	finish := header.FinishTime
	fmt.Fprint(&timeDetails, "**`Time Details:`**\n")

	finishTime := "**Finish Time:** `%s:%s:%s`\n"
	fmt.Fprintf(&timeDetails, finishTime, strTime(finish.Minutes), strTime(finish.Seconds), strTime(finish.Milliseconds))
	fmt.Fprintf(&timeDetails, "%s\n", divider(finishTime))
	for index, lap := range header.Laps {
		fmt.Fprintf(&timeDetails, "**Lap %d:** `%s:%s:%s`\n", index+1, strTime(lap.Minutes), strTime(lap.Seconds), strMs(lap.Milliseconds))
	}

	embed.AddField("", timeDetails.String(), false)

	embed.AddField("", fmt.Sprintf("**File:** %s", fileUrl), false)

	embed.URL = fileUrl

	date := fmt.Sprintf("Date created at: %d/%d/%d", header.Day, header.Month, header.Year)
	embed.SetFooter(date, interaction.Member.AvatarURL(""))
	return embed
}

func strTime(t int) string {
	str := fmt.Sprintf("%d", t)

	if len(str) == 1 {
		str = "0" + str
	}

	return str
}

func strMs(t int) string {
	str := fmt.Sprintf("%d", t)

	if len(str) == 1 {
		str = "00" + str
	} else if len(str) == 2 {
		str = "0" + str
	}

	return str
}

func divider(sample string) string {
	var divide strings.Builder

	for range len(sample) {
		divide.WriteString("-")
	}

	return divide.String()
}
