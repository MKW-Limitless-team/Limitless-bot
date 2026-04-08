package responses

import (
	"limitless-bot/commands"

	"github.com/bwmarrin/discordgo"
)

var (
	CommandResponses      = map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse{}
	InteractionResps      = make([]*InteractionResp, 0)
	ModalResponses        = map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse{}
	AutoCompleteResponses = map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate, focusedOption *discordgo.ApplicationCommandInteractionDataOption) *discordgo.InteractionResponse{}
)

type InteractionResp struct {
	ID         string
	Respond    func(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse
	Permission int64
}

func RegisterResponses() {
	// Add command responses here
	CommandResponses[commands.HELP_COMMAND] = HelpResponse
	CommandResponses[commands.PING_COMMAND] = PingResponse
	CommandResponses[commands.LEADERBOARD_COMMAND] = LeaderBoardResponse
	CommandResponses[commands.REGISTER_COMMAND] = Register
	CommandResponses[commands.SUBMIT_TIME_COMMAND] = SubmitTimeResponse
	CommandResponses[commands.LICENSE_COMMAND] = LicenseResponse
	CommandResponses[commands.ONLINE_COMMAND] = OnlineResponse
	CommandResponses[commands.TABLE_COMMAND] = TableRequest
	CommandResponses[commands.GENERATE_EVENTS_COMMAND] = GenerateEventsFormRequest
	CommandResponses[commands.TRACKLIST_COMMAND] = TracklistResponse
	CommandResponses[commands.TRACKFOLDER_COMMAND] = TrackFolderResponse
	CommandResponses[commands.RKG_COMMAND] = RKGResponse

	// Add interaction reponses here
	InteractionResps = append(InteractionResps, &InteractionResp{ID: PREVIOUS_BUTTON, Respond: IncPage, Permission: int64(discordgo.PermissionViewChannel)})
	InteractionResps = append(InteractionResps, &InteractionResp{ID: HOME_BUTTON, Respond: LeaderBoardResponse, Permission: int64(discordgo.PermissionViewChannel)})
	InteractionResps = append(InteractionResps, &InteractionResp{ID: NEXT_BUTTON, Respond: IncPage, Permission: int64(discordgo.PermissionViewChannel)})
	InteractionResps = append(InteractionResps, &InteractionResp{ID: TABLE_EDIT_BUTTON, Respond: EditTableRequest, Permission: int64(discordgo.PermissionManageMessages)})

	// Add modal responses here
	ModalResponses[TABLE_SUBMIT] = TableResponse
	ModalResponses[EDIT_TABLE_SUBMIT] = EditTableResponse
	ModalResponses[EVENT_SUBMIT] = GenerateEventsResponse

	// Add autocomplete responses here
	AutoCompleteResponses[commands.TRACK_OPTION_NAME] = TrackNameAutoComplete
}

func GetInteraction(ID string, responses []*InteractionResp) *InteractionResp {
	for _, response := range responses {
		if response.ID == ID {
			return response
		}
	}

	return nil
}
