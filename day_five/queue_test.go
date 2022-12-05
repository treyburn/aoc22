package day_five

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_Find(t *testing.T) {
	type testCase struct {
		name     string
		input    *Node
		position int
		want     *Node
	}

	var tests = []testCase{
		{"self", &Node{value: []rune("X")[0]}, 0, &Node{value: []rune("X")[0]}},
		{"too long", &Node{value: []rune("X")[0]}, 1, nil},
		{"last", &Node{value: []rune("X")[0], next: &Node{value: []rune("Y")[0], next: &Node{value: []rune("Z")[0]}}}, 2, &Node{value: []rune("Z")[0]}},
		{"mid", &Node{value: []rune("X")[0], next: &Node{value: []rune("Y")[0], next: &Node{value: []rune("Z")[0]}}}, 1, &Node{value: []rune("Y")[0], next: &Node{value: []rune("Z")[0]}}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.Find(tc.position)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNode_PopNext(t *testing.T) {
	type testCase struct {
		name  string
		input *Node
		want  *Node
	}

	var tests = []testCase{
		{"no next", &Node{}, nil},
		{"has next", &Node{value: []rune("X")[0], next: &Node{value: []rune("Y")[0], prev: &Node{value: []rune("X")[0]}}}, &Node{value: []rune("Y")[0]}},
		{"has next with next", &Node{value: []rune("X")[0], next: &Node{value: []rune("Y")[0], prev: &Node{value: []rune("X")[0]}, next: &Node{value: []rune("Z")[0]}}}, &Node{value: []rune("Y")[0], next: &Node{value: []rune("Z")[0]}}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.PopNext()

			assert.Equal(t, tc.want, got)
			assert.Nil(t, tc.input.next)

			if tc.want != nil {
				assert.Nil(t, got.prev)
			}
		})
	}
}

func TestNode_Last(t *testing.T) {
	type testCase struct {
		name  string
		input *Node
		want  *Node
	}

	var tests = []testCase{
		{"self", &Node{value: []rune("X")[0]}, &Node{value: []rune("X")[0]}},
		{"length of 2", &Node{value: []rune("X")[0], next: &Node{value: []rune("Y")[0]}}, &Node{value: []rune("Y")[0]}},
		{"length of 3", &Node{value: []rune("X")[0], next: &Node{value: []rune("Y")[0], next: &Node{value: []rune("Z")[0]}}}, &Node{value: []rune("Z")[0]}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.Last()

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewQueue(t *testing.T) {
	q := NewQueue()

	assert.Equal(t, q.tail, q.head.next)
	assert.Equal(t, q.head, q.tail.prev)
	assert.Nil(t, q.head.prev)
	assert.Nil(t, q.tail.next)
}

func TestQueue_Enqueue_Singles(t *testing.T) {

	x := &Node{value: []rune("X")[0]}
	y := &Node{value: []rune("Y")[0]}
	z := &Node{value: []rune("Z")[0]}

	q := NewQueue()

	q.Enqueue(z)
	assert.Equal(t, z, q.head.next)
	assert.Equal(t, q.head, z.prev)
	assert.Equal(t, q.tail.prev, z)
	assert.Equal(t, q.tail, z.next)

	q.Enqueue(y)
	assert.Equal(t, q.head.next, y)
	assert.Equal(t, q.head, y.prev)
	assert.Equal(t, z, y.next)
	assert.Equal(t, y, z.prev)
	assert.Equal(t, q.tail.prev, z)
	assert.Equal(t, q.tail, z.next)

	q.Enqueue(x)
	assert.Equal(t, q.head.next, x)
	assert.Equal(t, q.head, x.prev)
	assert.Equal(t, y, x.next)
	assert.Equal(t, x, y.prev)
	assert.Equal(t, z, y.next)
	assert.Equal(t, y, z.prev)
	assert.Equal(t, q.tail.prev, z)
	assert.Equal(t, q.tail, z.next)
}

func TestQueue_Enqueue_MultipleLinked(t *testing.T) {
	x := &Node{value: []rune("X")[0]}
	y := &Node{value: []rune("Y")[0]}
	z := &Node{value: []rune("Z")[0]}

	x.next = y
	y.prev = x
	y.next = z
	z.prev = y

	q := NewQueue()

	q.Enqueue(x)
	assert.Equal(t, x, q.head.next)
	assert.Equal(t, q.head, x.prev)
	assert.Equal(t, y, x.next)
	assert.Equal(t, x, y.prev)
	assert.Equal(t, z, y.next)
	assert.Equal(t, y, z.prev)
	assert.Equal(t, q.tail, z.next)
	assert.Equal(t, z, q.tail.prev)
}

func TestQueue_Dequeue_Single(t *testing.T) {
	want := &Node{value: []rune("X")[0]}

	q := NewQueue()
	q.Enqueue(want)

	got := q.Dequeue(1)

	assert.Equal(t, want, got)
	assert.Nil(t, got.prev)
	assert.Nil(t, got.next)
	assert.Equal(t, q.head, q.tail.prev)
	assert.Equal(t, q.tail, q.head.next)
}

func TestQueue_Dequeue_ZeroGetsNilNoChanges(t *testing.T) {
	x := &Node{value: []rune("X")[0]}

	q := NewQueue()
	q.Enqueue(x)

	got := q.Dequeue(0)

	assert.Nil(t, got)
	assert.Equal(t, x, q.tail.prev)
	assert.Equal(t, x, q.head.next)
	assert.Equal(t, q.head, x.prev)
	assert.Equal(t, q.tail, x.next)
}

func TestQueue_Dequeue_Single_WithMultipleInQueue(t *testing.T) {
	x := &Node{value: []rune("X")[0]}
	y := &Node{value: []rune("Y")[0]}
	x.next = y
	y.prev = x

	q := NewQueue()
	q.Enqueue(x)

	got := q.Dequeue(1)

	assert.Equal(t, x, got)
	assert.Nil(t, got.prev)
	assert.Nil(t, got.next)
	assert.Equal(t, y, q.tail.prev)
	assert.Equal(t, q.tail, y.next)
	assert.Equal(t, q.head, y.prev)
	assert.Equal(t, y, q.head.next)
}

func TestQueue_Dequeue_Multiple(t *testing.T) {
	x := &Node{value: []rune("X")[0]}
	y := &Node{value: []rune("Y")[0]}
	x.next = y
	y.prev = x

	q := NewQueue()
	q.Enqueue(x)

	got := q.Dequeue(2)

	assert.Equal(t, x, got)
	assert.Nil(t, x.prev)
	assert.Equal(t, y, x.next)
	assert.Nil(t, y.next)
	assert.Equal(t, x, y.prev)
	assert.Equal(t, q.head, q.tail.prev)
	assert.Equal(t, q.tail, q.head.next)
}

func TestQueue_Dequeue_Multiple_WithRemainder(t *testing.T) {
	x := &Node{value: []rune("X")[0]}
	y := &Node{value: []rune("Y")[0]}
	z := &Node{value: []rune("Z")[0]}
	x.next = y
	y.prev = x
	y.next = z
	z.prev = y

	q := NewQueue()
	q.Enqueue(x)

	got := q.Dequeue(2)

	assert.Equal(t, x, got)
	assert.Nil(t, x.prev)
	assert.Equal(t, y, x.next)
	assert.Nil(t, y.next)
	assert.Equal(t, x, y.prev)
	assert.Equal(t, z, q.tail.prev)
	assert.Equal(t, q.tail, z.next)
	assert.Equal(t, z, q.head.next)
	assert.Equal(t, q.head, z.prev)
}

func TestQueue_AddToBack_Single(t *testing.T) {
	x := &Node{value: []rune("X")[0]}

	q := NewQueue()
	q.AddToBack(x)

	assert.Equal(t, x, q.head.next)
	assert.Equal(t, q.head, x.prev)
	assert.Equal(t, q.tail, x.next)
	assert.Equal(t, x, q.tail.prev)
}

func TestQueue_AddToBack_Single_WithExisting(t *testing.T) {
	x := &Node{value: []rune("X")[0]}
	y := &Node{value: []rune("Y")[0]}

	q := NewQueue()
	q.Enqueue(x)
	q.AddToBack(y)

	assert.Equal(t, x, q.head.next)
	assert.Equal(t, q.head, x.prev)
	assert.Equal(t, y, x.next)
	assert.Equal(t, q.tail, y.next)
	assert.Equal(t, x, y.prev)
	assert.Equal(t, y, q.tail.prev)
}

func TestQueue_AddToBack_Multiple(t *testing.T) {
	x := &Node{value: []rune("X")[0]}
	y := &Node{value: []rune("Y")[0]}
	z := &Node{value: []rune("Z")[0]}
	x.next = y
	y.prev = x
	y.next = z
	z.prev = y

	q := NewQueue()
	q.AddToBack(x)

	assert.Equal(t, x, q.head.next)
	assert.Equal(t, q.head, x.prev)
	assert.Equal(t, y, x.next)
	assert.Equal(t, z, y.next)
	assert.Equal(t, x, y.prev)
	assert.Equal(t, z, q.tail.prev)
}
