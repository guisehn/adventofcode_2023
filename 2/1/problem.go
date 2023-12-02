package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type GameSet struct {
	red   int
	green int
	blue  int
}

func NewGameSet(str string) GameSet {
	set := GameSet{}
	cubes := strings.Split(str, ", ")

	for _, cube := range cubes {
		parts := strings.Split(cube, " ")
		num := toInt(parts[0])

		switch color := parts[1]; color {
		case "red":
			set.red = num
		case "green":
			set.green = num
		case "blue":
			set.blue = num
		}
	}

	return set
}

func (set GameSet) IsValid() bool {
	return set.red <= 12 && set.green <= 13 && set.blue <= 14
}

type Game struct {
	number int
	sets   []GameSet
}

var gameRe = regexp.MustCompile("Game ([0-9]+): (.+)")

func NewGame(str string) Game {
	gameMatch := gameRe.FindStringSubmatch(str)

	number := toInt(gameMatch[1])
	setStrings := strings.Split(gameMatch[2], "; ")
	sets := []GameSet{}

	for _, setStr := range setStrings {
		sets = append(sets, NewGameSet(setStr))
	}

	return Game{number: number, sets: sets}
}

func (game Game) IsValid() bool {
	for _, set := range game.sets {
		if !set.IsValid() {
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

func readInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}

func main() {
	input := readInput()
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		game := NewGame(line)
		if game.IsValid() {
			sum += game.number
		}
	}

	fmt.Println(sum)
}
