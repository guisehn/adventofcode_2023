package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

var cardScores = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"1": 1,
	"J": 0,
}

type HandType = int

const (
	FiveOfAKind  HandType = 7
	FourOfAKind  HandType = 6
	FullHouse    HandType = 5
	ThreeOfAKind HandType = 4
	TwoPair      HandType = 3
	OnePair      HandType = 2
	HighCard     HandType = 1
)

type Hand struct {
	cards []string
	bid   int
}

func NewHand(input string) Hand {
	parts := strings.Split(input, " ")

	return Hand{
		cards: strings.Split(parts[0], ""),
		bid:   toInt(parts[1]),
	}
}

func (hand Hand) JokerizedCards() []string {
	mostCommonCard := hand.MostCommonCard()
	newCards := []string{}

	for _, card := range hand.cards {
		if card == "J" {
			newCards = append(newCards, mostCommonCard)
		} else {
			newCards = append(newCards, card)
		}
	}
	return newCards
}

func (hand Hand) MostCommonCard() string {
	counts := countCards(hand.cards)
	delete(counts, "J")

	max, mostCommonCard := 0, ""
	for card, count := range counts {
		if count > max {
			max = count
			mostCommonCard = card
		}
	}

	return mostCommonCard
}

func (hand Hand) HandType() HandType {
	counts := countCards(hand.JokerizedCards())

	if len(counts) == 1 {
		return FiveOfAKind
	}

	values := maps.Values(counts)
	sort.Ints(values)

	if values[0] == 1 && values[1] == 4 {
		return FourOfAKind
	}
	if values[0] == 2 && values[1] == 3 {
		return FullHouse
	}
	if values[0] == 1 && values[1] == 1 && values[2] == 3 {
		return ThreeOfAKind
	}
	if values[0] == 1 && values[1] == 2 && values[2] == 2 {
		return TwoPair
	}
	if values[0] == 1 && values[1] == 1 && values[2] == 1 && values[3] == 2 {
		return OnePair
	}

	return HighCard
}

type ByTypeAndScore []Hand

func (list ByTypeAndScore) Len() int      { return len(list) }
func (list ByTypeAndScore) Swap(i, j int) { list[i], list[j] = list[j], list[i] }
func (list ByTypeAndScore) Less(i, j int) bool {
	a, b := list[i], list[j]
	aType, bType := a.HandType(), b.HandType()

	if aType != bType {
		return aType < bType
	}

	for k := 0; k < 5; k++ {
		if a.cards[k] != b.cards[k] {
			return cardScores[a.cards[k]] < cardScores[b.cards[k]]
		}
	}

	return true
}

func countCards(cards []string) map[string]int {
	counts := make(map[string]int)

	for _, card := range cards {
		value, ok := counts[card]
		if ok {
			counts[card] = value + 1
		} else {
			counts[card] = 1
		}
	}

	return counts
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
	lines := strings.Split(readInput(), "\n")
	hands := []Hand{}

	for _, line := range lines {
		hands = append(hands, NewHand(line))
	}

	sort.Sort(ByTypeAndScore(hands))

	result := 0
	for i, hand := range hands {
		rank := i + 1
		result += rank * hand.bid
	}

	fmt.Println(result)
}
