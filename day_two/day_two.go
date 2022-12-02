package day_two

import (
	"fmt"
	"strings"
)

type Move int

const (
	rock Move = iota + 1
	paper
	scissors
)

var initialDecode = map[string]Move{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var beats = map[Move]Move{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

var loses = map[Move]Move{
	rock:     paper,
	paper:    scissors,
	scissors: rock,
}

type Outcome int

const (
	lost Outcome = iota * 3
	draw
	win
)

var opponentsOutcome = map[Outcome]Outcome{
	win:  lost,
	draw: draw,
	lost: win,
}

var requiredOutcome = map[string]Outcome{
	"X": lost,
	"Y": draw,
	"Z": win,
}

type Strategy struct {
	OpponentsMove Move
	YourMove      Move
}

type Ruleset func(string, string) (Strategy, error)

type GameScore struct {
	OpponentScore int
	YourScore     int
}

func BuildStrategy(input string, rule Ruleset) []Strategy {
	moves := make([]Strategy, 0)
	lines := strings.Split(input, "\n")
	for idx, pair := range lines {
		if pair != "" {
			ops := strings.Split(pair, " ")
			if len(ops) == 2 {
				strat, err := rule(ops[0], ops[1])
				if err != nil {
					fmt.Println(fmt.Sprintf("Error on line %v with values %v: %v", idx, pair, err))
					continue
				}
				moves = append(moves, strat)

			} else {
				fmt.Println(fmt.Sprintf("Bad initialDecode on line %v: %v decoded to %v", idx, pair, ops))
				continue
			}
		}
	}

	return moves
}

func InitialStrategy(first, second string) (Strategy, error) {
	oOp, ok := initialDecode[strings.ToUpper(first)]
	if !ok {
		return Strategy{}, fmt.Errorf("unexpected value: %v", first)
	}
	yOp, ok := initialDecode[strings.ToUpper(second)]
	if !ok {
		return Strategy{}, fmt.Errorf("unexpected value: %v", second)
	}

	return Strategy{
		OpponentsMove: oOp,
		YourMove:      yOp,
	}, nil
}

func FixedStrategy(first, second string) (Strategy, error) {
	oOp, ok := initialDecode[strings.ToUpper(first)]
	if !ok {
		return Strategy{}, fmt.Errorf("unexpected value: %v", first)
	}
	outcome, ok := requiredOutcome[strings.ToUpper(second)]
	if !ok {
		return Strategy{}, fmt.Errorf("unexpected value: %v", second)
	}

	yOp := fixOutcome(oOp, outcome)

	return Strategy{OpponentsMove: oOp, YourMove: yOp}, nil
}

func Score(moves []Strategy) GameScore {
	score := GameScore{}

	for _, round := range moves {
		yourOutcome := determineYourOutcome(round)
		score.YourScore += int(yourOutcome) + int(round.YourMove)
		score.OpponentScore += int(opponentsOutcome[yourOutcome]) + int(round.OpponentsMove)
	}

	return score
}

// determineYourOutcome determines the win/draw/lost Outcome of a strategy from "your" perspective
func determineYourOutcome(s Strategy) Outcome {
	opponentWins := beats[s.OpponentsMove]

	switch s.YourMove {
	case opponentWins:
		return lost
	case s.OpponentsMove:
		return draw
	default:
		return win
	}
}

func fixOutcome(opponent Move, required Outcome) Move {
	switch required {
	case win:
		return loses[opponent]
	case lost:
		return beats[opponent]
	default:
		return opponent
	}
}