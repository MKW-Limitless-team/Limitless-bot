package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type RandomOption struct {
	Name   string
	Chance int
}

var (
	Modes     []*RandomOption = make([]*RandomOption, 0)
	Modifiers []*RandomOption = make([]*RandomOption, 0)
)

func SumTotal(options []*RandomOption) int {
	total := 0
	for _, option := range options {
		total += option.Chance
	}

	return total
}

func PickOne(source *rand.Rand, options []*RandomOption) *RandomOption {
	total := SumTotal(options)

	roll := source.Intn(total)
	fmt.Printf("Roll: %d Total: %d\n", roll, total)

	cumulative := 0
	for _, option := range options {
		cumulative += option.Chance
		if roll < cumulative {
			return option
		}
	}

	return nil
}

func PickMany(source *rand.Rand, options []*RandomOption, amount int) []*RandomOption {
	picked := make([]*RandomOption, 0)
	opts := make([]*RandomOption, 0)
	opts = append(opts, options...)

	for range amount {
		option := PickOne(source, opts)
		picked = append(picked, option)
		opts = deleteOption(opts, option)
	}

	return picked
}

func deleteOption(options []*RandomOption, option *RandomOption) []*RandomOption {
	opts := make([]*RandomOption, 0)
	for index, opt := range options {
		if opt.Name == option.Name {
			opts = append(options[:index], options[index+1:]...)
		}
	}

	return opts
}

func PopulateList(csvFile string, list []*RandomOption) []*RandomOption {
	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := strings.Split(scanner.Text(), ",")[0]
		change, err := strconv.Atoi(strings.Split(scanner.Text(), ",")[1])

		if err != nil {
			panic(err)
		}

		option := &RandomOption{Name: name, Chance: change}
		list = append(list, option)
	}

	return list
}
