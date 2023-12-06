package day6

import (
	"fmt"
	utils "github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"math"
	"strconv"
)

func SolveDay6() error {
	lines, err := utils.ReadStringFromArgs()
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

type Record struct {
	Distance   int
	RecordTime int
}

func generateRecords(lines []string) ([]*Record, error) {
	times, _, err := utils.MakeIntSlice(lines[0])
	if err != nil {
		return nil, err
	}
	distances, _, err := utils.MakeIntSlice(lines[1])
	if err != nil {
		return nil, err
	}

	races := make([]*Record, 0)

	for i := range times {
		races = append(races, &Record{RecordTime: int(times[i]), Distance: int(distances[i])})
	}

	return races, nil
}

func solvePart1(lines []string) (int, error) {
	records, err := generateRecords(lines)
	if err != nil {
		return -1, err
	}

	var result = 1
	for _, record := range records {
		var waysToWin = calc(record)
		result *= waysToWin
	}

	return result, nil
}

func solvePart2(lines []string) (int, error) {
	records, err := generateRecords(lines)
	if err != nil {
		return -1, err
	}

	timeString := ""
	distanceString := ""

	for _, record := range records {
		timeString += fmt.Sprintf("%d", record.RecordTime)
		distanceString += fmt.Sprintf("%d", record.Distance)
	}

	t, err := strconv.Atoi(timeString)
	if err != nil {
		return -1, err
	}
	d, err := strconv.Atoi(distanceString)
	if err != nil {
		return -1, err
	}

	record := &Record{RecordTime: t, Distance: d}
	waysToWin := calc(record)
	return waysToWin, nil
}

func calc(record *Record) int {
	// x * (t - x) = d
	// -x^2 + tx = d
	// x^2 - tx + d = 0
	// x gesucht | t = record.Time | d record.Distance
	discriminant := math.Pow(float64(-record.RecordTime), 2) - (4 * float64(record.Distance))
	sqrt := math.Sqrt(discriminant)
	h1 := int((float64(record.RecordTime) - sqrt) / 2)
	h2 := int((float64(record.RecordTime) + sqrt) / 2)

	if h1 > h2 {
		return h1 - h2
	}
	return h2 - h1
}
