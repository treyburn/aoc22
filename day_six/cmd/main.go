package main

import (
	"fmt"

	"github.com/treyburn/aoc-21/day_six"
)

func main() {
	scan := day_six.NewScanner(4)
	got := scan.FindMarker(day_six.InputData)
	fmt.Println("Day 6: Part 1: ", got)

	scan2 := day_six.NewScanner(14)
	got2 := scan2.FindMarker(day_six.InputData)
	fmt.Println("Day 6: Part 2: ", got2)
}
