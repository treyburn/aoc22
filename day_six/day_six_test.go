package day_six

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewScanner(t *testing.T) {
	wantCap := 2

	s := NewScanner(wantCap)

	assert.Equal(t, wantCap, s.cap)
	assert.Equal(t, 0, len(s.set))
	assert.Equal(t, 0, len(s.current))
}

func TestScanner_add(t *testing.T) {
	scan := NewScanner(2)

	scan.add(testRune("a"))
	assert.Equal(t, 1, len(scan.current))
	assert.Equal(t, 1, len(scan.set))
	assert.Equal(t, testRune("a"), scan.current[0])

	scan.add(testRune("b"))
	assert.Equal(t, 2, len(scan.current))
	assert.Equal(t, 2, len(scan.set))
	assert.Equal(t, testRune("a"), scan.current[0])
	assert.Equal(t, testRune("b"), scan.current[1])

	scan.add(testRune("c"))
	assert.Equal(t, 2, len(scan.current))
	assert.Equal(t, 2, len(scan.set))
	assert.Equal(t, testRune("b"), scan.current[0])
	assert.Equal(t, testRune("c"), scan.current[1])
}

func TestScanner_isUnique(t *testing.T) {
	scan := NewScanner(2)

	got := scan.isUnique(testRune("a"))
	assert.True(t, got)

	scan.add(testRune("a"))

	got = scan.isUnique(testRune("a"))
	assert.False(t, got)

	got = scan.isUnique(testRune("b"))
	assert.True(t, got)
}

func TestScanner_dropLeft(t *testing.T) {
	scan := NewScanner(2)
	scan.add(testRune("a"))
	scan.add(testRune("b"))

	got := scan.isUnique(testRune("a"))
	require.False(t, got)
	got = scan.isUnique(testRune("b"))
	require.False(t, got)

	scan.dropLeft()
	got = scan.isUnique(testRune("a"))
	assert.True(t, got)
	got = scan.isUnique(testRune("b"))
	require.False(t, got)

	assert.Equal(t, 1, len(scan.current))
	assert.Equal(t, 1, len(scan.set))
}

func TestScanner_len(t *testing.T) {
	scan := NewScanner(2)
	assert.Equal(t, 0, scan.len())

	scan.add(testRune("a"))
	assert.Equal(t, 1, scan.len())

	scan.add(testRune("b"))
	assert.Equal(t, 2, scan.len())

	scan.add(testRune("c"))
	assert.Equal(t, 2, scan.len())
}

func TestScanner_FindMarker(t *testing.T) {
	type testCase struct {
		name  string
		input string
		cap   int
		want  int
	}

	var tests = []testCase{
		{"first", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7},
		{"second", "bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5},
		{"third", "nppdvjthqldpwncqszvftbrmjlhg", 4, 6},
		{"fourth", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10},
		{"fifth", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11},
		{"first", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19},
		{"second", "bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23},
		{"third", "nppdvjthqldpwncqszvftbrmjlhg", 14, 23},
		{"fourth", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29},
		{"fifth", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			input := []rune(tc.input)

			scan := NewScanner(tc.cap)

			got := scan.FindMarker(input)

			assert.Equal(t, tc.want, got)
		})
	}
}

func testRune(r string) rune {
	return []rune(r)[0]
}
