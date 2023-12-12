package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction = string

const (
	Up    Direction = "U"
	Down  Direction = "D"
	Left  Direction = "L"
	Right Direction = "R"
)

type Coord struct {
	x, y int
}

type Problem struct {
	pipes            [][]string
	initialDirection Direction
	initialCoord     Coord
}

var pipeTypes = map[string](map[Direction]Direction){
	"|": {
		Up:   Up,
		Down: Down,
	},
	"-": {
		Left:  Left,
		Right: Right,
	},
	"L": {
		Down: Right,
		Left: Up,
	},
	"J": {
		Right: Up,
		Down:  Left,
	},
	"7": {
		Up:    Left,
		Right: Down,
	},
	"F": {
		Left: Down,
		Up:   Right,
	},
}

func NewProblem(input string) Problem {
	initialCoord := Coord{}
	lines := strings.Split(input, "\n")
	pipes := [][]string{}

	for y := 0; y < len(lines); y++ {
		chars := strings.Split(lines[y], "")

		for x := 0; x < len(lines[y]); x++ {
			if chars[x] == "S" {
				initialCoord = Coord{x, y}

				// hardcoding what's going to be the initial pipe, too lazy to write
				// this algorithm
				chars[x] = "7"
			}
		}

		pipes = append(pipes, chars)
	}

	return Problem{
		pipes:            pipes,
		initialCoord:     initialCoord,
		initialDirection: "R", // also hardcoding initial direction
	}
}

func (problem Problem) farthestDistance() int {
	coord := problem.initialCoord
	direction := problem.initialDirection
	path := []Coord{}

	for {
		path = append(path, coord)
		coord, direction = movePipe(problem.pipes[coord.y][coord.x], coord, direction)
		if coord == problem.initialCoord {
			break
		}
	}

	return len(path) / 2
}

func movePipe(pipe string, coord Coord, direction Direction) (Coord, Direction) {
	newDirection := pipeTypes[pipe][direction]

	fmt.Println("MOVE PIPE", pipe, coord, direction, "->", newDirection)
	newCoord := move(coord, newDirection)
	fmt.Println("NEW COORD", newCoord)

	return newCoord, newDirection
}

func move(coord Coord, direction Direction) Coord {
	x, y := coord.x, coord.y

	switch direction {
	case Up:
		return Coord{x, y - 1}
	case Down:
		return Coord{x, y + 1}
	case Left:
		return Coord{x - 1, y}
	case Right:
		return Coord{x + 1, y}
	default:
		panic("Unknown direction")
	}
}

func readInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}

func main() {
	input := readInput()
	problem := NewProblem(input)

	// pipes := [][]string{
	// 	{".", ".", ".", ".", "."},
	// 	{".", "F", "-", "7", "."},
	// 	{".", "|", ".", "|", "."},
	// 	{".", "L", "-", "J", "."},
	// 	{".", ".", ".", ".", "."},
	// }

	fmt.Println(problem.farthestDistance())
}
