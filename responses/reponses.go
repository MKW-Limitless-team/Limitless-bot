package responses

import (
	"limitless-bot/commands"

	"github.com/bwmarrin/discordgo"
)

var CommandResponses = map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse{}
var InteractionResponses = map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse{}

func RegisterResponses() {
	// Add command responses here
	CommandResponses[commands.HELP_COMMAND] = HelpResponse
	CommandResponses[commands.PING_COMMAND] = PingResponse
	CommandResponses[commands.LEADERBOARD_COMMAND] = LeaderBoardResponse
	CommandResponses[commands.REGISTER_COMMAND] = RegistrationResponse
	CommandResponses[commands.SUBMIT_TIME_COMMAND] = SubmitTimeResponse
	CommandResponses[commands.EDIT_MII_COMMAND] = EditMiiResponse
	CommandResponses[commands.LICENSE_COMMAND] = LicenseResponse
	CommandResponses[commands.ONLINE_COMMAND] = OnlineResponse

	// Add interaction reponses here
	InteractionResponses[PREVIOUS_BUTTON] = IncPage
	InteractionResponses[HOME_BUTTON] = LeaderBoardResponse
	InteractionResponses[NEXT_BUTTON] = IncPage

	// Add modal responses here
}
