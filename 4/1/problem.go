package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	cardNumber     int
	winningNumbers []int
	numbers        []int
}

var cardRe = regexp.MustCompile("Card +([0-9]+): ([0-9 ]+) \\| ([0-9 ]+)")
var splitRe = regexp.MustCompile("\\S+")

func NewCard(str string) Card {
	cardMatch := cardRe.FindStringSubmatch(str)

	cardNumber := toInt(cardMatch[1])
	winningNumbers := toInts(splitRe.FindAllString(cardMatch[2], -1))
	numbers := toInts(splitRe.FindAllString(cardMatch[3], -1))

	return Card{cardNumber: cardNumber, winningNumbers: winningNumbers, numbers: numbers}
}

func (card Card) Points() int {
	winningMap := toMap(card.winningNumbers)
	points := 0

	for _, number := range card.numbers {
		if winningMap[number] {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	return points
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

func toInts(strings []string) []int {
	ints := []int{}
	for _, str := range strings {
		ints = append(ints, toInt(str))
	}
	return ints
}

func toMap[K comparable](values []K) map[K]bool {
	m := make(map[K]bool)
	for _, val := range values {
		m[val] = true
	}
	return m
}

func main() {
	input := readInput()
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		card := NewCard(line)
		sum += card.Points()
	}

	fmt.Println(sum)
}
