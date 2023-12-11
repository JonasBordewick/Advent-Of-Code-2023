package day10

import (
	"fmt"
	"github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"slices"
)

const (
	VerticalPipe int = iota
	HorizontalPipe
	NorthEastPipe
	NorthWestPipe
	SouthWestPipe
	SouthEastPipe
	Ground
	Start
	In
	Out
)

type Pipe struct {
	PipeType int
	XPos     int
	YPos     int
	Visited  bool
}

type Node struct {
	Neighbours []*Node
	Visited    bool
	Pipe       *Pipe
	Distance   int
}

func (pipe *Pipe) EvaluatePipeType(pipes [][]*Pipe) {
	if pipe.PipeType != Start {
		return
	}

	cTop, cLeft, cRight, cBottom := false, false, false, false

	if pipe.YPos > 0 {
		pipeType := pipes[pipe.YPos-1][pipe.XPos].PipeType
		if pipeType == VerticalPipe || pipeType == SouthEastPipe || pipeType == SouthWestPipe {
			cTop = true
		}
	}

	if pipe.YPos < len(pipes)-1 {
		pipeType := pipes[pipe.YPos+1][pipe.XPos].PipeType
		if pipeType == VerticalPipe || pipeType == NorthEastPipe || pipeType == NorthWestPipe {
			cBottom = true
		}
	}

	if pipe.XPos > 0 {
		pipeType := pipes[pipe.YPos][pipe.XPos-1].PipeType
		if pipeType == HorizontalPipe || pipeType == NorthEastPipe || pipeType == SouthEastPipe {
			cLeft = true
		}
	}

	if pipe.XPos < len(pipes[pipe.YPos])-1 {
		pipeType := pipes[pipe.YPos][pipe.XPos+1].PipeType
		if pipeType == HorizontalPipe || pipeType == NorthWestPipe || pipeType == SouthWestPipe {
			cRight = true
		}
	}

	if cTop && cBottom {
		pipe.PipeType = HorizontalPipe
	} else if cTop && cLeft {
		pipe.PipeType = NorthWestPipe
	} else if cTop && cRight {
		pipe.PipeType = NorthEastPipe
	} else if cBottom && cLeft {
		pipe.PipeType = SouthWestPipe
	} else if cBottom && cRight {
		pipe.PipeType = SouthEastPipe
	} else {
		pipe.PipeType = VerticalPipe
	}

}

func (pipe *Pipe) getGroundNeighbours(pipes *[][]*Pipe) []*Pipe {
	var neighbours []*Pipe = []*Pipe{}

	if pipe.YPos > 0 && (*pipes)[pipe.YPos-1][pipe.XPos].PipeType > SouthEastPipe {
		neighbours = append(neighbours, (*pipes)[pipe.YPos-1][pipe.XPos])
	}
	if pipe.YPos < len(*pipes)-1 && (*pipes)[pipe.YPos+1][pipe.XPos].PipeType > SouthEastPipe {
		neighbours = append(neighbours, (*pipes)[pipe.YPos+1][pipe.XPos])
	}
	if pipe.XPos > 0 && (*pipes)[pipe.YPos][pipe.XPos-1].PipeType > SouthEastPipe {
		neighbours = append(neighbours, (*pipes)[pipe.YPos][pipe.XPos-1])
	}
	if pipe.XPos < len((*pipes)[pipe.YPos])-1 && (*pipes)[pipe.YPos][pipe.XPos+1].PipeType > SouthEastPipe {
		neighbours = append(neighbours, (*pipes)[pipe.YPos][pipe.XPos+1])
	}

	return neighbours
}

