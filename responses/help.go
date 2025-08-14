package responses

import (
	"fmt"
	"limitless-bot/commands"
	"limitless-bot/components/embed"
	"limitless-bot/response"

	"github.com/bwmarrin/discordgo"
)

func HelpResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.
		NewMessageResponse().
		SetInteractionResponseData(HelpData(session, interaction))

	return response.InteractionResponse
}

func HelpData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	embed := embed.NewRichEmbed("**Commands**", "All the info on the bot's commands ", 0xff00e4)

	for _, command := range commands.GlobalCommands {
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
