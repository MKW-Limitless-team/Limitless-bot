package playerdata

type PlayerData struct {
	Name       string
	FriendCode string
	DiscordId  string
	Mmr        int
	Mii        string
}

type Season struct {
	Name         string
	Active       bool
	Participants Participant
}

type Participant struct {
	Name string
	Mmr  int
}
