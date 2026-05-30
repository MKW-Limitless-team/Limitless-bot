package responses

import (
	"encoding/json"
	"fmt"
	"limitless-bot/commands"
	"limitless-bot/components"
	"limitless-bot/components/button"
	e "limitless-bot/components/embed"
	r "limitless-bot/response"
	"limitless-bot/utils"
	"net/http"
	"sort"
	"strconv"
	"strings"

	ltrcResponses "github.com/MKW-Limitless-team/limitless-types/responses"
	"github.com/bwmarrin/discordgo"
)

var (
	USAGE_BUTTON = "usage:"
)

type UsageEntry struct {
	Name  string
	Count int
}

type UsageSpec struct {
	StatsType string
	Title     string
	Endpoint  string
}

var characterNames = map[string]string{
	"baby_daisy":     "Baby Daisy",
	"baby_luigi":     "Baby Luigi",
	"baby_mario":     "Baby Mario",
	"baby_peach":     "Baby Peach",
	"birdo":          "Birdo",
	"bowser":         "Bowser",
	"bowser_jr":      "Bowser Jr.",
	"daisy":          "Daisy",
	"diddy_kong":     "Diddy Kong",
	"donkey_kong":    "Donkey Kong",
	"dry_bones":      "Dry Bones",
	"dry_bowser":     "Dry Bowser",
	"funky_kong":     "Funky Kong",
	"king_boo":       "King Boo",
	"koopa_troopa":   "Koopa Troopa",
	"luigi":          "Luigi",
	"mario":          "Mario",
	"mii_l_a_female": "Large Mii Outfit A (Female)",
	"mii_l_a_male":   "Large Mii Outfit A (Male)",
	"mii_l_b_female": "Large Mii Outfit B (Female)",
	"mii_l_b_male":   "Large Mii Outfit B (Male)",
	"mii_l_c_female": "Large Mii Outfit C (Female)",
	"mii_l_c_male":   "Large Mii Outfit C (Male)",
	"mii_large":      "Large Mii",
	"mii_m_a_female": "Medium Mii Outfit A (Female)",
	"mii_m_a_male":   "Medium Mii Outfit A (Male)",
	"mii_m_b_female": "Medium Mii Outfit B (Female)",
	"mii_m_b_male":   "Medium Mii Outfit B (Male)",
	"mii_m_c_female": "Medium Mii Outfit C (Female)",
	"mii_m_c_male":   "Medium Mii Outfit C (Male)",
	"mii_medium":     "Medium Mii",
	"mii_s_a_female": "Small Mii Outfit A (Female)",
	"mii_s_a_male":   "Small Mii Outfit A (Male)",
	"mii_s_b_female": "Small Mii Outfit B (Female)",
	"mii_s_b_male":   "Small Mii Outfit B (Male)",
	"mii_s_c_female": "Small Mii Outfit C (Female)",
	"mii_s_c_male":   "Small Mii Outfit C (Male)",
	"mii_small":      "Small Mii",
	"peach":          "Peach",
	"rosalina":       "Rosalina",
	"toad":           "Toad",
	"toadette":       "Toadette",
	"waluigi":        "Waluigi",
	"wario":          "Wario",
	"yoshi":          "Yoshi",
}

