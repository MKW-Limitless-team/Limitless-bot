package responses

import (
	"bytes"
	"encoding/json"
	"fmt"
	"limitless-bot/commands"
	"limitless-bot/components"
	"limitless-bot/components/button"
	e "limitless-bot/components/embed"
	r "limitless-bot/response"
	"limitless-bot/utils"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	ltrcResponses "github.com/MKW-Limitless-team/limitless-types/responses"
	"github.com/bwmarrin/discordgo"
)

var (
	USAGE_BUTTON = "usage:"
)

var discordMentionRegex = regexp.MustCompile(`^<@!?([0-9]+)>$`)

type UsageEntry struct {
	Name  string
	Count int
	Rank  int
}

type UsageSpec struct {
	StatsType string
	Title     string
	Endpoint  string
}

type PInfoRequest struct {
	PID uint64 `json:"pid"`
}

type PInfoResponse struct {
	Player  *PInfoPlayer `json:"player"`
	Success bool         `json:"success"`
	Error   string       `json:"error"`
}

type PInfoPlayer struct {
	ProfileID uint64 `json:"profile_id"`
	MiiName   string `json:"mii_name"`
	DiscordID string `json:"discord_id"`
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
	messageComponent := interaction.MessageComponentData()
	parts := strings.Split(messageComponent.CustomID, ":")
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

	sortMode := parts[3]
	page := UsagePage(parts[4])
	if len(messageComponent.Values) > 0 {
		sortMode = messageComponent.Values[0]
		page = 1
	}

	response := r.NewMessageResponse().
		SetResponseData(UsageEmbedData(session, interaction.GuildID, spec, pid, sortMode, page))
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
	pid, errMessage := GetUsageRequest(session, interaction)
	if errMessage != "" {
		return r.NewResponseData(errMessage).InteractionResponseData
	}

	return UsageEmbedData(session, interaction.GuildID, spec, pid, "uses_desc", page)
}

func UsageEmbedData(session *discordgo.Session, guildID string, spec *UsageSpec, pid uint64, sortMode string, page int) *discordgo.InteractionResponseData {
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

	displayName := GetUsageDisplayName(session, guildID, pid)
	embed := e.NewRichEmbed(fmt.Sprintf("**%s's %s usage**", displayName, strings.TrimSuffix(spec.StatsType, "s")), "", 0xd70ccf)
	embed.SetFooter(fmt.Sprintf("Page : %d / %d | Sort: %s", page, pageCount, UsageSortLabel(sortMode)), "")

	start := (page - 1) * entriesPerPage
	end := start + entriesPerPage
	if end > len(entries) {
		end = len(entries)
	}

	for _, entry := range entries[start:end] {
		embed.Description += fmt.Sprintf("%d. **%s** — *%d %s*\n", entry.Rank, entry.Name, entry.Count, UseText(entry.Count))
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

	actionRow := components.NewActionRow()
	actionRow.AddComponent(UsageSortMenu(spec.StatsType, pid, sortMode))
	data.AddComponent(actionRow)

	return data.InteractionResponseData
}

func GetUsageRequest(session *discordgo.Session, interaction *discordgo.InteractionCreate) (uint64, string) {
	options := interaction.ApplicationCommandData().Options
	userOption := utils.GetOption(options, "user")

	userID := interaction.Member.User.ID
	if userOption != nil {
		user := strings.TrimSpace(userOption.StringValue())
		mention := discordMentionRegex.FindStringSubmatch(user)
		if len(mention) == 2 {
			userID = mention[1]
		} else {
			pid, err := GetUsagePID(user)
			if err != nil {
				return 0, "Invalid user or PID"
			}

			return pid, ""
		}
	}

	pid, err := GetRegisteredPID(userID)
	if err != nil {
		return 0, "This user is not registered yet"
	}

	return pid, ""
}

func GetUsagePID(user string) (uint64, error) {
	pid, err := strconv.ParseUint(user, 10, 64)
	if err != nil || pid == 0 {
		return 0, fmt.Errorf("invalid pid")
	}

	return pid, nil
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

func GetUsageDisplayName(session *discordgo.Session, guildID string, pid uint64) string {
	pinfoRequest := &PInfoRequest{PID: pid}
	marshalled, err := json.Marshal(pinfoRequest)
	if err != nil {
		return fmt.Sprintf("%d", pid)
	}

	resp, err := http.Post("http://wfc.blazico.nl/api/pinfo", "application/json", bytes.NewBuffer(marshalled))
	if err != nil {
		return fmt.Sprintf("%d", pid)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("%d", pid)
	}

	var pinfoResponse *PInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&pinfoResponse)
	if err != nil || pinfoResponse == nil || !pinfoResponse.Success || pinfoResponse.Player == nil {
		return fmt.Sprintf("%d", pid)
	}

	if pinfoResponse.Player.DiscordID != "" {
		return GetDiscordDisplayName(session, guildID, pinfoResponse.Player.DiscordID)
	}

	if pinfoResponse.Player.MiiName != "" {
		return pinfoResponse.Player.MiiName
	}

	return fmt.Sprintf("%d", pid)
}

func GetDiscordDisplayName(session *discordgo.Session, guildID string, userID string) string {
	member, err := session.GuildMember(guildID, userID)
	if err == nil && member != nil {
		return member.DisplayName()
	}

	user, err := session.User(userID)
	if err == nil && user != nil {
		return user.Username
	}

	return fmt.Sprintf("<@%s>", userID)
}

func UseText(count int) string {
	if count == 1 {
		return "use"
	}

	return "uses"
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
		if key == "profile_id" {
			continue
		}

		entries = append(entries, &UsageEntry{
			Name:  UsageDisplayName(spec.StatsType, key),
			Count: count,
		})
	}

	SetUsageRanks(entries)
	SortUsageEntries(entries, sortMode)

	return entries, nil
}

