package day8

import (
	"fmt"
	"github.com/JonasBordewick/Advent-Of-Code-2023/utils"
	"regexp"
	"slices"
)

type Node struct {
	Left  string
	Value string
	Right string
}

func (n *Node) walk(route string, nodes map[string]*Node, endNode string) int {
	foundZZZ := false

	routeLen := len(route)
	routeIndex := 0
	steps := 0

	currentNode := n

	for !foundZZZ {
		if routeIndex == routeLen {
			routeIndex = 0
		}

		if route[routeIndex] == 'L' {
			currentNode = nodes[currentNode.Left]
		} else {
			currentNode = nodes[currentNode.Right]
		}

		steps++

		if currentNode.Value == endNode {
			foundZZZ = true
		}
		routeIndex++
	}

	return steps
}

func parseInput(lines []string) (string, map[string]*Node) {
	route := lines[0]

	var nodeMap map[string]*Node = make(map[string]*Node)

	var re = regexp.MustCompile(`(?m)([A-Z0-9]+)\s*=\s*\(([^,]+),\s*([^)]+)\)`)

	for i := 2; i < len(lines); i++ {
		match := re.FindAllStringSubmatch(lines[i], -1)
		nodeMap[match[0][1]] = &Node{Value: match[0][1], Left: match[0][2], Right: match[0][3]}
	}

	return route, nodeMap
}

func findDirection(route string, move int) string {
	return string(route[move%len(route)])
}

func euclideanGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(numbers []int, index int) int {
	if index == len(numbers)-1 {
		return numbers[index]
	}
	a := numbers[index]
	b := LCM(numbers, index+1)
	return (a * b) / euclideanGCD(a, b)
}

func SolveDay() error {
	lines, err := utils.ReadStringFromArgs()
	if err != nil {
		fmt.Printf("error at solve day8 with readstringfromargs(): %s", err.Error())
		return err
	}

	route, nodes := parseInput(lines)
	//steps := nodes["AAA"].walk(route, nodes, "ZZZ")
	//fmt.Printf("Die Lösung für Tag 8 Part 1 ist %d!\n", steps)

	var keys []string = make([]string, 0)

	for k, _ := range nodes {
		keys = append(keys, k)
	}

	var startRe = regexp.MustCompile("[A-Z0-9][A-Z0-9]A")
	var endRe = regexp.MustCompile("[A-Z0-9][A-Z0-9]Z")

	startNodes := make([]string, 0)
	endNodes := make([]string, 0)

	for _, nodeValue := range keys {
		if startRe.FindString(nodeValue) != "" {
			startNodes = append(startNodes, nodeValue)
		} else if endRe.FindString(nodeValue) != "" {
			endNodes = append(endNodes, nodeValue)
		}
	}
	fmt.Println(startNodes)
	fmt.Println(endNodes)
	var moves []int = make([]int, 0)
	for i, startNode := range startNodes {
		fmt.Printf("Start %s | End %s\n", startNode, endNodes[i])
		node := startNode
		move := 0
		for !slices.Contains(endNodes, node) {
			direction := findDirection(route, move)
			if direction == "L" {
				node = nodes[node].Left
			} else {
				node = nodes[node].Right
			}
			move++
		}
		moves = append(moves, move)
	}

	fmt.Println(moves)

	fmt.Println(LCM(moves, 0))

	return nil
}
