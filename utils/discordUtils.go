package utils

import "github.com/bwmarrin/discordgo"

func GetGuild(session *discordgo.Session, guildID string) *discordgo.Guild {
	guild, _ := session.Guild(guildID)
	return guild
}