func SetUsageRanks(entries []*UsageEntry) {
	sorted := make([]*UsageEntry, len(entries))
	copy(sorted, entries)

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Count == sorted[j].Count {
			return sorted[i].Name < sorted[j].Name
		}

		return sorted[i].Count > sorted[j].Count
	})

	rank := 0
	lastCount := -1
	for i, entry := range sorted {
		if entry.Count != lastCount {
			rank = i + 1
			lastCount = entry.Count
		}

		entry.Rank = rank
	}
}

func SortUsageEntries(entries []*UsageEntry, sortMode string) {
	sort.Slice(entries, func(i, j int) bool {
		if sortMode == "alphabetical" {
			return entries[i].Name < entries[j].Name
		}

		if sortMode == "uses_asc" {
			if entries[i].Count == entries[j].Count {
				return entries[i].Name < entries[j].Name
			}

			return entries[i].Count < entries[j].Count
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

func UsageSortMenu(statsType string, pid uint64, sortMode string) discordgo.MessageComponent {
	minValues := 1
	options := []discordgo.SelectMenuOption{
		UsageSortOption("Uses descending", "uses_desc", sortMode),
		UsageSortOption("Uses ascending", "uses_asc", sortMode),
		UsageSortOption("Alphabetical", "alphabetical", sortMode),
	}

	return discordgo.SelectMenu{
		MenuType:    discordgo.StringSelectMenu,
		CustomID:    UsageButtonID(statsType, pid, sortMode, 1),
		Placeholder: "Sort usage",
		MinValues:   &minValues,
		MaxValues:   1,
		Options:     options,
	}
}

func UsageSortOption(label string, value string, sortMode string) discordgo.SelectMenuOption {
	return discordgo.SelectMenuOption{
		Label:   label,
		Value:   value,
		Default: sortMode == value,
	}
}

func UsageSortLabel(sortMode string) string {
	switch sortMode {
	case "alphabetical":
		return "alphabetical"
	case "uses_asc":
		return "uses ascending"
	default:
		return "uses descending"
	}
}

func UsagePage(pageStr string) int {
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 1
	}

	return page
}
