package day2

import (
	"fmt"
	u "github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"regexp"
	"strconv"
	"strings"
)

type GameBag struct {
	RedCubes   int
	GreenCubes int
	BlueCubes  int
}

type Game struct {
	Id                 int
	RevealedRedCubes   []int
	RevealedGreenCubes []int
	RevealedBlueCubes  []int
}

func (game *Game) Play(bag *GameBag) (int, int) {
	var maxRounds int = 0
	if len(game.RevealedBlueCubes) > maxRounds {
		maxRounds = len(game.RevealedBlueCubes)
	}
	if len(game.RevealedRedCubes) > maxRounds {
		maxRounds = len(game.RevealedRedCubes)
	}
	if len(game.RevealedGreenCubes) > maxRounds {
		maxRounds = len(game.RevealedGreenCubes)
	}

	var maxRed, maxGreen, maxBlue = 1, 1, 1

	var valid = true

	for round := 0; round < maxRounds; round++ {
		if round < len(game.RevealedBlueCubes) {
			if bag.BlueCubes < game.RevealedBlueCubes[round] {
				valid = false
			}
			if maxBlue < game.RevealedBlueCubes[round] {
				maxBlue = game.RevealedBlueCubes[round]
			}
		}
		if round < len(game.RevealedRedCubes) {
			if bag.RedCubes < game.RevealedRedCubes[round] {
				valid = false
			}
			if maxRed < game.RevealedRedCubes[round] {
				maxRed = game.RevealedRedCubes[round]
			}
		}
		if round < len(game.RevealedGreenCubes) {
			if bag.GreenCubes < game.RevealedGreenCubes[round] {
				valid = false
			}
			if maxGreen < game.RevealedGreenCubes[round] {
				maxGreen = game.RevealedGreenCubes[round]
			}
		}
	}

	if valid {
		return game.Id, maxRed * maxBlue * maxGreen
	}
	return 0, maxRed * maxBlue * maxGreen

}

func SolveDayTwo() error {
	lines, err := u.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day2 with readstringfromargs(): %s", err.Error())
		return err
	}

	solutionPart1, err := solvePart1(lines)
	if err != nil {
		fmt.Printf("error at solve day2 with solvePart1(): %s\n", err.Error())
		return err
	}
	fmt.Printf("Die Lösung für Tag 2 Part 1 ist %d!\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("error at solve day2 with solvePart2(): %s\n", err.Error())
		return err
	}

	fmt.Printf("Die Lösung für Tag 2 Part 2 ist %d!\n", solutionPart2)

	return nil
}

func solvePart1(lines []string) (int, error) {
	bag := &GameBag{RedCubes: 12, GreenCubes: 13, BlueCubes: 14}
	games, err := generateGames(lines)
	if err != nil {
		return -1, err
	}

	var sum = 0

	for _, game := range games {
		result, _ := game.Play(bag)
		sum += result
	}

	return sum, nil
}

func generateGames(lines []string) ([]*Game, error) {
	var games []*Game = []*Game{}
	for _, line := range lines {
		gameID, err := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		if err != nil {
			return nil, err
		}
		var game *Game = &Game{Id: gameID, RevealedRedCubes: make([]int, 0), RevealedGreenCubes: make([]int, 0), RevealedBlueCubes: make([]int, 0)}

		rounds := strings.Split(strings.Split(line, ":")[1], ";")
		var re = regexp.MustCompile(`(?m)\d+|green|blue|red`)
		for _, round := range rounds {
			cubes := strings.Split(round, ",")
			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")
				matches := re.FindAllStringSubmatch(cube, -1)
				number, err := strconv.Atoi(matches[0][0])
				if err != nil {
					return nil, err
				}
				color := matches[1][0]

				switch color {
				case "red":
					game.RevealedRedCubes = append(game.RevealedRedCubes, number)
				case "green":
					game.RevealedGreenCubes = append(game.RevealedGreenCubes, number)
				case "blue":
					game.RevealedBlueCubes = append(game.RevealedBlueCubes, number)
				}
			}
		}

		games = append(games, game)
	}
	return games, nil
}

func solvePart2(lines []string) (int, error) {
	bag := &GameBag{RedCubes: 12, GreenCubes: 13, BlueCubes: 14}
	games, err := generateGames(lines)
	if err != nil {
		return -1, err
	}

	var sum = 0

	for _, game := range games {
		_, result := game.Play(bag)
		sum += result
	}

	return sum, nil
}
