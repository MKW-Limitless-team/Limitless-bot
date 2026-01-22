package messeges

import (
	"fmt"
	e "limitless-bot/components/embed"
	"limitless-bot/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nwoik/generate-mii/rkg"
)

func RkgEmbed(message *discordgo.Message, attachment *discordgo.MessageAttachment, readable *rkg.ReadbleRKG) *discordgo.MessageEmbed {
	header := readable.Header
	embed := e.NewRichEmbed(header.Track, "", 0xcb2b83)

	folderName, ok := utils.FolderNames[header.Track]
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

	finishTime := "**Finish Time:** `%d:%d:%d`\n"
	fmt.Fprintf(&timeDetails, finishTime, finish.Minutes, finish.Seconds, finish.Milliseconds)
	fmt.Fprintf(&timeDetails, "%s\n", divider(finishTime))
	for index, lap := range header.Laps {
		fmt.Fprintf(&timeDetails, "**Lap %d:** `%d:%d:%d`\n", index+1, lap.Minutes, lap.Seconds, lap.Milliseconds)
	}

	embed.AddField("", timeDetails.String(), false)

	embed.AddField("", fmt.Sprintf("**File:** %s", attachment.URL), false)

	embed.URL = attachment.URL

	date := fmt.Sprintf("Date created at: %d/%d/%d", header.Day, header.Month, header.Year)
	embed.SetFooter(date, message.Author.AvatarURL(""))
	return embed.MessageEmbed
}

func divider(sample string) string {
	var divide strings.Builder

	for range len(sample) {
		divide.WriteString("-")
	}

	return divide.String()
}
