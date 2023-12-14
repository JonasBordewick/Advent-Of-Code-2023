package day12

import (
	"fmt"
	"github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"strings"
)

func calc(pattern string, groupOfDamagedSprings []int64) int {
	if len(pattern) == 0 {
		if len(groupOfDamagedSprings) == 0 {
			return 0
		}
		return 1
	}

	if len(groupOfDamagedSprings) == 0 {
		notInPattern := true
		for i := 0; i < len(pattern); i++ {
			if pattern[i] == '#' {
				notInPattern = false
			}
		}
		if notInPattern {
			return 1
		}
		return 0
	}

	result := 0

	if pattern[0] == '.' || pattern[0] == '?' {
		result += calc(pattern[1:], groupOfDamagedSprings)
	}

	if pattern[0] == '#' || pattern[0] == '?' {
		if groupOfDamagedSprings[0] <= int64(len(pattern)) {
			notInPattern := true
			for i := 0; i < len(pattern[:groupOfDamagedSprings[0]]); i++ {
				if pattern[i] == '.' {
					notInPattern = false
				}
			}
			if notInPattern {
				if groupOfDamagedSprings[0] == int64(len(pattern)) || pattern[groupOfDamagedSprings[0]] != '#' {
					result += calc(pattern[groupOfDamagedSprings[0]:], groupOfDamagedSprings[1:])
				}
			}
		}
	}

	return result
}

func SolveDay() error {
	lines, err := utils.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day8 with readstringfromargs(): %s", err.Error())
		return err
	}

	total := 0

	for _, line := range lines {
		splitted := strings.Split(line, " ")
		pattern := splitted[0]
		groupsOfDamagedSprings, _, err := utils.MakeIntSlice(splitted[1])
		if err != nil {
			return err
		}
		total += calc(pattern, groupsOfDamagedSprings)
	}

	fmt.Println(total)
	return nil
}
