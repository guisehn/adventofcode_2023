package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ConvertMap struct {
	sourceStart      int
	destinationStart int
	mapRange         int
}

type Category struct {
	maps []ConvertMap
}

func (category Category) Convert(number int) int {
	for _, cmap := range category.maps {
		sourceStart := cmap.sourceStart
		sourceEnd := sourceStart + cmap.mapRange

		if number >= sourceStart && number <= sourceEnd {
			// fmt.Println("found", number, "is between", sourceStart, "and", sourceEnd)
			distance := number - sourceStart
			return cmap.destinationStart + distance
		}
	}

	return number
}

type Problem struct {
	seeds      []int
	categories []Category
}

func NewProblem(input string) Problem {
	parts := strings.Split(input, "\n\n")
	seeds := toInts(strings.Split(strings.Replace(parts[0], "seeds: ", "", 1), " "))

	categories := []Category{}
	for i := 1; i < len(parts); i++ {
		part := parts[i]
		lines := strings.Split(part, "\n")

		maps := []ConvertMap{}
		for j := 1; j < len(lines); j++ {
			parts := strings.Split(lines[j], " ")
			cmap := ConvertMap{
				destinationStart: toInt(parts[0]),
				sourceStart:      toInt(parts[1]),
				mapRange:         toInt(parts[2]),
			}
			maps = append(maps, cmap)
		}

		categories = append(categories, Category{maps: maps})
	}

	return Problem{seeds: seeds, categories: categories}
}

func (problem Problem) ConvertAllSeeds() []int {
	seeds := []int{}
	for _, seed := range problem.seeds {
		seeds = append(seeds, problem.ConvertSeed(seed))
	}
	return seeds
}

func (problem Problem) ConvertSeed(number int) int {
	for _, category := range problem.categories {
		number = category.Convert(number)
	}

	return number
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

func main() {
	input := readInput()
	problem := NewProblem(input)
	seeds := problem.ConvertAllSeeds()
	min := slices.Min(seeds)
	fmt.Println(min)
}
