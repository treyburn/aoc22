package main

import (
	"fmt"
	"github.com/treyburn/aoc-21/day_three"
)

func main() {
	rucks := day_three.BuildRucks(day_three.InputData)

	sum := day_three.SumDupePriorities(rucks)

	fmt.Println("Day 3: Part 1: ", sum)

	groups, err := day_three.BuildGroups(rucks)
	if err != nil {
		panic(err)
	}

	badgeSum, err := day_three.SumBadgePriorities(groups)
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 3: Part 2: ", badgeSum)
}
