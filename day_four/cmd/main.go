package main

import (
	"fmt"
	"github.com/treyburn/aoc-21/day_four"
)

func main() {
	pairs, err := day_four.BuildSectorPairs(day_four.InputData)
	if err != nil {
		panic(err)
	}

	overlaps := day_four.CountFullOverlaps(pairs)

	fmt.Println("Day 4: Part 1: ", overlaps)

	partialOverlaps := day_four.CountPartialOverlaps(pairs)

	fmt.Println("Day 4: Part 2: ", partialOverlaps)
}
