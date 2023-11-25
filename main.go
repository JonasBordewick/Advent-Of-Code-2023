package main

import (
	"fmt"
	"github.com/JonasBordewick/Advent-Of-Code-2023/utils"
)

func main() {
	day, err := utils.GetDayFromArgs()
	if err != nil {
		fmt.Println("Usage: advent <day1-25> <path to file>")
		return
	}
	switch day {
	case 0:
		fmt.Println("Starting to solve advent of code puzzle of day 1")
	}
}
