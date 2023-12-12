package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func predictNextNumber(history []int) int {
	sequence := history
	lastNums := []int{history[len(history)-1]}

	// fmt.Println(sequence)
	for !isAllZero(sequence) {
		diffs := []int{}
		for i := 0; i < len(sequence)-1; i++ {
			diffs = append(diffs, sequence[i+1]-sequence[i])
		}
		sequence = diffs
		// fmt.Println(sequence)
		lastNums = append([]int{sequence[len(sequence)-1]}, lastNums...)
	}

	result := 0
	for _, n := range lastNums {
		result += n
	}
	return result
}

func isAllZero(list []int) bool {
	for _, n := range list {
		if n != 0 {
			return false
		}
	}

	return true
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func toInts(strings []string) []int {
	ints := []int{}
	for _, str := range strings {
		ints = append(ints, toInt(str))
	}
	return ints
}

func readInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}

func parseInput(str string) [][]int {
	lines := strings.Split(str, "\n")
	items := [][]int{}

	for _, line := range lines {
		items = append(items, toInts(strings.Split(line, " ")))
	}

	return items
}

func main() {
	input := readInput()
	histories := parseInput(input)
	sum := 0

	for _, history := range histories {
		sum += predictNextNumber(history)
	}

	fmt.Println(sum)
}
