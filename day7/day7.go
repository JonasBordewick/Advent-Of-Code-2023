package day7

import (
	"fmt"
	"github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"sort"
	"strconv"
	"strings"
)

type HandRanking int

const (
	FiveOfAKind HandRanking = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

type Hand struct {
	Cards []rune
	Bid   int
	Rank  HandRanking
}

func (hand *Hand) Evaluate(part2 bool) HandRanking {
	var cardMap map[rune]int = make(map[rune]int)

	for _, card := range hand.Cards {
		cardMap[card] += 1
	}

	if part2 && len(cardMap) > 1 {

		jokers, hasJokers := cardMap['1']
		if hasJokers {

			maxKey := '1'
			maxValue := 0

			for key, value := range cardMap {
				if key == '1' {
					continue
				}
				if value > maxValue {
					maxKey = key
					maxValue = value
				} else if value == maxValue && key > maxKey {
					maxKey = key
					maxValue = value
				}
			}
			cardMap[maxKey] += jokers
			delete(cardMap, '1')
		}
	}

	if len(cardMap) == 1 {
		return FiveOfAKind
	} else if len(cardMap) == 2 {
		for _, value := range cardMap {
			if value == 4 || value == 1 {
				return FourOfAKind
			}
			return FullHouse
		}
	} else if len(cardMap) == 3 {
		for _, value := range cardMap {
			if value == 3 {
				return ThreeOfAKind
			} else if value == 2 {
				return TwoPair
			}
		}
	} else if len(cardMap) == 4 {
		return OnePair
	}
	return HighCard
}

func parseHands(lines []string, part2 bool) ([]*Hand, error) {
	var hands = make([]*Hand, 0)
	for _, line := range lines {
		line = strings.ReplaceAll(line, "A", "Z")
		line = strings.ReplaceAll(line, "K", "Y")
		line = strings.ReplaceAll(line, "Q", "X")
		if part2 {
			line = strings.ReplaceAll(line, "J", "1")
		} else {
			line = strings.ReplaceAll(line, "J", "W")
		}
		splitted := strings.Split(line, " ")
		bid, err := strconv.Atoi(splitted[1])
		if err != nil {
			return nil, err
		}
		hand := &Hand{Cards: []rune(splitted[0]), Bid: bid}
		hand.Rank = hand.Evaluate(part2)
		hands = append(hands, hand)
	}
	return hands, nil
}

func SolveDay() error {
	lines, err := utils.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day7 with readstringfromargs(): %s", err.Error())
		return err
	}

	solutionPart1, err := solvePart1(lines)
	if err != nil {
		fmt.Printf("error at solve day7 with solvePart1(): %s\n", err.Error())
		return err
	}
	fmt.Printf("Die Lösung für Tag 7 Part 1 ist %d!\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("error at solve day7 with solvePart2(): %s\n", err.Error())
		return err
	}

	fmt.Printf("Die Lösung für Tag 7 Part 2 ist %d!\n", solutionPart2)

	return nil
}

func solvePart1(lines []string) (int, error) {
	hands, err := parseHands(lines, false)
	if err != nil {
		return -1, err
	}
	sort.Slice(hands, func(i, j int) bool {
		hand := hands[i]
		other := hands[j]
		if hand.Rank < other.Rank {
			return true
		} else if hand.Rank > other.Rank {
			return false
		}
		for i := 0; i < 5; i++ {
			if hand.Cards[i] > other.Cards[i] {
				return true
			} else if hand.Cards[i] < other.Cards[i] {
				return false
			}
		}
		return false
	})

	sumOfWinnings := 0

	for i := 0; i < len(hands); i++ {
		sumOfWinnings += (len(hands) - i) * hands[i].Bid
	}

	return sumOfWinnings, nil
}

func solvePart2(lines []string) (int, error) {
	hands, err := parseHands(lines, true)
	if err != nil {
		return -1, err
	}
	sort.Slice(hands, func(i, j int) bool {
		hand := hands[i]
		other := hands[j]
		if hand.Rank < other.Rank {
			return true
		} else if hand.Rank > other.Rank {
			return false
		}
		for i := 0; i < 5; i++ {
			if hand.Cards[i] > other.Cards[i] {
				return true
			} else if hand.Cards[i] < other.Cards[i] {
				return false
			}
		}
		return false
	})

	sumOfWinnings := 0

	for i := 0; i < len(hands); i++ {
		sumOfWinnings += (len(hands) - i) * hands[i].Bid
	}

	return sumOfWinnings, nil
}
