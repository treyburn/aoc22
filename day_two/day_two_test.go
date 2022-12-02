package day_two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildStrategy_Initial(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  []Strategy
	}

	var tests = []testCase{
		{"standard",
			"A Y\nB X\nC Z",
			[]Strategy{
				{rock, paper},
				{paper, rock},
				{scissors, scissors}}},
		{"empty", "", []Strategy{}},
		{"bad values", "T K\nI S", []Strategy{}},
		{"mixed values", "A Z\nI S\nT K", []Strategy{{rock, scissors}}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := BuildStrategy(tc.input, InitialRules)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBuildStrategy_Fixed(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  []Strategy
	}

	var tests = []testCase{
		{"standard",
			"A Y\nB X\nC Z",
			[]Strategy{
				{rock, rock},
				{paper, rock},
				{scissors, rock}}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := BuildStrategy(tc.input, FixedRules)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDetermineWinner(t *testing.T) {
	type testCase struct {
		name  string
		input Strategy
		want  Outcome
	}

	var tests = []testCase{
		{"you win", Strategy{rock, paper}, win},
		{"you lose", Strategy{rock, scissors}, lost},
		{"you draw", Strategy{rock, rock}, draw},
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
		input []Strategy
		want  GameScore
	}

	var tests = []testCase{
		{"standard", []Strategy{
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
			strat := BuildStrategy(tc.input, FixedRules)

			got := Score(strat)

			assert.Equal(t, tc.want, got)
		})
	}
}