var vehicleNames = map[string]string{
	"bit_bike":         "Bit Bike",
	"blue_falcon":      "Blue Falcon",
	"booster_seat":     "Booster Seat",
	"bullet_bike":      "Bullet Bike",
	"cheep_charger":    "Cheep Charger",
	"classic_dragster": "Classic Dragster",
	"daytripper":       "Daytripper",
	"dolphin_dasher":   "Dolphin Dasher",
	"flame_flyer":      "Flame Flyer",
	"flame_runner":     "Flame Runner",
	"honeycoupe":       "Honeycoupe",
	"jet_bubble":       "Jet Bubble",
	"jetsetter":        "Jetsetter",
	"mach_bike":        "Mach Bike",
	"magikruiser":      "Magikruiser",
	"mini_beast":       "Mini Beast",
	"offroader":        "Offroader",
	"phantom":          "Phantom",
	"piranha_prowler":  "Piranha Prowler",
	"quacker":          "Quacker",
	"shooting_star":    "Shooting Star",
	"sneakster":        "Sneakster",
	"spear":            "Spear",
	"sprinter":         "Sprinter",
	"standard_bike_l":  "Standard Bike L",
	"standard_bike_m":  "Standard Bike M",
	"standard_bike_s":  "Standard Bike S",
	"standard_kart_l":  "Standard Kart L",
	"standard_kart_m":  "Standard Kart M",
	"standard_kart_s":  "Standard Kart S",
	"sugarscoot":       "Sugarscoot",
	"super_blooper":    "Super Blooper",
	"tiny_titan":       "Tiny Titan",
	"wario_bike":       "Wario Bike",
	"wild_wing":        "Wild Wing",
	"zip_zip":          "Zip Zip",
}

func CharactersResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := r.NewMessageResponse().
		SetResponseData(UsageData(session, interaction, CharacterUsageSpec(), 1))

	return response.InteractionResponse
}

func VehiclesResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := r.NewMessageResponse().
		SetResponseData(UsageData(session, interaction, VehicleUsageSpec(), 1))

	return response.InteractionResponse
}

func UsagePageResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	parts := strings.Split(interaction.MessageComponentData().CustomID, ":")
	if len(parts) != 5 {
		return r.NewMessageResponse().SetResponseData(r.NewResponseData("Unable to change page").InteractionResponseData).InteractionResponse
	}

	spec := CharacterUsageSpec()
	if parts[1] == "vehicles" {
		spec = VehicleUsageSpec()
	}

	pid, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return r.NewMessageResponse().SetResponseData(r.NewResponseData("Invalid PID").InteractionResponseData).InteractionResponse
	}

	page, err := strconv.Atoi(parts[4])
	if err != nil {
		page = 1
	}

	response := r.NewMessageResponse().
		SetResponseData(UsageEmbedData(spec, pid, parts[3], page))
	response.Type = discordgo.InteractionResponseUpdateMessage

	return response.InteractionResponse
}

func CharacterUsageSpec() *UsageSpec {
	return &UsageSpec{
		StatsType: "characters",
		Title:     "**Character Usage**",
		Endpoint:  "http://wfc.blazico.nl/api/mkw_characters?profile_id=%d",
	}
}

func VehicleUsageSpec() *UsageSpec {
	return &UsageSpec{
		StatsType: "vehicles",
		Title:     "**Vehicle Usage**",
		Endpoint:  "http://wfc.blazico.nl/api/mkw_vehicles?profile_id=%d",
	}
}

func UsageData(session *discordgo.Session, interaction *discordgo.InteractionCreate, spec *UsageSpec, page int) *discordgo.InteractionResponseData {
	pid, sortMode, errMessage := GetUsageRequest(session, interaction)
	if errMessage != "" {
		return r.NewResponseData(errMessage).InteractionResponseData
	}

	return UsageEmbedData(spec, pid, sortMode, page)
}

func UsageEmbedData(spec *UsageSpec, pid uint64, sortMode string, page int) *discordgo.InteractionResponseData {
	data := r.NewResponseData("")
	entries, err := GetUsageEntries(spec, pid, sortMode)
	if err != nil || len(entries) == 0 {
		return r.NewResponseData("This user is not registered yet").InteractionResponseData
	}

	entriesPerPage := 10
	pageCount := (len(entries) + entriesPerPage - 1) / entriesPerPage
	if page < 1 {
		page = 1
	}
	if page > pageCount {
		page = pageCount
	}

	embed := e.NewRichEmbed(spec.Title, fmt.Sprintf("PID: `%d`", pid), 0xd70ccf)
	embed.SetFooter(fmt.Sprintf("Page : %d / %d | Sort: %s", page, pageCount, sortMode), "")

	start := (page - 1) * entriesPerPage
	end := start + entriesPerPage
	if end > len(entries) {
		end = len(entries)
	}

	for _, entry := range entries[start:end] {
		embed.AddField("", fmt.Sprintf("**%s:** %d", entry.Name, entry.Count), false)
	}

	data.AddEmbed(embed)

	if pageCount > 1 {
		actionRow := components.NewActionRow()
		previousButton := button.NewBasicButton("Previous", UsageButtonID(spec.StatsType, pid, sortMode, page-1), discordgo.PrimaryButton, page == 1)
		nextButton := button.NewBasicButton("Next", UsageButtonID(spec.StatsType, pid, sortMode, page+1), discordgo.PrimaryButton, page == pageCount)

		actionRow.AddComponent(previousButton)
		actionRow.AddComponent(nextButton)
		data.AddComponent(actionRow)
	}

	return data.InteractionResponseData
}

func GetUsageRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) (uint64, string, string) {
	options := interaction.ApplicationCommandData().Options
	userOption := utils.GetOption(options, "user")
	pidOption := utils.GetOption(options, "pid")
	sortOption := utils.GetOption(options, "sort")

	if userOption != nil && pidOption != nil {
		return 0, "", "Please provide either a user or a PID, not both"
	}

	sortMode := "number"
	if sortOption != nil {
		sortMode = sortOption.StringValue()
	}

	if sortMode != "alphabetical" {
		sortMode = "number"
	}

	if pidOption != nil {
		pid, err := strconv.ParseUint(pidOption.StringValue(), 10, 64)
		if err != nil || pid == 0 {
			return 0, "", "Invalid PID"
		}

		return pid, sortMode, ""
	}

	userID := interaction.Member.User.ID
	if userOption != nil {
		userID = userOption.UserValue(session).ID
	}

	pid, err := GetRegisteredPID(userID)
	if err != nil {
		return 0, "", "This user is not registered yet"
	}

	return pid, sortMode, ""
}

func GetRegisteredPID(userID string) (uint64, error) {
	resp, err := http.Get("http://localhost:5000/player?discord_id=" + userID)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("player request failed")
	}

	var jsonResponse *ltrcResponses.PlayerInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		return 0, err
	}

	if jsonResponse == nil || jsonResponse.Status == ltrcResponses.Failure || jsonResponse.PlayerData == nil {
		return 0, fmt.Errorf("player not registered")
	}

	return jsonResponse.PlayerData.ProfileID, nil
}

func GetUsageEntries(spec *UsageSpec, pid uint64, sortMode string) ([]*UsageEntry, error) {
	resp, err := http.Get(fmt.Sprintf(spec.Endpoint, pid))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("usage request failed")
	}

	stats := map[string]int{}
	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		return nil, err
	}

	entries := make([]*UsageEntry, 0)
	for key, count := range stats {
		if key == "profile_id" || count == 0 {
			continue
		}

		entries = append(entries, &UsageEntry{
			Name:  UsageDisplayName(spec.StatsType, key),
			Count: count,
		})
	}

	SortUsageEntries(entries, sortMode)

	return entries, nil
}

func SortUsageEntries(entries []*UsageEntry, sortMode string) {
	sort.Slice(entries, func(i, j int) bool {
		if sortMode == "alphabetical" {
			return entries[i].Name < entries[j].Name
		}

		if entries[i].Count == entries[j].Count {
			return entries[i].Name < entries[j].Name
		}

		return entries[i].Count > entries[j].Count
	})
}

func UsageDisplayName(statsType string, key string) string {
	if statsType == commands.CHARACTERS_COMMAND {
		name, ok := characterNames[key]
		if ok {
			return name
		}
	}

	if statsType == commands.VEHICLES_COMMAND {
		name, ok := vehicleNames[key]
		if ok {
			return name
		}
	}

	parts := strings.Split(key, "_")
	for i, part := range parts {
		if part == "" {
			continue
		}

		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}

	return strings.Join(parts, " ")
}

func UsageButtonID(statsType string, pid uint64, sortMode string, page int) string {
	return fmt.Sprintf("%s%s:%d:%s:%d", USAGE_BUTTON, statsType, pid, sortMode, page)
}
