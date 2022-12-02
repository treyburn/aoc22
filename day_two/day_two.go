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

type MoveSet struct {
	OpponentsMove Move
	YourMove      Move
}

type Strategy func(opponent, you string) (MoveSet, error)

type GameScore struct {
	OpponentScore int
	YourScore     int
}

// BuildMoves takes a newline separated string of space separated instructions and builds a MoveSet based on the provided Strategy
func BuildMoves(input string, build Strategy) []MoveSet {
	moves := make([]MoveSet, 0)
	lines := strings.Split(input, "\n")
	for idx, pair := range lines {
		if pair != "" {
			ops := strings.Split(pair, " ")
			if len(ops) == 2 {
				strat, err := build(ops[0], ops[1])
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

// InitialStrategy builds a MoveSet based on the initially assumed code
func InitialStrategy(opponent, you string) (MoveSet, error) {
	oOp, ok := initialDecode[strings.ToUpper(opponent)]
	if !ok {
		return MoveSet{}, fmt.Errorf("unexpected value: %v", opponent)
	}
	yOp, ok := initialDecode[strings.ToUpper(you)]
	if !ok {
		return MoveSet{}, fmt.Errorf("unexpected value: %v", you)
	}

	return MoveSet{
		OpponentsMove: oOp,
		YourMove:      yOp,
	}, nil
}

// FixedStrategy builds a MoveSet where your move is determined by a desired Outcome from the opponents Move
func FixedStrategy(opponent, you string) (MoveSet, error) {
	oOp, ok := initialDecode[strings.ToUpper(opponent)]
	if !ok {
		return MoveSet{}, fmt.Errorf("unexpected value: %v", opponent)
	}
	outcome, ok := requiredOutcome[strings.ToUpper(you)]
	if !ok {
		return MoveSet{}, fmt.Errorf("unexpected value: %v", you)
	}

	yOp := fixOutcome(oOp, outcome)

	return MoveSet{OpponentsMove: oOp, YourMove: yOp}, nil
}

// Score gives a GameScore based on a series of MoveSets
func Score(moves []MoveSet) GameScore {
	score := GameScore{}

	for _, round := range moves {
		yourOutcome := determineYourOutcome(round)
		score.YourScore += int(yourOutcome) + int(round.YourMove)
		score.OpponentScore += int(opponentsOutcome[yourOutcome]) + int(round.OpponentsMove)
	}

	return score
}

// determineYourOutcome determines the win/draw/lost Outcome of a strategy from "your" perspective
func determineYourOutcome(s MoveSet) Outcome {
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

// fixOutcome gives "you" the appropriate Move based on your opponents Move and the required Outcome
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
