package main

import "fmt"

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

func main() {
	races := []Race{
		{time: 38, recordDistance: 241},
		{time: 94, recordDistance: 1549},
		{time: 79, recordDistance: 1074},
		{time: 70, recordDistance: 1091},
	}

	result := 1
	for _, race := range races {
		result *= race.countImprovements()
	}
	fmt.Println(result)
}
