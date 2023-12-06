package main

import (
	"fmt"
	"sort"
)

type Range struct {
	min int
	max int
}

type ConversionMap struct {
	source Range
	dest   Range
}

type Category struct {
	maps []ConversionMap
}

func convert(r Range, m ConversionMap) Range {
	dist1 := r.min - m.source.min
	dist2 := r.max - m.source.min
	return Range{min: m.dest.min + dist1, max: m.dest.min + dist2}
}

func generateNewRanges(r Range, category Category) []Range {
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
			overlap := Range{min: max(m.source.min, r.min), max: min(m.source.max, r.max)}
			newRanges = append(newRanges, convert(overlap, m))
		}
	}

	lastCat := category.maps[len(category.maps)-1]
	if r.max > lastCat.source.max {
		newRanges = append(newRanges, Range{min: lastCat.source.max + 1, max: r.max})
	}

	return newRanges
}

func main() {
	r := Range{min: 40, max: 105}

	maps := []ConversionMap{
		{source: Range{min: 98, max: 99}, dest: Range{min: 50, max: 51}},
		{source: Range{min: 50, max: 97}, dest: Range{min: 52, max: 99}},
	}

	sort.Slice(maps, func(i, j int) bool {
		return maps[i].source.min < maps[j].source.max
	})

	cat := Category{maps: maps}

	x := generateNewRanges(r, cat)

	fmt.Println(x)
}
