package day1

import (
	"fmt"
	u "github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"regexp"
	"strconv"
	"strings"
)

func SolveDayOne() error {
	lines, err := u.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day1 with readstringfromargs(): %s", err.Error())
		return err
	}

	solutionPart1, err := solvePart1(lines)
	if err != nil {
		fmt.Printf("error at solve day1 with solvePart1(): %s\n", err.Error())
		return err
	}
	fmt.Printf("Die Lösung für Tag 1 Part 1 ist %d!\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("error at solve day1 with solvePart2(): %s\n", err.Error())
		return err
	}

	fmt.Printf("Die Lösung für Tag 1 Part 2 ist %d!\n", solutionPart2)

	return nil
}

func solvePart1(lines []string) (int, error) {
	var sum = 0
	var re = regexp.MustCompile(`\d`)
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		calibrationValue, err := strconv.Atoi(fmt.Sprintf("%c%c", matches[0][0][0], matches[len(matches)-1][0][0]))
		if err != nil {
			fmt.Printf("error at solve day1 with strconv.Atoi(): %s", err.Error())
			return -1, err
		}
		sum += calibrationValue
	}
	return sum, nil
}

func solvePart2(lines []string) (int, error) {
	var re = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	var sum = 0
	for _, line := range lines {
		line = strings.ReplaceAll(line, "one", "onee")
		line = strings.ReplaceAll(line, "two", "twoo")
		line = strings.ReplaceAll(line, "three", "three")
		line = strings.ReplaceAll(line, "five", "fivee")
		line = strings.ReplaceAll(line, "seven", "sevenn")
		line = strings.ReplaceAll(line, "eight", "eightt")
		line = strings.ReplaceAll(line, "nine", "ninee")
		matches := re.FindAllStringSubmatch(line, -1)
		calibrationValue, err := strconv.Atoi(
			fmt.Sprintf("%c%c", mapToRune(matches[0][0]), mapToRune(matches[len(matches)-1][0])))
		if err != nil {
			fmt.Printf("error at solve day1 with strconv.Atoi(): %s", err.Error())
			return -1, err
		}
		sum += calibrationValue
	}
	return sum, nil
}

func mapToRune(match string) rune {
	switch match {
	case "one":
		return '1'
	case "two":
		return '2'
	case "three":
		return '3'
	case "four":
		return '4'
	case "five":
		return '5'
	case "six":
		return '6'
	case "seven":
		return '7'
	case "eight":
		return '8'
	case "nine":
		return '9'
	default:
		return rune(match[0])
	}
}
