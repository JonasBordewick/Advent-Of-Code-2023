package day3

import (
	"fmt"
	u "github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"math"
	"regexp"
	"strconv"
	"unicode"
)

func SolveDayThree() error {
	lines, err := u.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day3 with readstringfromargs(): %s", err.Error())
		return err
	}

	solutionPart1, err := solvePart1(lines)
	if err != nil {
		fmt.Printf("error at solve day3 with solvePart1(): %s\n", err.Error())
		return err
	}
	fmt.Printf("Die Lösung für Tag 2 Part 1 ist %d!\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("error at solve day3 with solvePart2(): %s\n", err.Error())
		return err
	}

	fmt.Printf("Die Lösung für Tag 3 Part 2 ist %d!\n", solutionPart2)

	return nil
}

func solvePart1(lines []string) (int, error) {
	var sum = 0
	for lineIndex, line := range lines {
		var re = regexp.MustCompile(`(?m)\d+`)
		matches := re.FindAllStringSubmatch(line, -1)
		indices := re.FindAllStringIndex(line, -1)
		for i := range matches {
			start := indices[i][0]
			end := indices[i][1]
			adjacentToASymbol := false

			if start > 0 {
				start -= 1
			}

			if end == len(line) {
				end -= 1
			}

			if isSpecialCharacter(rune(line[start])) || isSpecialCharacter(rune(line[end])) {
				adjacentToASymbol = true
			}

			if lineIndex > 0 {
				lineAbove := lines[lineIndex-1]

				for index := start; index <= end; index++ {
					if isSpecialCharacter(rune(lineAbove[index])) {
						adjacentToASymbol = true
					}
				}

			}

			if lineIndex < len(lines)-1 {
				lineAbove := lines[lineIndex+1]

				for index := start; index <= end; index++ {
					if isSpecialCharacter(rune(lineAbove[index])) {
						adjacentToASymbol = true
					}
				}
			}

			if adjacentToASymbol {
				number, err := strconv.Atoi(matches[i][0])
				if err != nil {
					return -1, err
				}
				sum += number
			}

		}
	}

	return sum, nil
}

func isSpecialCharacter(char rune) bool {
	return !unicode.IsDigit(char) && char != '.'
}

func solvePart2(lines []string) (int, error) {

	var sum = 0

	for lineIndex, line := range lines {
		var gr = regexp.MustCompile(`(?m)\*`)
		var nr = regexp.MustCompile(`(?m)\d+`)

		gears := gr.FindAllStringSubmatch(line, -1)
		indices := gr.FindAllStringIndex(line, -1)
		for i := range gears {
			gearIndex := indices[i][0]

			var gearRatioA int = -1
			var gearRatioB int = -1

			numbersInLine := nr.FindAllStringSubmatch(line, -1)
			numberIndices := nr.FindAllStringIndex(line, -1)
			for j := range numbersInLine {
				start := numberIndices[j][0]
				end := numberIndices[j][1] - 1
				if isAdjacent(start, end, gearIndex) {
					number, err := strconv.Atoi(numbersInLine[j][0])
					if err != nil {
						return -1, err
					}

					if gearRatioA == -1 {
						gearRatioA = number
					} else if gearRatioB == -1 {
						gearRatioB = number
					}
				}

			}

			if lineIndex > 0 {
				lineAbove := lines[lineIndex-1]
				numbersInLine = nr.FindAllStringSubmatch(lineAbove, -1)
				numberIndices = nr.FindAllStringIndex(lineAbove, -1)
				for j := range numbersInLine {
					start := numberIndices[j][0]
					end := numberIndices[j][1] - 1
					if isAdjacent(start, end, gearIndex) {
						number, err := strconv.Atoi(numbersInLine[j][0])
						if err != nil {
							return -1, err
						}

						if gearRatioA == -1 {
							gearRatioA = number
						} else if gearRatioB == -1 {
							gearRatioB = number
						}
					}

				}
			}

			if lineIndex < len(lines)-1 {
				lineBelow := lines[lineIndex+1]
				numbersInLine = nr.FindAllStringSubmatch(lineBelow, -1)
				numberIndices = nr.FindAllStringIndex(lineBelow, -1)
				for j := range numbersInLine {
					start := numberIndices[j][0]
					end := numberIndices[j][1] - 1
					if isAdjacent(start, end, gearIndex) {
						number, err := strconv.Atoi(numbersInLine[j][0])
						if err != nil {
							return -1, err
						}

						if gearRatioA == -1 {
							gearRatioA = number
						} else if gearRatioB == -1 {
							gearRatioB = number
						}
					}
				}
			}
			if gearRatioA > -1 && gearRatioB > -1 {
				sum += gearRatioA * gearRatioB
			}
		}
	}

	return sum, nil
}

func isAdjacent(start, end, to int) bool {
	for index := start; index <= end; index++ {
		if math.Abs(float64(index-to)) <= 1 {
			return true
		}
	}
	return false
}
