package day13

import (
	"fmt"
	"github.com/JonasBordewick/Advent-Of-Code-2023/utils"
)

const (
	MirrorAxisHorizontal int = iota
	MirrorAxisVertical
)

type Pattern struct {
	Pattern        []string
	cols           []string
	MirrorAxisType int
	MirrorAxisAtIndex int
}

func (p *Pattern) search() {
	// Horizontale Spiegel Achse
	var rowMap map[string]int = make(map[string]int)
	var rowCounter = 0
	var x []int = make([]int, 0)
	var possibleHorizontalAxis []int = make([]int, 0)
	for i, row := range p.Pattern {
		if _, ok := rowMap[row]; !ok {
			rowMap[row] = rowCounter
			rowCounter++
		}
		if len(x) > 0 && rowMap[row] == x[len(x)-1] {
			possibleHorizontalAxis = append(possibleHorizontalAxis, i)
		}
		x = append(x, rowMap[row])
	}

	var colMap map[string]int = make(map[string]int)
	var colCounter = 0
	var y []int = make([]int, 0)
	var possibleVertAxis []int = make([]int, 0)
	for i, col := range p.cols {
		if _, ok := colMap[col]; !ok {
			colMap[col] = colCounter
			colCounter++
		}
		if len(y) > 0 && colMap[col] == y[len(y)-1] {
			possibleVertAxis = append(possibleVertAxis, i)
		}
		y = append(y, colMap[col])
	}

	horizontalAxis := -1

	for _, hori := range possibleHorizontalAxis {
		a := x[hori:]
		b := x[:hori]

		var possible = true

		for i := 0; i < len(a); i++ {
			if i >= len(b) {
				break
			}
			if a[i] != b[len(b)-(i+1)] && possible {
				possible = false
				break
			}
		}

		if possible {
			horizontalAxis = hori
			break
		}
	}

	verticalAxis := -1

	for _, vert := range possibleVertAxis {
		a := y[vert:]
		b := y[:vert]

		var possible = true

		for i := 0; i < len(a); i++ {
			if i >= len(b) {
				break
			}
			if a[i] != b[len(b)-(i+1)] && possible {
				possible = false
				break
			}
		}

		if possible {
			verticalAxis = vert
			break
		}
	}

	if verticalAxis == -1 && horizontalAxis == -1 {
		fmt.Println("Keine Achse gefunden")
	}
	if verticalAxis > -1 && horizontalAxis > -1 {
		fmt.Println("Zwei Achsen gefunden")
	}

	if verticalAxis > -1 {
		p.MirrorAxisType = MirrorAxisVertical
		p.MirrorAxisAtIndex = verticalAxis
		return
	}
	p.MirrorAxisType = MirrorAxisHorizontal
	p.MirrorAxisAtIndex = horizontalAxis
}

func (p *Pattern) findSmudge() {
	var a string = ""
	var b string = ""
	for i := 0; i < len(p.Pattern); i++ {
		a = p.Pattern[i]
		var hemmingDiff bool = false
		for j := i + 1; j < len(p.Pattern); j++ {
			b = p.Pattern[j]
			if diffOfString(a, b) == 1 {
				hemmingDiff = true
				break
			}
		}
		if hemmingDiff {
			break
		}
	}
	fmt.Println(a)
	fmt.Println(b)

}

func diffOfString(a, b string) int {
	if len(a) != len(b) {
		return -1
	}

	if a == b {
		return 0
	}

	var diff = 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff
}

func SolveDay() error {
	lines, err := utils.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day8 with readstringfromargs(): %s", err.Error())
		return err
	}

	patterns := make([]*Pattern, 0)

	currentPattern := &Pattern{Pattern: make([]string, 0)}

	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, currentPattern)
			currentPattern = &Pattern{Pattern: make([]string, 0)}
			continue
		}
		currentPattern.Pattern = append(currentPattern.Pattern, line)
	}
	patterns = append(patterns, currentPattern)

	sum := 0

	for _, pattern := range patterns {
		var cols []string = make([]string, 0)
		for i := 0; i < len(pattern.Pattern[0]); i++ {
			col := ""
			for _, row := range pattern.Pattern {
				col += string(row[i])
			}
			cols = append(cols, col)
		}
		pattern.search()
		if pattern.MirrorAxisType == MirrorAxisVertical {
			sum += pattern.MirrorAxisAtIndex
		} else {
			sum += (100 * pattern.MirrorAxisAtIndex)
		}
	}

	fmt.Printf("Sum of all notes %d\n", sum)

	for _, pattern := range patterns {
		pattern.findSmudge()
	}

	return nil
}
