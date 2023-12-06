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

func NewRace(input string) Race {
	input = regexp.MustCompile("([^0-9\n])").ReplaceAllString(input, "")
	lines := strings.Split(input, "\n")
	return Race{time: toInt(lines[0]), recordDistance: toInt(lines[1])}
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

func main() {
	race := NewRace(readInput())
	fmt.Println(race.countImprovements())
}
