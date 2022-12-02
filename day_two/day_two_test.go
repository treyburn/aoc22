package day_two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildStrategy_Initial(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  []MoveSet
	}

	var tests = []testCase{
		{"standard",
			"A Y\nB X\nC Z",
			[]MoveSet{
				{rock, paper},
				{paper, rock},
				{scissors, scissors}}},
		{"empty", "", []MoveSet{}},
		{"bad values", "T K\nI S", []MoveSet{}},
		{"mixed values", "A Z\nI S\nT K", []MoveSet{{rock, scissors}}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := BuildMoves(tc.input, InitialStrategy)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBuildStrategy_Fixed(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  []MoveSet
	}

	var tests = []testCase{
		{"standard",
			"A Y\nB X\nC Z",
			[]MoveSet{
				{rock, rock},
				{paper, rock},
				{scissors, rock}}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := BuildMoves(tc.input, FixedStrategy)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDetermineWinner(t *testing.T) {
	type testCase struct {
		name  string
		input MoveSet
		want  Outcome
	}

	var tests = []testCase{
		{"you win", MoveSet{rock, paper}, win},
		{"you lose", MoveSet{rock, scissors}, lost},
		{"you draw", MoveSet{rock, rock}, draw},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := determineYourOutcome(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestScore(t *testing.T) {
	type testCase struct {
		name  string
		input []MoveSet
		want  GameScore
	}

	var tests = []testCase{
		{"standard", []MoveSet{
			{rock, paper},
			{paper, rock},
			{scissors, scissors}}, GameScore{1 + 0 + 2 + 6 + 3 + 3, 15}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := Score(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestScore_FixedStrategy(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  GameScore
	}

	var tests = []testCase{
		{"standard", "A Y\nB X\nC Z", GameScore{
			OpponentScore: 4 + 2 + 6 + 3,
			YourScore:     12,
		}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			strat := BuildMoves(tc.input, FixedStrategy)

			got := Score(strat)

			assert.Equal(t, tc.want, got)
		})
	}
}
