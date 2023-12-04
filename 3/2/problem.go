package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func readNumber(line *[]string, x int) int {
	number := ""

	// go backwards until we find start of number
	for ; x > 0 && isDigit((*line)[x-1]); x-- {
	}

	// go forward until digit ends
	for ; x < len(*line) && isDigit((*line)[x]); x++ {
		number += (*line)[x]
	}

	return toInt(number)
}

func getAdjacentNumbers(matrix *[][]string, x int, y int) []int {
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
	numbersSet := make(map[int]bool)

	for _, coord := range coords {
		if coord.x < 0 || coord.y < 0 || coord.y > max_y || coord.x > max_x {
			continue
		}

		char := (*matrix)[coord.y][coord.x]

		if isDigit(char) {
			number := readNumber(&((*matrix)[coord.y]), coord.x)
			numbersSet[number] = true
		}
	}

	numbers := []int{}
	for number, _ := range numbersSet {
		numbers = append(numbers, number)
	}
	return numbers
}

func findGearRatios(matrix [][]string) []int {
	gearRatios := []int{}

	for y, line := range matrix {
		for x, char := range line {
			if char == "*" {
				numbers := getAdjacentNumbers(&matrix, x, y)
				if len(numbers) == 2 {
					gearRatio := numbers[0] * numbers[1]
					gearRatios = append(gearRatios, gearRatio)
				}
			}
		}
	}

	return gearRatios
}

func main() {
	// n := []string{"2", "3", "4", "1", "2", "3", "4", "5", "6", "7"}

	// fmt.Println(readNumber(&n, 5))

	input := readInput()
	matrix := toMatrix(input)
	gearRatios := findGearRatios(matrix)

	sum := 0

	for _, gearRatio := range gearRatios {
		sum += gearRatio
	}

	fmt.Println(sum)
}
