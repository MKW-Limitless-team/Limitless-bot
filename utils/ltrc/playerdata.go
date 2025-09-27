package ltrc

import (
	"fmt"
	"sort"
)

type PlayerData struct {
	Name       string
	FriendCode string
	DiscordID  string
	Mmr        int64
	Mii        string
	Flag       string
}

// type Season struct {
// 	Name         string
// 	Active       bool
// 	Participants Participant
// }

// type Participant struct {
// 	Name string
// 	Mmr  int64
// }

func (playerData *PlayerData) ShowMii() string {
	return ShowMii(playerData.Mii)
}

func ShowMii(mii string) string {
	return fmt.Sprintf("https://mii-unsecure.ariankordi.net/miis/image.png?data=%s&expression=normal&cameraYRotate=30", mii)
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
