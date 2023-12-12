package day11

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

	universe := make([][]*struct {
		YPos int
		XPos int
		Type int
	}, 0)

	for y := 0; y < len(lines); y++ {
		row := make([]*struct {
			YPos int
			XPos int
			Type int
		}, 0)
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == '#' {
				row = append(row, &struct {
					YPos int
					XPos int
					Type int
				}{
					YPos: y,
					XPos: x,
					Type: 1,
				})
			} else {
				row = append(row, &struct {
					YPos int
					XPos int
					Type int
				}{
					YPos: y,
					XPos: x,
					Type: 0,
				})
			}
		}
		universe = append(universe, row)
	}

	emptyRows := []int{}

	for y := 0; y < len(universe); y++ {
		space := true
		for x := 0; x < len(universe[y]); x++ {
			if universe[y][x].Type == 1 {
				space = false
				break
			}
		}
		if space {
			emptyRows = append(emptyRows, y)
		}
	}
	emptyColumns := []int{}
	for x := 0; x < len(universe[0]); x++ {
		space := true
		for y := 0; y < len(universe); y++ {
			if universe[y][x].Type != 0 {
				space = false
				break
			}
		}
		if space {
			emptyColumns = append(emptyColumns, x)
		}
	}

	galaxies := make([]*struct {
		YPos int
		XPos int
		Type int
	}, 0)
	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[y]); x++ {

			element := universe[y][x]

			rowsBelow := 0
			columnsBelow := 0

			for i := 0; i < len(emptyRows); i++ {
				if emptyRows[i] < y {
					rowsBelow++
				} else {
					break
				}
			}

			for i := 0; i < len(emptyColumns); i++ {
				if emptyColumns[i] < x {
					columnsBelow++
				} else {
					break
				}
			}

			element.YPos = y + rowsBelow*1
			element.XPos = x + columnsBelow*1

			if element.Type == 1 {
				galaxies = append(galaxies, element)
			}

		}
	}

	distances := 0

	for i := 0; i < len(galaxies); i++ {
		element := galaxies[i]
		for _, other := range galaxies[i+1:] {
			distances += abs(other.YPos-element.YPos) + abs(other.XPos-element.XPos)
		}
	}

	fmt.Printf("the sum of the shortest path between all %d pairs of galaxies is %d\n", ((len(galaxies)-1)*len(galaxies))/2, distances)

	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[y]); x++ {

			element := universe[y][x]

			rowsBelow := 0
			columnsBelow := 0

			for i := 0; i < len(emptyRows); i++ {
				if emptyRows[i] < y {
					rowsBelow++
				} else {
					break
				}
			}

			for i := 0; i < len(emptyColumns); i++ {
				if emptyColumns[i] < x {
					columnsBelow++
				} else {
					break
				}
			}

			element.YPos = y + rowsBelow*(1000000-1)
			element.XPos = x + columnsBelow*(1000000-1)

		}
	}

	distances = 0

	for i := 0; i < len(galaxies); i++ {
		element := galaxies[i]
		for _, other := range galaxies[i+1:] {
			distances += abs(other.YPos-element.YPos) + abs(other.XPos-element.XPos)
		}
	}

	fmt.Printf("the sum of the shortest path between all %d pairs of galaxies is %d\n", ((len(galaxies)-1)*len(galaxies))/2, distances)

	return nil
}

func abs(value int) int {
	if value < 0 {
		return -1 * value
	}
	return value
}
