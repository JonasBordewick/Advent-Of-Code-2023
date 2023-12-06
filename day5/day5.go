package day5

import (
	"fmt"
	u "github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"math"
	"sync"
)

func SolveDayFive() error {
	lines, err := u.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day5 with readstringfromargs(): %s", err.Error())
		return err
	}

	solutionPart1, err := solvePart1(lines)
	if err != nil {
		fmt.Printf("error at solve day5 with solvePart1(): %s\n", err.Error())
		return err
	}
	fmt.Printf("Die Lösung für Tag 5 Part 1 ist %d!\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("error at solve day4 with solvePart2(): %s\n", err.Error())
		return err
	}

	fmt.Printf("Die Lösung für Tag 5 Part 2 ist %d!\n", solutionPart2)

	return nil
}

type AlmanacContent struct {
	DestinationStart int64
	SourceStart      int64
	Range            int64
}

func makeContent(lines []string, lineIndex int) ([]*AlmanacContent, int, error) {
	var result []*AlmanacContent = make([]*AlmanacContent, 0)

	for lineIndex < len(lines) && lines[lineIndex] != "" {
		numbers, _, err := u.MakeIntSlice(lines[lineIndex])
		if err != nil {
			return nil, -1, err
		}
		result = append(result, &AlmanacContent{DestinationStart: numbers[0], SourceStart: numbers[1], Range: numbers[2]})
		lineIndex++
	}
	lineIndex += 2

	return result, lineIndex, nil
}

func parseLines(lines []string) ([]int64, []*AlmanacContent, []*AlmanacContent, []*AlmanacContent, []*AlmanacContent, []*AlmanacContent, []*AlmanacContent, []*AlmanacContent, error) {
	seeds, _, err := u.MakeIntSlice(lines[0])
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	var lineIndex int = 3

	seedToSoil, lineIndex, err := makeContent(lines, lineIndex)

	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	soilToFertilizer, lineIndex, err := makeContent(lines, lineIndex)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	fertilizerToWater, lineIndex, err := makeContent(lines, lineIndex)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	waterToLight, lineIndex, err := makeContent(lines, lineIndex)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	lightToTemp, lineIndex, err := makeContent(lines, lineIndex)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	tempToHum, lineIndex, err := makeContent(lines, lineIndex)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	humToLocation, lineIndex, err := makeContent(lines, lineIndex)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	return seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHum, humToLocation, nil
}

func calcMapValue(source int64, selectedMap []*AlmanacContent) int64 {
	var dest int64 = source
	for _, content := range selectedMap {
		if source >= content.SourceStart && source < content.SourceStart+content.Range {
			dest = content.DestinationStart + (source - content.SourceStart)
			break
		}
	}
	return dest
}

func seedToLocation(seed int64, maps map[string][]*AlmanacContent) int64 {
	soil := calcMapValue(seed, maps["sts"])
	fert := calcMapValue(soil, maps["stf"])
	wate := calcMapValue(fert, maps["ftw"])
	ligh := calcMapValue(wate, maps["wtl"])
	temp := calcMapValue(ligh, maps["ltt"])
	humi := calcMapValue(temp, maps["tth"])
	location := calcMapValue(humi, maps["htl"])
	return location
}

func solvePart1(lines []string) (int64, error) {
	seeds, sts, stf, ftw, wtl, ltt, tth, htl, err := parseLines(lines)
	var maps map[string][]*AlmanacContent = make(map[string][]*AlmanacContent)
	maps["sts"] = sts
	maps["stf"] = stf
	maps["ftw"] = ftw
	maps["wtl"] = wtl
	maps["ltt"] = ltt
	maps["tth"] = tth
	maps["htl"] = htl
	if err != nil {
		return -1, err
	}
	var location int64 = math.MaxInt64
	for _, seed := range seeds {
		var tmp int64 = seedToLocation(seed, maps)
		if location > tmp {
			location = tmp
		}
	}
	return location, nil
}

func solvePart2(lines []string) (int64, error) {
	seeds, sts, stf, ftw, wtl, ltt, tth, htl, err := parseLines(lines)
	if err != nil {
		return -1, err
	}
	var maps map[string][]*AlmanacContent = make(map[string][]*AlmanacContent)
	maps["sts"] = sts
	maps["stf"] = stf
	maps["ftw"] = ftw
	maps["wtl"] = wtl
	maps["ltt"] = ltt
	maps["tth"] = tth
	maps["htl"] = htl
	var wg = sync.WaitGroup{}
	wg.Add(int(len(seeds) / 2))

	distances := make([]int64, 0)

	for i := 0; i < len(seeds); i += 2 {
		go func(rangeStart, rangeEnd int64, maps map[string][]*AlmanacContent) {
			defer wg.Done()

			var location int64 = math.MaxInt64
			for seed := rangeStart; seed < rangeEnd; seed++ {
				var tmp int64 = seedToLocation(seed, maps)
				if tmp < location {
					location = tmp
				}
			}
			fmt.Printf("Add Location %d\n", location)
			distances = append(distances, location)
		}(seeds[i], seeds[i]+seeds[i+1], maps)
	}
	wg.Wait()

	var min int64 = math.MaxInt64
	for i := range distances {
		location := distances[i]
		if location < min {
			min = location
		}
	}
	return min, nil
}
