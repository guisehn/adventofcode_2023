package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func readInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}

func toMatrix(input string) [][]string {
	matrix := [][]string{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}

	return matrix
}

var digitRe = regexp.MustCompile("[0-9]")

func isDigit(str string) bool {
	return digitRe.MatchString(str)
}

func hasAdjacentSymbol(matrix *[][]string, x int, y int) bool {
	coords := []struct {
		x int
		y int
	}{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		{x - 1, y},
		{x + 1, y},
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}

	max_x := len((*matrix)[y]) - 1
	max_y := len(*matrix) - 1

	for _, coord := range coords {
		if coord.x < 0 || coord.y < 0 || coord.y > max_y || coord.x > max_x {
			continue
		}

		char := (*matrix)[coord.y][coord.x]
		if char != "." && !isDigit(char) {
			return true
		}
	}

	return false
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func findPartNumbers(matrix [][]string) []int {
	acc := ""
	capturing := false
	foundSymbol := false
	partNumbers := []int{}

	for y, line := range matrix {
		for x, char := range line {
			if isDigit(char) {
				acc += char
				capturing = true

				if !foundSymbol && hasAdjacentSymbol(&matrix, x, y) {
					foundSymbol = true
				}
			} else if capturing {
				if foundSymbol {
					partNumbers = append(partNumbers, toInt(acc))
				}

				acc = ""
				capturing = false
				foundSymbol = false
			}
		}
	}

	return partNumbers
}

func main() {
	input := readInput()
	matrix := toMatrix(input)
	partNumbers := findPartNumbers(matrix)
	sum := 0

	for _, partNumber := range partNumbers {
		sum += partNumber
	}

	fmt.Println(sum)
}
