package utils

import "sort"

type PlayerData struct {
	Name       string
	FriendCode string
	DiscordID  string
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

type byMmr []*PlayerData

// Len implements sort.Interface.
func (b byMmr) Len() int {
	return len(b)
}

// Less implements sort.Interface.
func (b byMmr) Less(i int, j int) bool {
	return b[i].Mmr > b[j].Mmr
}

// Swap implements sort.Interface.
func (b byMmr) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func SortByMMR(players []*PlayerData) []*PlayerData {
	sort.Sort(byMmr(players))

	return players
}
