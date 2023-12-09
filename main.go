package main

import (
	"fmt"
	"github.com/JonasBordewick/Advent-Of-Code-2023/day1"
	"github.com/JonasBordewick/Advent-Of-Code-2023/day2"
	"github.com/JonasBordewick/Advent-Of-Code-2023/day3"
	"github.com/JonasBordewick/Advent-Of-Code-2023/day4"
	"github.com/JonasBordewick/Advent-Of-Code-2023/day5"
	"github.com/JonasBordewick/Advent-Of-Code-2023/day6"
	"github.com/JonasBordewick/Advent-Of-Code-2023/day7"
	"github.com/JonasBordewick/Advent-Of-Code-2023/day8"
	"github.com/JonasBordewick/Advent-Of-Code-2023/day9"
	"github.com/JonasBordewick/Advent-Of-Code-2023/utils"
)

func main() {
	day, err := utils.GetDayFromArgs()
	if err != nil {
		fmt.Println("Usage: advent <day1-25> <path to file>")
		return
	}
	fmt.Printf("Starting to solve advent of code puzzle of day %d\n", day+1)
	switch day {
	case 0:
		err := day1.SolveDayOne()
		if err != nil {
			return
		}
	case 1:
		err := day2.SolveDayTwo()
		if err != nil {
			return
		}
	case 2:
		err := day3.SolveDayThree()
		if err != nil {
			return
		}
	case 3:
		err := day4.SolveDayFour()
		if err != nil {
			return
		}
	case 4:
		err := day5.SolveDayFive()
		if err != nil {
			return
		}
	case 5:
		err := day6.SolveDay6()
		if err != nil {
			return
		}
	case 6:
		err := day7.SolveDay()
		if err != nil {
			return
		}
	case 7:
		err := day8.SolveDay()
		if err != nil {
			return
		}
	case 8:
		err := day9.SolveDay()
		if err != nil {
			return
		}
	case 9:
		fmt.Printf("the puzzle has not yet been solved")
	case 10:
		fmt.Printf("the puzzle has not yet been solved")
	case 11:
		fmt.Printf("the puzzle has not yet been solved")
	case 12:
		fmt.Printf("the puzzle has not yet been solved")
	case 13:
		fmt.Printf("the puzzle has not yet been solved")
	case 14:
		fmt.Printf("the puzzle has not yet been solved")
	case 15:
		fmt.Printf("the puzzle has not yet been solved")
	case 16:
		fmt.Printf("the puzzle has not yet been solved")
	case 17:
		fmt.Printf("the puzzle has not yet been solved")
	case 18:
		fmt.Printf("the puzzle has not yet been solved")
	case 19:
		fmt.Printf("the puzzle has not yet been solved")
	case 20:
		fmt.Printf("the puzzle has not yet been solved")
	case 21:
		fmt.Printf("the puzzle has not yet been solved")
	case 22:
		fmt.Printf("the puzzle has not yet been solved")
	case 23:
		fmt.Printf("the puzzle has not yet been solved")
	case 24:
		fmt.Printf("the puzzle has not yet been solved")
	}
}
