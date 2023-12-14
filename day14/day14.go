package day14

import (
	"fmt"
	"github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"time"
)

func SolveDay() error {
	lines, err := utils.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day8 with readstringfromargs(): %s", err.Error())
		return err
	}

	startTime := time.Now()

	pattern := makePattern(lines)

	tiltNorth(&pattern)

	sum := calNorthBeam(&pattern)

	fmt.Printf("The total load is %d\n", sum)

	pattern = makePattern(lines)

	patternStore := make([][][]int32, 0)

	cycleStart := 0
	cycleLength := 0

	for i := 0; i < 1000000000; i++ {
		tiltNorth(&pattern)
		tiltWest(&pattern)
		tiltSouth(&pattern)
		tiltEast(&pattern)

		//fmt.Printf("After %d cycles:\n", i+1)
		//for y := range pattern {
		//	for x := range pattern[y] {
		//		fmt.Print(string(pattern[y][x]))
		//	}
		//	fmt.Print("\n")
		//}
		//fmt.Print("\n")

		foundPattern := false

		for index, oldPattern := range patternStore {
			equals := true
			for y := 0; y < len(pattern); y++ {
				for x := 0; x < len(pattern[y]); x++ {
					if pattern[y][x] != oldPattern[y][x] {
						equals = false
						break
					}
				}
				if !equals {
					break
				}
			}
			if equals {
				foundPattern = true
				cycleStart = index
				cycleLength = i - index
				break
			}
		}

		if foundPattern {
			break
		}

		patternCopy := make([][]int32, 0)

		for y := 0; y < len(pattern); y++ {
			rowCopy := make([]int32, len(pattern[y]))
			copy(rowCopy, pattern[y])
			patternCopy = append(patternCopy, rowCopy)
		}

		patternStore = append(patternStore, patternCopy)
	}

	loopCounter := (1000000000-cycleStart)%cycleLength - 1

	if loopCounter <= 0 {
		loopCounter = cycleLength - 1
	}

	for i := 0; i < loopCounter; i++ {
		tiltNorth(&pattern)
		tiltWest(&pattern)
		tiltSouth(&pattern)
		tiltEast(&pattern)
	}

	sum = calNorthBeam(&pattern)

	fmt.Printf("The total load is %d\n", sum)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Die Ausführungszeit betrug: %s\n", elapsedTime)

	return nil
}

func calNorthBeam(pattern *[][]int32) int {
	sum := 0
	for j := 0; j < len(*pattern); j++ {
		stoneCounter := 0
		for i := 0; i < len((*pattern)[j]); i++ {
			if (*pattern)[j][i] == 'O' {
				stoneCounter++
			}
		}
		sum += stoneCounter * (len(*pattern) - j)
	}
	return sum
}

func makePattern(lines []string) [][]int32 {
	pattern := make([][]int32, 0)
	for _, line := range lines {
		list := make([]int32, 0)
		for _, char := range line {
			list = append(list, char)
		}
		pattern = append(pattern, list)
	}
	return pattern
}

func tiltNorth(pattern *[][]int32) {
	for i := 0; i < len((*pattern)[0]); i++ {
		freeIndex := -1
		for j := 0; j < len(*pattern); j++ {
			if (*pattern)[j][i] == 'O' {
				if freeIndex > -1 {
					(*pattern)[j][i] = '.'
					(*pattern)[freeIndex][i] = 'O'
					freeIndex += 1
				}
			} else if (*pattern)[j][i] == '#' {
				freeIndex = -1
			} else if (*pattern)[j][i] == '.' {
				if freeIndex == -1 {
					freeIndex = j
				}
			}
		}
	}
}
func tiltWest(pattern *[][]int32) {
	for j := 0; j < len(*pattern); j++ {
		freeIndex := -1
		for i := 0; i < len((*pattern)[j]); i++ {
			if (*pattern)[j][i] == 'O' {
				if freeIndex > -1 {
					(*pattern)[j][i] = '.'
					(*pattern)[j][freeIndex] = 'O'
					freeIndex += 1
				}
			} else if (*pattern)[j][i] == '#' {
				freeIndex = -1
			} else if (*pattern)[j][i] == '.' {
				if freeIndex == -1 {
					freeIndex = i
				}
			}
		}
	}
}

func tiltSouth(pattern *[][]int32) {
	for i := 0; i < len((*pattern)[0]); i++ {
		freeIndex := -1
		for j := len(*pattern) - 1; j >= 0; j-- {
			if (*pattern)[j][i] == 'O' {
				if freeIndex > -1 {
					(*pattern)[j][i] = '.'
					(*pattern)[freeIndex][i] = 'O'
					freeIndex -= 1
				}
			} else if (*pattern)[j][i] == '#' {
				freeIndex = -1
			} else if (*pattern)[j][i] == '.' {
				if freeIndex == -1 {
					freeIndex = j
				}
			}
		}
	}
}

func tiltEast(pattern *[][]int32) {
	for j := 0; j < len(*pattern); j++ {
		freeIndex := -1
		for i := len((*pattern)[j]) - 1; i >= 0; i-- {
			if (*pattern)[j][i] == 'O' {
				if freeIndex > -1 {
					(*pattern)[j][i] = '.'
					(*pattern)[j][freeIndex] = 'O'
					freeIndex -= 1
				}
			} else if (*pattern)[j][i] == '#' {
				freeIndex = -1
			} else if (*pattern)[j][i] == '.' {
				if freeIndex == -1 {
					freeIndex = i
				}
			}
		}
	}
}
