package events

import (
	"limitless-bot/responses"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func OnMessage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	message := msg.Message

	if msg.Author.Bot && strings.Contains(msg.Content, "Seed:") {
		for _, emoji := range responses.LABELS {
			session.MessageReactionAdd(msg.ChannelID, message.ID, emoji)
		}
	}

	// attachments := message.Attachments
	// if !msg.Author.Bot && message.Content == "" && isTimeTrialSubmission(message.Attachments) {
	// 	ok := true
	// 	for _, attachment := range attachments {
	// 		resp, err := http.Get(attachment.URL)
	// 		if err == nil {
	// 			rkgData, err := io.ReadAll(resp.Body)
	// 			defer resp.Body.Close()
	// 			if err == nil {
	// 				readable := rkg.ConvertRkg(rkg.ParseRKG(rkgData))
	// 				embed := messeges.RkgEmbed(message, attachment, readable)
	// 				_, err := session.ChannelMessageSendEmbed(msg.ChannelID, embed)
	// 				if err != nil {
	// 					ok = false
	// 				}
	// 			}
	// 		}
	// 	}

	// 	if ok {
	// 		session.ChannelMessageDelete(msg.ChannelID, msg.ID)
	// 	}
	// }
}

func isTimeTrialSubmission(attachments []*discordgo.MessageAttachment) bool {
	result := true
	for _, attachment := range attachments {
		if !strings.HasSuffix(attachment.Filename, ".rkg") {
			result = false
		}
	}

	return result
}
