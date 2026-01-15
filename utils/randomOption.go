package utils

import (
	"bufio"
	"maps"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	Modes     map[string]int = make(map[string]int)
	Modifiers map[string]int = make(map[string]int)
)

func SumTotal(options map[string]int) int {
	total := 0
	for _, chance := range options {
		total += chance
	}

	return total
}

func PickOne(source *rand.Rand, options map[string]int) string {
	total := SumTotal(options)

	roll := source.Intn(total)
	println(roll)

	cumulative := 0
	for option, chance := range options {
		cumulative += chance
		if roll < cumulative {
			return option
		}
	}

	return ""
}

func PickMany(source *rand.Rand, options map[string]int, amount int) []string {
	picked := make([]string, 0)
	opts := maps.Clone(options)

	for range amount {
		option := PickOne(source, opts)
		picked = append(picked, option)
		delete(opts, option)
	}

	return picked
}

func PopulateMap(csvFile string, dstMap map[string]int) {
	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key := strings.Split(scanner.Text(), ",")[0]
		value, err := strconv.Atoi(strings.Split(scanner.Text(), ",")[1])

		if err != nil {
			panic(err)
		}

		dstMap[key] = value
	}
}
