package main

import (
	"fmt"

	"github.com/treyburn/aoc-21/day_one"
)

func main() {
	calories := day_one.CountCalories(day_one.InputData)

	highestCal := day_one.FindHighestCalorie(calories)

	fmt.Println("Day 1: Part 1: ", highestCal)

	sumCals := day_one.SumTop3Calories(calories)

	fmt.Println("Day 1: Part 2: ", sumCals)

}
