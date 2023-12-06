package day4

import (
	"fmt"
	u "github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	Id             int
	WinningNumbers []int64
}

func (card *ScratchCard) play(numbers []int64) (bool, int, int) {
	var matches int = 0
	for _, winningNumber := range card.WinningNumbers {
		for _, number := range numbers {
			if winningNumber == number {
				matches += 1
				break
			}
		}
	}
	return matches > 0, matches, pow(2, matches-1)
}

func pow(n, m int) int {
	if m < 0 {
		return -1
	}
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func parseLine(input string) (*ScratchCard, []int64, error) {
	splitted := strings.Split(input, ":")
	cardID, err := strconv.Atoi(regexp.MustCompile(`(?m)\d+`).FindAllStringSubmatch(splitted[0], -1)[0][0])
	if err != nil {
		return nil, nil, err
	}
	x := strings.Split(splitted[1], "|")
	scratchNumbers, _, err := u.MakeIntSlice(x[0])
	if err != nil {
		return nil, nil, err
	}
	myNumbers, _, err := u.MakeIntSlice(x[1])
	if err != nil {
		return nil, nil, err
	}
	return &ScratchCard{Id: cardID, WinningNumbers: scratchNumbers}, myNumbers, nil
}

func SolveDayFour() error {
	lines, err := u.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day4 with readstringfromargs(): %s", err.Error())
		return err
	}

	solutionPart1, err := solvePart1(lines)
	if err != nil {
		fmt.Printf("error at solve day4 with solvePart1(): %s\n", err.Error())
		return err
	}
	fmt.Printf("Die Lösung für Tag 2 Part 1 ist %d!\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("error at solve day4 with solvePart2(): %s\n", err.Error())
		return err
	}

	fmt.Printf("Die Lösung für Tag 3 Part 2 ist %d!\n", solutionPart2)

	return nil
}

func solvePart1(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		card, numbers, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		hasMatch, _, points := card.play(numbers)
		if hasMatch {
			// fmt.Printf("Card %d has %d winning numbers, so it is worth %d points\n", card.Id, matches, points)
			sum += points
		} else {
			// fmt.Printf("Card %d has no winning numbers\n", card.Id)
		}
	}
	return sum, nil
}

func solvePart2(lines []string) (int, error) {
	var cards []int = make([]int, 0)
	var numbersMap map[int][]int64 = make(map[int][]int64)
	var cardMap map[int]*ScratchCard = make(map[int]*ScratchCard)
	for _, line := range lines {
		card, numbers, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		cards = append(cards, card.Id)
		cardMap[card.Id] = card
		numbersMap[card.Id] = numbers
	}

	for i := 0; i < len(cards); i++ {
		card := cardMap[cards[i]]
		hasMatch, matches, _ := card.play(numbersMap[card.Id])
		if hasMatch {
			for i := 1; i <= matches; i++ {
				cards = append(cards, card.Id+i)
			}
		}
	}

	return len(cards), nil
}
