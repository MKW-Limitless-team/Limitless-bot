package playerdata

type PlayerData struct {
	Name       string
	FriendCode string
	DiscordId  string
	Mmr        float64
	Mii        string
}

type Season struct {
	Name         string
	Active       bool
	Participants Participant
}

type Participant struct {
	Name string
	Mmr  float64
}
