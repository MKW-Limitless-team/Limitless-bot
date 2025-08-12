package commands

import (
	"fmt"
	"limitless-bot/command"
	"limitless-bot/command/response"
	"limitless-bot/components/embed"

	"github.com/bwmarrin/discordgo"
)

func HelpCommand() *discordgo.ApplicationCommand {
	help := command.NewChatApplicationCommand("help", "Lists the bot's commands").
		SetDefaultMemberPermissions(discordgo.PermissionViewChannel)

	return help.ApplicationCommand
}

func HelpResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.
		NewMessageResponse().
		SetInteractionResponseData(HelpData(session, interaction))

	return response.InteractionResponse
}

func HelpData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	embed := embed.NewRichEmbed("**Commands**", "All the info on the bot's commands ", 0xff00e4)

	for _, command := range globalCommands {
		if hasPermission(interaction.Member, command.DefaultMemberPermissions) {
			embed.AddField(fmt.Sprintf("**/%s**", command.Name), command.Description, false)
		}
	}

	data := response.NewResponseData("").AddEmbed(embed)
	return data.InteractionResponseData
}

func hasPermission(member *discordgo.Member, requiredPermission *int64) bool {
	if member == nil || requiredPermission == nil {
		return false
	}

	return member.Permissions&(*requiredPermission) == *requiredPermission
}
