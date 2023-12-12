package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func readInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}

func parseInput(str string) ([]string, map[string]Node) {
	lines := strings.Split(str, "\n")
	directions := strings.Split(lines[0], "")

	nodes := make(map[string]Node)
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		nodes[line[0:3]] = Node{left: line[7:10], right: line[12:15]}
	}

	return directions, nodes
}

func locationsEndingWith(nodes map[string]Node, char string) []string {
	result := []string{}

	for node := range nodes {
		if node[2:3] == char {
			result = append(result, node)
		}
	}

	return result
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(integers []int) int {
	result := integers[0] * integers[1] / gcd(integers[0], integers[1])

	if len(integers) > 2 {
		for i := 0; i < len(integers); i++ {
			result = lcm(append([]int{result}, integers[2:]...))
		}
	}

	return result
}

func main() {
	input := readInput()
	directions, nodes := parseInput(input)

	locations := locationsEndingWith(nodes, "A")
	stepCounts := []int{}

	for i := 0; i < len(locations); i++ {
		location := locations[i]
		steps := 0
		firstZ := ""
		stepsTilFirstZ := -1
		stepsTilCycle := -1
		path := []string{location}

		for {
			di := steps % len(directions)
			direction := directions[di]

			if direction == "L" {
				location = nodes[location].left
			} else {
				location = nodes[location].right
			}

			path = append(path, location)
			steps++

			if location[2:3] == "Z" {
				if stepsTilFirstZ == -1 {
					firstZ = location
					stepsTilFirstZ = steps
				} else if location == firstZ {
					stepsTilCycle = steps - stepsTilFirstZ
					stepCounts = append(stepCounts, stepsTilCycle)
					break
				}
			}
		}

		// fmt.Println("Path", path, "Cycle found", stepsTilFirstZ, stepsTilCycle)
	}

	fmt.Println(lcm(stepCounts))
}