func SolveDay() error {
	lines, err := utils.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day8 with readstringfromargs(): %s", err.Error())
		return err
	}

	pipes, startPipe, grounds := makePipes(lines)
	startPipe.EvaluatePipeType(pipes)

	graph := buildGraph(pipes, startPipe)

	graph, maxDistance := bfs(graph)

	fmt.Println(maxDistance)

	var ins []*Pipe = make([]*Pipe, 0)

	for _, ground := range grounds {
		if ground.YPos == 0 || ground.YPos == len(pipes)-1 || ground.XPos == 0 || ground.XPos == len(pipes[0])-1 {
			ground.PipeType = Out
		} else {
			ground.PipeType = In
			ins = append(ins, ground)
		}
	}

	var groups [][]*Pipe = make([][]*Pipe, 0)

	for _, in := range ins {
		if in.Visited {
			continue
		}
		var group = make([]*Pipe, 0)
		var queue = []*Pipe{in}
		var current *Pipe
		for len(queue) > 0 {
			current = queue[0]
			queue = queue[1:]
			current.Visited = true
			group = append(group, current)
			nextPipes := current.getGroundNeighbours(&pipes)
			for _, next := range nextPipes {
				if next.Visited {
					continue
				}
				if !slices.Contains(queue, next) {
					queue = append(queue, next)
				}
			}
		}
		groups = append(groups, group)
	}

	var sum = 0

	for _, group := range groups {
		var yPos = group[0].YPos
		var xPos = group[0].XPos

		var diffYPos = false
		var diffXPos = false

		for _, x := range group {
			result := yPos == x.YPos
			if !result {
				diffYPos = true
			}
			result = xPos == x.XPos
			if !result {
				diffXPos = true
			}
			if diffYPos && diffXPos {
				break
			}
		}

		if !(diffXPos && diffYPos) {
			sum += len(group)
		} else {
			for _, x := range group {
				x.PipeType = Out
			}
		}

	}

	fmt.Println(sum)

	printPipes(pipes)

	return nil
}

func printPipes(pipes [][]*Pipe) {
	for _, y := range pipes {
		for _, pipe := range y {
			switch pipe.PipeType {
			case VerticalPipe:
				fmt.Print("$")
			case Ground:
				fmt.Print("$")
			case HorizontalPipe:
				fmt.Print("$")
			case NorthEastPipe:
				fmt.Print("$")
			case NorthWestPipe:
				fmt.Print("$")
			case SouthWestPipe:
				fmt.Print("$")
			case SouthEastPipe:
				fmt.Print("$")
			case In:
				fmt.Print("1")
			case Out:
				fmt.Print("0")
			}
		}
		fmt.Print("\n")
	}
}

func makePipes(lines []string) ([][]*Pipe, *Pipe, []*Pipe) {
	pipes := make([][]*Pipe, 0)
	grounds := make([]*Pipe, 0)
	var startPipe *Pipe
	for yPos := 0; yPos < len(lines); yPos++ {
		currentLine := lines[yPos]
		pipesOfLine := make([]*Pipe, 0)
		for xPos := 0; xPos < len(currentLine); xPos++ {
			var currentChar = rune(currentLine[xPos])
			switch currentChar {
			case '|':
				pipesOfLine = append(pipesOfLine, &Pipe{PipeType: VerticalPipe, XPos: xPos, YPos: yPos})
			case '-':
				pipesOfLine = append(pipesOfLine, &Pipe{PipeType: HorizontalPipe, XPos: xPos, YPos: yPos})
			case 'L':
				pipesOfLine = append(pipesOfLine, &Pipe{PipeType: NorthEastPipe, XPos: xPos, YPos: yPos})
			case 'J':
				pipesOfLine = append(pipesOfLine, &Pipe{PipeType: NorthWestPipe, XPos: xPos, YPos: yPos})
			case '7':
				pipesOfLine = append(pipesOfLine, &Pipe{PipeType: SouthWestPipe, XPos: xPos, YPos: yPos})
			case 'F':
				pipesOfLine = append(pipesOfLine, &Pipe{PipeType: SouthEastPipe, XPos: xPos, YPos: yPos})
			case '.':
				var ground = &Pipe{PipeType: Ground, XPos: xPos, YPos: yPos}
				pipesOfLine = append(pipesOfLine, ground)
				grounds = append(grounds, ground)
			case 'S':
				startPipe = &Pipe{PipeType: Start, XPos: xPos, YPos: yPos}
				pipesOfLine = append(pipesOfLine, startPipe)
			}
		}
		pipes = append(pipes, pipesOfLine)
	}

	return pipes, startPipe, grounds
}

