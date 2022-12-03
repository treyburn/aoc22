package day_three

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRuck(t *testing.T) {
	type testCase struct {
		name string
		input string
		wantLen int
	}

	var tests = []testCase{
		{"first", "vJrwpWtwJgWrhcsFMMfFFhFp", 12},
		{"second", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 16},
		{"third", "PmmdzqPrVvPwwTWBwg", 9},
		{"fourth", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 15},
		{"fifth", "ttgJtRGJQctTZtZT", 8},
		{"sixth", "CrZsJsPPZsGzwwsLwLmpwMDw", 12},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := NewRuck(tc.input)

			assert.Equal(t, tc.wantLen, got.itemsPerComp)
		})
	}
}

func TestRuck_DeDupe(t *testing.T) {
	type testCase struct {
		name string
		input string
		wantDupe rune
	}

	var tests = []testCase{
		{"first", "vJrwpWtwJgWrhcsFMMfFFhFp", testRune("p")},
		{"second", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", testRune("L")},
		{"third", "PmmdzqPrVvPwwTWBwg", testRune("P")},
		{"fourth", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", testRune("v")},
		{"fifth", "ttgJtRGJQctTZtZT", testRune("t")},
		{"sixth", "CrZsJsPPZsGzwwsLwLmpwMDw", testRune("s")},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			ruck := NewRuck(tc.input)

			got := ruck.DeDupe()

			assert.Equal(t, tc.wantDupe, got)

		})
	}
}

func TestRuck_IsDeDuped(t *testing.T) {
	type testCase struct {
		name string
		input string
	}

	var want = true
	var noDupe rune

	var tests = []testCase{
		{"first", "vJrwpWtwJgWrhcsFMMfFFhFp"},
		{"second", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
		{"third", "PmmdzqPrVvPwwTWBwg"},
		{"fourth", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
		{"fifth", "ttgJtRGJQctTZtZT"},
		{"sixth", "CrZsJsPPZsGzwwsLwLmpwMDw"},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			ruck := NewRuck(tc.input)

			_ = ruck.DeDupe()

			got := ruck.IsDeDuped()

			assert.Equal(t, want, got)

			gotDupe := ruck.DeDupe()

			assert.Equal(t, noDupe, gotDupe)

		})
	}
}

func TestBuildRucks(t *testing.T) {
	var testData = `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg

wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

	want := 6

	got := BuildRucks(testData)

	assert.Equal(t, want, len(got))
}

func TestGetPriority(t *testing.T) {
	type testCase struct {
		name string
		input rune
		want int
	}

	var tests = []testCase{
		{"first", testRune("p"), 16},
		{"second", testRune("L"), 38},
		{"third", testRune("P"), 42},
		{"fourth", testRune("v"), 22},
		{"fifth", testRune("t"), 20},
		{"sixth", testRune("s"), 19},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := GetPriority(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSumDupePriorities(t *testing.T) {
	var testData = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	want := 157

	rucks := BuildRucks(testData)

	got := SumDupePriorities(rucks)

	assert.Equal(t, want, got)
}

func TestBuildGroups(t *testing.T) {
	type testCase struct {
		name string
		input string
		want int
	}

	var tests = []testCase{
		{"one", `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg`, 1},
		{"two", `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`, 2},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			rucks := BuildRucks(tc.input)
			got, err := BuildGroups(rucks)

			assert.NoError(t, err)
			assert.Equal(t, tc.want, len(got))
		})
	}
}

func TestGetGroupBadge(t *testing.T) {
	type testCase struct {
		name string
		input string
		want rune
	}

	var tests = []testCase{
		{"one", `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg`, testRune("r")},
		{"two", `wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`, testRune("Z")},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			rucks := BuildRucks(tc.input)
			for _, ruck := range rucks {
				_ = ruck.DeDupe()
			}
			group, err := BuildGroups(rucks)

			require.NoError(t, err)
			require.Equal(t, 1, len(group))

			got, err := GetGroupBadge(group[0])
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSumBadgePriorities(t *testing.T) {
	type testCase struct {
		name string
		input string
		want int
	}

	var tests = []testCase{
		{"combo", `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`, 70},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			rucks := BuildRucks(tc.input)
			for _, ruck := range rucks {
				_ = ruck.DeDupe()
			}

			groups, err := BuildGroups(rucks)
			require.NoError(t, err)

			got, err := SumBadgePriorities(groups)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func testRune(s string) rune {
	return []rune(s)[0]
}