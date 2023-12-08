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

func main() {
	input := readInput()
	directions, nodes := parseInput(input)

	location := "AAA"
	steps := 0
	for location != "ZZZ" {
		di := steps % len(directions)
		direction := directions[di]

		if direction == "L" {
			location = nodes[location].left
		} else {
			location = nodes[location].right
		}

		steps++
	}

	fmt.Println(steps)
}
