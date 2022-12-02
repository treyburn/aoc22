package main

import (
	"fmt"

	"github.com/treyburn/aoc-21/day_two"
)

func main() {
	game := day_two.BuildStrategy(day_two.InputData, day_two.InitialStrategy)

	score := day_two.Score(game)

	fmt.Println("Your score: ", score.YourScore)
	fmt.Println("Opponents score: ", score.OpponentScore)

	fixedGame := day_two.BuildStrategy(day_two.InputData, day_two.FixedStrategy)

	fixedScore := day_two.Score(fixedGame)

	fmt.Println("Your score - fixed: ", fixedScore.YourScore)
	fmt.Println("Opponents score - fixed: ", fixedScore.OpponentScore)
}
