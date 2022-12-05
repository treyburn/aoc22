package main

import (
	"fmt"

	"github.com/treyburn/aoc-21/day_five"
)

func main() {
	result, err := day_five.RunCrane(day_five.InputData, day_five.ProcessMoves)
	if err != nil {
		panic(err)
	}

	fmt.Println("day 5: Part 1: ", result)

	result2, err := day_five.RunCrane(day_five.InputData, day_five.ProcessMovesUpdated)
	if err != nil {
		panic(err)
	}

	fmt.Println("day 5: Part 2: ", result2)
}
