package responses

import (
	"fmt"
	"limitless-bot/components"
	"limitless-bot/components/modal"
	"limitless-bot/response"
	"limitless-bot/utils"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	EVENT_SUBMIT = "event_submit"
)

func GenerateEventsFormRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewModalResponse().
		SetResponseData(GenerateEventsForm())

	return response.InteractionResponse
}

func GenerateEventsForm() *discordgo.InteractionResponseData {
	data := response.NewFormData("Event", EVENT_SUBMIT)

	actionRow := components.NewActionRow()
	date := modal.NewTextField("Date (01-01-2026)", "date", "dd-mm-yyyy", true)
	actionRow.AddComponent(date)
	data.AddComponent(actionRow)

	actionRow = components.NewActionRow()
	seed := modal.NewTextField("Seed (Optional)", "seed", "Seed (Optional)", false)
	actionRow.AddComponent(seed)
	data.AddComponent(actionRow)

	return data.InteractionResponseData
}

func GenerateEventsResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := response.NewMessageResponse().
		SetResponseData(GenerateEventsData(interaction))

	return response.InteractionResponse
}

func GenerateEventsData(interaction *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	data := response.NewResponseData("")
	submitData := interaction.ModalSubmitData()

	seedStr, _ := utils.GetSubmitDataValueByID(submitData, "seed")
	seed := getSeed(seedStr, interaction.ID)

	dateStr, _ := utils.GetSubmitDataValueByID(submitData, "date")
	date, err := utils.GetTime(dateStr)

	if err != nil {
		data.SetContent("Incorrect date format. E.g `01-01-2026`")
		return data.InteractionResponseData
	}

	dates := make([]time.Time, 0)
	dates = append(dates, date)
	dates = append(dates, date.Add(time.Hour*24))
	dates = append(dates, date.Add(time.Hour*48))

	labels := []string{":regional_indicator_a:", ":regional_indicator_b:", ":regional_indicator_c:", ":regional_indicator_d:", ":regional_indicator_e:", ":regional_indicator_f:"}

	var msg strings.Builder
	msg.WriteString("# All events require a minimum 8 players, except 3v3s which require 9\n")
	source := rand.New(rand.NewSource(seed))

	for i, date := range dates {
		fmt.Fprintf(&msg, "## %s of %s\n", utils.DayToString(date.Day()), date.Month().String())

		events := utils.PickMany(source, utils.Modes, 2)

		for j, event := range events {
			fmt.Fprintf(&msg, "### Event %s | %s | \n", labels[(i*2)+j], event.Name)

			fmt.Fprintf(&msg, "Starting Time: Between %s and %s\n",
				utils.CreateTimeStamp(date), utils.CreateTimeStamp(date.Add(time.Hour*1+time.Minute*14)))

			date = date.Add(time.Hour*1 + time.Minute*15)
		}

	}

	fmt.Fprintf(&msg, "`Seed: %d`", seed)

	data.SetContent(msg.String())

	return data.InteractionResponseData
}

func getSeed(seedStr string, fallback string) int64 {
	var seed int
	var err error

	if seedStr == "" {
		seed, err = strconv.Atoi(fallback)
		if err != nil {
			bytes := []byte(fallback)
			seed = int(big.NewInt(0).SetBytes(bytes).Int64())
		}
	} else {
		seed, err = strconv.Atoi(seedStr)
		if err != nil {
			bytes := []byte(seedStr)
			seed = int(big.NewInt(0).SetBytes(bytes).Int64())
		}
	}

	return int64(seed)
}
