package day_five

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeQueues(t *testing.T) {
	type testCase struct {
		name    string
		input   int
		wantLen int
	}

	var tests = []testCase{
		{"none", 0, 0},
		{"one", 1, 1},
		{"five", 5, 5},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := MakeQueues(tc.input)

			assert.Equal(t, tc.wantLen, len(got))

			for i := 1; i <= tc.wantLen; i++ {
				q, ok := got[i]
				assert.True(t, ok)
				assert.NotNil(t, q)
			}
		})
	}
}

func TestEndOfDiagram(t *testing.T) {
	type testCase struct {
		name      string
		input     string
		wantCount int
		wantBool  bool
	}

	var tests = []testCase{
		{"invalid 1", "abc", 0, false},
		{"invalid 2", " 1 2 ", 0, false},
		{"valid 1", " 1 ", 1, true},
		{"valid 2", " 1   2 ", 2, true},
		{"valid 3", " 1   2   3 ", 3, true},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			count, ok := isEndDiagram(tc.input)

			assert.Equal(t, tc.wantCount, count)
			assert.Equal(t, tc.wantBool, ok)
		})
	}
}

func TestChunkStringToRunsPtrs(t *testing.T) {
	a := []rune("a")[0]
	b := []rune("b")[0]
	c := []rune("c")[0]

	type testCase struct {
		name   string
		input  string
		chunks int
		want   []*rune
	}

	var tests = []testCase{
		{"invalid", "", 0, []*rune{}},
		{"one rune", " a ", 1, []*rune{&a}},
		{"two rune", " a   b ", 2, []*rune{&a, &b}},
		{"with a nil", "     b ", 2, []*rune{nil, &b}},
		{"with a nil", "     b   c ", 3, []*rune{nil, &b, &c}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := chunkStringToRunePtrs(tc.input, tc.chunks)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMakeInitialQueues(t *testing.T) {
	rawInput := `    [D]    
[N] [C]    
[Z] [M] [P]`

	testData := strings.Split(rawInput, "\n")

	wantLen := 3

	queues, err := MakeInitialQueues(testData, wantLen)
	assert.NoError(t, err)

	assert.Equal(t, wantLen, len(queues))

	one, ok := queues[1]
	assert.True(t, ok)
	item := one.Dequeue(1)
	assert.Equal(t, "N", item.Value())
	item = one.Dequeue(1)
	assert.Equal(t, "Z", item.Value())
	item = one.Dequeue(1)
	assert.Nil(t, item)

	two, ok := queues[2]
	assert.True(t, ok)
	item = two.Dequeue(1)
	assert.Equal(t, "D", item.Value())
	item = two.Dequeue(1)
	assert.Equal(t, "C", item.Value())
	item = two.Dequeue(1)
	assert.Equal(t, "M", item.Value())
	item = two.Dequeue(1)
	assert.Nil(t, item)

	three, ok := queues[3]
	assert.True(t, ok)
	item = three.Dequeue(1)
	assert.Equal(t, "P", item.Value())
	item = three.Dequeue(1)
	assert.Nil(t, item)
}

func TestParseMoves(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  Move
	}

	var tests = []testCase{
		{"first", "move 1 from 2 to 1", Move{1, 2, 1}},
		{"second", "move 3 from 1 to 3", Move{3, 1, 3}},
		{"third", "move 2 from 2 to 1", Move{2, 2, 1}},
		{"fourth", "move 1 from 1 to 2", Move{1, 1, 2}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseMoves(tc.input)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)

		})
	}
}

func TestProcessMoves(t *testing.T) {
	rawInput := `    [D]    
[N] [C]    
[Z] [M] [P]`

	testData := strings.Split(rawInput, "\n")
	dataLen := 3

	queues, err := MakeInitialQueues(testData, dataLen)
	require.NoError(t, err)

	move, err := parseMoves("move 1 from 2 to 1")
	require.NoError(t, err)

	err = ProcessMoves(queues, move)
	assert.NoError(t, err)

	assert.Equal(t, "D", queues[1].Dequeue(1).Value())
	assert.Equal(t, "C", queues[2].Dequeue(1).Value())
	assert.Equal(t, "P", queues[3].Dequeue(1).Value())
}

func TestProcessMoves_Multiple(t *testing.T) {
	rawInput := `    [D]    
[N] [C]    
[Z] [M] [P]`

	testData := strings.Split(rawInput, "\n")
	dataLen := 3

	queues, err := MakeInitialQueues(testData, dataLen)
	require.NoError(t, err)

	move, err := parseMoves("move 1 from 2 to 1")
	require.NoError(t, err)
	err = ProcessMoves(queues, move)
	require.NoError(t, err)

	move, err = parseMoves("move 3 from 1 to 3")
	require.NoError(t, err)
	err = ProcessMoves(queues, move)
	require.NoError(t, err)

	assert.Nil(t, queues[1].Dequeue(1))
	assert.Equal(t, "C", queues[2].Dequeue(1).Value())
	assert.Equal(t, "Z", queues[3].Dequeue(1).Value())
}

func TestRunCrane(t *testing.T) {
	rawInput := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	want := []string{"C", "M", "Z"}

	got, err := RunCrane(rawInput, ProcessMoves)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestProcessMovesUpdated(t *testing.T) {
	rawInput := `    [D]    
[N] [C]    
[Z] [M] [P]`

	testData := strings.Split(rawInput, "\n")
	dataLen := 3

	queues, err := MakeInitialQueues(testData, dataLen)
	require.NoError(t, err)

	move, err := parseMoves("move 1 from 2 to 1")
	require.NoError(t, err)
	err = ProcessMovesUpdated(queues, move)
	require.NoError(t, err)

	move, err = parseMoves("move 3 from 1 to 3")
	require.NoError(t, err)
	err = ProcessMovesUpdated(queues, move)
	require.NoError(t, err)

	assert.Nil(t, queues[1].Dequeue(1))
	assert.Equal(t, "C", queues[2].Dequeue(1).Value())
	assert.Equal(t, "D", queues[3].Dequeue(1).Value())
}

func TestRunCrane_Updated(t *testing.T) {
	rawInput := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	want := []string{"M", "C", "D"}

	got, err := RunCrane(rawInput, ProcessMovesUpdated)
	assert.NoError(t, err)

	assert.Equal(t, want, got)
}
