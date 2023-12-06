package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	min int
	max int
}

type ConversionMap struct {
	source Range
	dest   Range
}

func (m ConversionMap) convert(r Range) Range {
	dist1 := r.min - m.source.min
	dist2 := r.max - m.source.min
	return Range{min: m.dest.min + dist1, max: m.dest.min + dist2}
}

type Category struct {
	name string
	maps []ConversionMap
}

func NewCategory(input string) Category {
	lines := strings.Split(input, "\n")
	maps := []ConversionMap{}

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		parts := strings.Split(line, " ")

		m := ConversionMap{
			source: Range{min: toInt(parts[1]), max: toInt(parts[1]) + toInt(parts[2]) - 1},
			dest:   Range{min: toInt(parts[0]), max: toInt(parts[0]) + toInt(parts[2]) - 1},
		}

		maps = append(maps, m)
	}

	sort.Slice(maps, func(i, j int) bool {
		return maps[i].source.min < maps[j].source.max
	})

	return Category{name: lines[0], maps: maps}
}

func (category Category) generateNewRanges(r Range) []Range {
	newRanges := []Range{}

	if r.min < category.maps[0].source.min {
		newRanges = append(newRanges, Range{min: r.min, max: min(category.maps[0].source.min-1, r.max)})
	}

	for _, m := range category.maps {
		// min                                 max
		// 40 41 42 43 44 45 46 47 48 49 50 51 52
		//                               50 51 52 53 54 55 56
		//                               min                 max
		if r.min <= m.source.max && r.max >= m.source.min {
			overlappedRange := Range{min: max(m.source.min, r.min), max: min(m.source.max, r.max)}
			newRanges = append(newRanges, m.convert(overlappedRange))
		}
	}

	lastCat := category.maps[len(category.maps)-1]
	if r.max > lastCat.source.max {
		newRanges = append(newRanges, Range{min: max(lastCat.source.max+1, r.min), max: r.max})
	}

	return newRanges
}

type Problem struct {
	seedRanges []Range
	categories []Category
}

func NewProblem(input string) Problem {
	parts := strings.Split(input, "\n\n")

	seedsInput := toInts(strings.Split(strings.Replace(parts[0], "seeds: ", "", 1), " "))
	seedRanges := []Range{}
	for i := 0; i < len(seedsInput); i += 2 {
		seedRanges = append(seedRanges, Range{min: seedsInput[i], max: seedsInput[i] + seedsInput[i+1] - 1})
	}

	categories := []Category{}
	for i := 1; i < len(parts); i++ {
		categories = append(categories, NewCategory(parts[i]))
	}

	return Problem{
		seedRanges: seedRanges,
		categories: categories,
	}
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
	ranges := problem.seedRanges

	for _, cat := range problem.categories {
		// fmt.Println(cat.name)
		// fmt.Printf("%+v\n", cat.maps)

		newRanges := []Range{}
		for _, r := range ranges {
			result := cat.generateNewRanges(r)
			// fmt.Println("transforming ranges", r, "they became", result)
			newRanges = append(newRanges, result...)
		}
		ranges = newRanges

		// fmt.Println("merged:", ranges)
		// fmt.Println("")
	}

	min := -1
	for _, r := range ranges {
		if min == -1 || r.min < min {
			min = r.min
		}
	}

	fmt.Println(min)
}
