package ltrc

import (
	"sort"
)

type Placement struct {
	ID           int
	Track        string
	DiscordID    string
	Minutes      int
	Seconds      int
	Milliseconds int
	Character    string
	Vehicle      string
	DriftType    string
	Category     string
	Url          string
	CRC          uint32
	Approved     bool
}

type byTime []*Placement

func (b byTime) Len() int {
	return len(b)
}

func (b byTime) Less(i int, j int) bool {
	time1 := b[i]
	time2 := b[j]

	if time1.Minutes > time2.Minutes {
		return true
	} else if time1.Seconds > time2.Seconds {
		return true
	} else if time1.Milliseconds > time2.Milliseconds {
		return true
	}

	return false
}

func (b byTime) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func SortByTime(placements []*Placement) []*Placement {
	sort.Sort(byTime(placements))

	return placements
}
