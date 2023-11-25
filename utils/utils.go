package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetDayFromArgs() (int, error) {
	if len(os.Args) < 2 {
		// There are no arguments (except of "application")
		return -1, errors.New("there is no filepath in args")
	}
	parts := strings.Split(os.Args[1], "day")
	if len(parts) != 2 {
		return 0, fmt.Errorf("wrong format(%s) should be dayX", os.Args[1])
	}
	dayNumber, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	if dayNumber < 1 || dayNumber > 25 {
		return 0, fmt.Errorf("wrong day %d, there are only 25 days of advent", dayNumber)
	}
	return dayNumber - 1, nil
}

func ReadStringFromArgs() ([]string, error) {
	// Check if there are arguments passed
	if len(os.Args) < 3 {
		// There are no arguments (except of "application")
		return nil, errors.New("there is no filepath in args")
	}
	filepath := os.Args[2]
	return ReadStringFromFile(filepath)
}

func ReadStringFromFile(filepath string) ([]string, error) {
	// Check if the passed file exists
	if _, err := os.Stat(filepath); err != nil {
		return nil, err
	}
	// Assertion: File exists
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	// Useful to close file when this Method is closed
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
