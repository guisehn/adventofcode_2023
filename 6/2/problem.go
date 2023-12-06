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
	race := Race{time: 38947970, recordDistance: 241154910741091}
	fmt.Println(race.countImprovements())
}
