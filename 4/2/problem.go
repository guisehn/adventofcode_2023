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

func (card Card) CountMatchingNumbers() int {
	winningMap := toMap(card.winningNumbers)
	matchingNumbers := 0

	for _, number := range card.numbers {
		if winningMap[number] {
			matchingNumbers++
		}
	}

	return matchingNumbers
}

func readInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
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

func incrementCount(counts *map[int]int, key int, quantity int) {
	if val, ok := (*counts)[key]; ok {
		(*counts)[key] = val + quantity
	} else {
		(*counts)[key] = quantity
	}
}

func main() {
	input := readInput()
	lines := strings.Split(input, "\n")
	len := len(lines)
	cardCounts := make(map[int]int)

	for _, line := range lines {
		card := NewCard(line)
		incrementCount(&cardCounts, card.cardNumber, 1)

		// fmt.Println("Processing card", card.cardNumber, "for which I have", cardCounts[card.cardNumber], "copies")

		quantity := cardCounts[card.cardNumber]
		max := min(len, card.cardNumber+card.CountMatchingNumbers())
		for i := card.cardNumber + 1; i <= max; i++ {
			// fmt.Println("Creating", quantity, "copies of card", i)
			incrementCount(&cardCounts, i, quantity)
		}

		// fmt.Println("")
		// fmt.Println(cardCounts)
		// fmt.Println("===")
	}

	sum := 0
	for _, count := range cardCounts {
		sum += count
	}

	fmt.Println(sum)
}
