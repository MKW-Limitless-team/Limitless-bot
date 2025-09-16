package ltrc

import (
	"sort"
	"strconv"
	"strings"
)

type Placement struct {
	Track     string
	DiscordID string
	Flag      string
	Time      string
	Character string
	Vehicle   string
	DriftType string
	Category  string
	Accepted  bool
}

type byTime []*Placement

func (b byTime) Len() int {
	return len(b)
}

func (b byTime) Less(i int, j int) bool {
	time1 := b[i].Time
	time2 := b[j].Time

	time1Min, _ := strconv.Atoi(strings.Split(time1, ":")[0])
	time2Min, _ := strconv.Atoi(strings.Split(time2, ":")[0])

	if time1Min > time2Min {
		return time1Min > time2Min
	}

	time1Sec, _ := strconv.Atoi(strings.Split(strings.Split(time1, ":")[1], ".")[0])
	time2Sec, _ := strconv.Atoi(strings.Split(strings.Split(time1, ":")[1], ".")[0])

	if time1Sec > time2Sec {
		return time1Sec > time2Sec
	}

	time1Ms, _ := strconv.Atoi(strings.Split(strings.Split(time1, ":")[1], ".")[1])
	time2Ms, _ := strconv.Atoi(strings.Split(strings.Split(time1, ":")[1], ".")[1])

	if time1Ms > time2Ms {
		return time1Ms > time2Ms
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