func buildGraph(pipes [][]*Pipe, start *Pipe) *Node {
	queue := make([]*Node, 0)
	startNode := &Node{
		Neighbours: make([]*Node, 0),
		Pipe:       start,
	}
	start.Visited = true
	queue = append(queue, startNode)
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		pipe := currentNode.Pipe
		var node *Node
		if pipe.PipeType == VerticalPipe {
			if !pipes[pipe.YPos-1][pipe.XPos].Visited {
				pipes[pipe.YPos-1][pipe.XPos].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos-1][pipe.XPos]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
			if !pipes[pipe.YPos+1][pipe.XPos].Visited {
				pipes[pipe.YPos+1][pipe.XPos].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos+1][pipe.XPos]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
		} else if pipe.PipeType == HorizontalPipe {
			if !pipes[pipe.YPos][pipe.XPos-1].Visited {
				pipes[pipe.YPos][pipe.XPos-1].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos][pipe.XPos-1]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
			if !pipes[pipe.YPos][pipe.XPos+1].Visited {
				pipes[pipe.YPos][pipe.XPos+1].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos][pipe.XPos+1]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
		} else if pipe.PipeType == NorthEastPipe {
			if !pipes[pipe.YPos-1][pipe.XPos].Visited {
				pipes[pipe.YPos-1][pipe.XPos].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos-1][pipe.XPos]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
			if !pipes[pipe.YPos][pipe.XPos+1].Visited {
				pipes[pipe.YPos][pipe.XPos+1].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos][pipe.XPos+1]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
		} else if pipe.PipeType == NorthWestPipe {
			if !pipes[pipe.YPos-1][pipe.XPos].Visited {
				pipes[pipe.YPos-1][pipe.XPos].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos-1][pipe.XPos]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
			if !pipes[pipe.YPos][pipe.XPos-1].Visited {
				pipes[pipe.YPos][pipe.XPos-1].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos][pipe.XPos-1]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
		} else if pipe.PipeType == SouthWestPipe {
			if !pipes[pipe.YPos+1][pipe.XPos].Visited {
				pipes[pipe.YPos+1][pipe.XPos].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos+1][pipe.XPos]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
			if !pipes[pipe.YPos][pipe.XPos-1].Visited {
				pipes[pipe.YPos+1][pipe.XPos].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos][pipe.XPos-1]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
		} else if pipe.PipeType == SouthEastPipe {
			if !pipes[pipe.YPos+1][pipe.XPos].Visited {
				pipes[pipe.YPos+1][pipe.XPos].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos+1][pipe.XPos]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
			if !pipes[pipe.YPos][pipe.XPos+1].Visited {
				pipes[pipe.YPos][pipe.XPos+1].Visited = true
				node = &Node{Neighbours: []*Node{currentNode}, Pipe: pipes[pipe.YPos][pipe.XPos+1]}
				queue = append(queue, node)
				currentNode.Neighbours = append(currentNode.Neighbours, node)
			}
		}
	}
	return startNode
}

func bfs(graph *Node) (*Node, int) {
	neighbours := graph.Neighbours
	distanceToStart := graph.Distance
	for len(neighbours) > 0 {
		distanceToStart++
		neighboursQueue := make([]*Node, 0)

		for _, n := range neighbours {
			if !n.Visited {
				n.Visited = true
				for _, neighbour := range n.Neighbours {
					if !neighbour.Visited {
						neighboursQueue = append(neighboursQueue, neighbour)
					}
				}
				n.Distance = distanceToStart
			}
		}

		neighbours = neighboursQueue
	}

	return graph, distanceToStart

}
