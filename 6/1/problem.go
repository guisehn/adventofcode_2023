package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time           int
	recordDistance int
}

func (race Race) countImprovements() int {
	count := 0

	for holdTime := 0; holdTime <= race.time; holdTime++ {
		speed := holdTime
		distance := (race.time - holdTime) * speed
		// fmt.Println("hold time", holdTime, "distance", distance)
		if distance > race.recordDistance {
			count++
		}
	}

	return count
}

func readInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func buildRaces(input string) []Race {
	input = regexp.MustCompile("([^0-9 \n])").ReplaceAllString(input, "")
	input = regexp.MustCompile("[ ]+").ReplaceAllString(input, " ")

	lines := strings.Split(input, "\n")
	times := strings.Split(strings.TrimSpace(lines[0]), " ")
	distances := strings.Split(strings.TrimSpace(lines[1]), " ")

	races := []Race{}
	for i, time := range times {
		races = append(races, Race{time: toInt(time), recordDistance: toInt(distances[i])})
	}
	return races
}

func main() {
	input := readInput()
	races := buildRaces(input)

	result := 1
	for _, race := range races {
		result *= race.countImprovements()
	}
	fmt.Println(result)
}
