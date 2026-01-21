package utils

import (
	"fmt"
	"math/rand"
)

type RandomOption struct {
	Name   string
	Chance int
}

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
