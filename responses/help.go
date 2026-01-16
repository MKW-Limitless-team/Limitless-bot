package responses

import (
	"fmt"
	"limitless-bot/commands"
	"limitless-bot/components/embed"
	"limitless-bot/response"
	"limitless-bot/utils"

	"github.com/bwmarrin/discordgo"
)

func HelpResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.
		NewMessageResponse().
		SetResponseData(HelpData(session, interaction))

	return response.InteractionResponse
}

func HelpData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	embed := embed.NewRichEmbed("**Commands**", "All the info on the bot's commands ", 0xff00e4)

	for _, command := range commands.GlobalCommands {
		if command.DefaultMemberPermissions != nil && utils.HasPermission(interaction.Member, *command.DefaultMemberPermissions) {
			embed.AddField(fmt.Sprintf("**/%s**", command.Name), command.Description, false)
		}
	}

	data := response.NewResponseData("").AddEmbed(embed)
	return data.InteractionResponseData
}
