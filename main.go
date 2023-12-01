package main

import (
	"fmt"
	"github.com/JonasBordewick/Advent-Of-Code-2023/day1"
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
		fmt.Printf("the puzzle has not yet been solved")
	case 2:
		fmt.Printf("the puzzle has not yet been solved")
	case 3:
		fmt.Printf("the puzzle has not yet been solved")
	case 4:
		fmt.Printf("the puzzle has not yet been solved")
	case 5:
		fmt.Printf("the puzzle has not yet been solved")
	case 6:
		fmt.Printf("the puzzle has not yet been solved")
	case 7:
		fmt.Printf("the puzzle has not yet been solved")
	case 8:
		fmt.Printf("the puzzle has not yet been solved")
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
