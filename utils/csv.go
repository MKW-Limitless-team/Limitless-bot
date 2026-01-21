package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	Modes       []*RandomOption = make([]*RandomOption, 0)
	Modifiers   []*RandomOption = make([]*RandomOption, 0)
	Tracks      []*RandomOption = make([]*RandomOption, 0)
	FolderNames                 = make(map[string]string, 0)
)

func PopulateFolderNames(csvFile string, folderMap map[string]string) map[string]string {
	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := strings.Split(scanner.Text(), ",")[0]
		folder := strings.Split(scanner.Text(), ",")[1]

		folderMap[name] = folder
	}

	return folderMap
}

func PopulateRandomOptions(csvFile string, list []*RandomOption) []*RandomOption {
	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := strings.Split(scanner.Text(), ",")[0]
		chance, err := strconv.Atoi(strings.Split(scanner.Text(), ",")[1])

		if err != nil {
			panic(err)
		}

		option := &RandomOption{Name: name, Chance: chance}
		list = append(list, option)
	}

	return list
}
