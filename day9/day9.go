package day9

import (
	"fmt"
	"github.com/JonasBordewick/Advent-Of-Code-2023/utils"
)

func SolveDay() error {
	lines, err := utils.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day8 with readstringfromargs(): %s", err.Error())
		return err
	}

	var histories [][]int64 = make([][]int64, 0)
	for _, line := range lines {
		history, _, err := utils.MakeIntSlice(line)
		if err != nil {
			return err
		}
		histories = append(histories, history)
	}

	var sum int64 = 0

	for _, history := range histories {
		sum += extrapolate(history)
	}
	fmt.Printf("Das Ergebnis für Tag9 Part 1 ist %d\n", sum)
	sum = 0
	for _, history := range histories {
		sum += extrapolateBackwards(history)
	}
	fmt.Printf("Das Ergebnis für Tag9 Part 2 ist %d\n", sum)

	return nil
}

func extrapolate(history []int64) int64 {
	var values [][]int64 = buildPyramid(history)
	var extrapolateValue int64 = 0
	for i := len(values) - 1; i >= 0; i-- {
		value := extrapolateValue + values[i][len(values[i])-1]
		extrapolateValue = value
	}
	return extrapolateValue
}

func extrapolateBackwards(history []int64) int64 {
	var values [][]int64 = buildPyramid(history)
	var extrapolateValue int64 = 0
	for i := len(values) - 1; i >= 0; i-- {
		value := values[i][0] - extrapolateValue
		extrapolateValue = value
	}

	return extrapolateValue
}

func buildPyramid(history []int64) [][]int64 {
	var values [][]int64 = [][]int64{history}
	var depth int = 0
	for {
		var diffs []int64 = make([]int64, 0)
		for i := 1; i < len(values[depth]); i++ {
			diffs = append(diffs, values[depth][i]-values[depth][i-1])
		}
		values = append(values, diffs)
		if utils.All(0, diffs) {
			break
		}
		depth++
	}
	return values
}
