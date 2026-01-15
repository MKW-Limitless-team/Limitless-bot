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
}
