package responses

import (
	"fmt"
	"limitless-bot/response"
	"limitless-bot/utils"
	"math/big"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	TRACKLIST_SUBMIT = "tracklist_submit"
)

func TracklistResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.
		NewMessageResponse().
		SetResponseData(TracklistData(session, interaction))

	return response.InteractionResponse
}

func TracklistData(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	data := response.NewResponseData("")

	amount := 32
	if len(interaction.ApplicationCommandData().Options) > 0 {
		amount = int(interaction.ApplicationCommandData().Options[0].IntValue())
	}

	seed, err := strconv.Atoi(interaction.ID)

	if err != nil {
		bytes := []byte(interaction.ID)
		seed = int(big.NewInt(0).SetBytes(bytes).Int64())
	}

	var msg strings.Builder
	msg.WriteString("# Tracklist:\n\n")
	source := rand.New(rand.NewSource(int64(seed)))

	tracks := utils.PickMany(source, utils.Tracks, amount)

	for _, track := range tracks {
		fmt.Fprintf(&msg, "%s\n", track.Name)
	}

	data.SetContent(msg.String())

	return data.InteractionResponseData
}
