package day_five

import (
	"fmt"
	"regexp"
	"strings"
)

var endOfDiagramRe = regexp.MustCompile(`^((\s\d\s)\s?)+$`)

var numQueuesRe = regexp.MustCompile(`\d`)

var spaceRune = []rune(" ")[0]

func MakeQueues(n int) map[int]*Queue {
	q := make(map[int]*Queue)
	if n < 1 {
		return q
	}
	for i := 1; i <= n; i++ {
		q[i] = NewQueue()
	}

	return q
}

// isEndDiagram returns true when it encounters a string matching the endOfDiagramRe regex
//
//	and a count of the digits in the end of diagram string
func isEndDiagram(input string) (int, bool) {
	if endOfDiagramRe.MatchString(input) {
		return len(numQueuesRe.FindAllString(input, -1)), true
	}
	return 0, false
}

func chunkStringToRunePtrs(input string, numChunks int) []*rune {
	res := make([]*rune, 0)
	raw := []rune(input)

	count := 0
	for i := 1; count < numChunks; i += 4 {
		r := raw[i]
		if r == spaceRune {
			res = append(res, nil)
		} else {
			res = append(res, &r)
		}
		count++
	}

	return res
}

func MakeInitialQueues(input []string, length int) (map[int]*Queue, error) {
	queues := MakeQueues(length)

	for _, row := range input {
		chunks := chunkStringToRunePtrs(row, length)
		for idx, val := range chunks {
			if val == nil {
				continue
			}
			n := &Node{value: *val}
			q, ok := queues[idx+1]
			if !ok {
				return nil, fmt.Errorf("invalid input: have %v queues - but tried to add to #%v - on input line '%v'", length, idx+1, row)
			}
			q.AddToBack(n)
		}
	}

	return queues, nil
}

type Move struct {
	count int
	from  int
	to    int
}

func parseMoves(input string) (Move, error) {
	var count, from, to int
	n, err := fmt.Sscanf(input, "move %d from %d to %d", &count, &from, &to)
	if err != nil {
		return Move{}, err
	}
	if n != 3 {
		return Move{}, fmt.Errorf("invalid move parse - got %v from '%v'", n, input)
	}

	return Move{
		count: count,
		from:  from,
		to:    to,
	}, nil
}

func ProcessMoves(queues map[int]*Queue, move Move) error {
	from, ok := queues[move.from]
	if !ok {
		return fmt.Errorf("invalid move: no queue for %v", move.from)
	}

	to, ok := queues[move.to]
	if !ok {
		return fmt.Errorf("invalid move: no queue for %v", move.from)
	}

	items := from.Dequeue(move.count)
	to.ReverseEnqueue(items)

	return nil
}

func ProcessMovesUpdated(queues map[int]*Queue, move Move) error {
	from, ok := queues[move.from]
	if !ok {
		return fmt.Errorf("invalid move: no queue for %v", move.from)
	}

	to, ok := queues[move.to]
	if !ok {
		return fmt.Errorf("invalid move: no queue for %v", move.from)
	}

	items := from.Dequeue(move.count)
	to.Enqueue(items)

	return nil
}

type Processor func(map[int]*Queue, Move) error

func RunCrane(input string, proc Processor) ([]string, error) {
	split := strings.Split(input, "\n")
	initalSetup := make([]string, 0)
	var doneCollecting bool
	var queues map[int]*Queue
	var move Move
	var err error
	var totalQueues int

	for _, row := range split {
		// start by building the initial queue setup
		if !doneCollecting {
			count, ok := isEndDiagram(row)
			if ok {
				totalQueues = count
				doneCollecting = true
				queues, err = MakeInitialQueues(initalSetup, count)
				if err != nil {
					return nil, err
				}
			} else {
				initalSetup = append(initalSetup, row)
			}
		} else {
			if row == "" {
				continue
			}
			move, err = parseMoves(row)
			if err != nil {
				return nil, err
			}

			err = proc(queues, move)
			if err != nil {
				return nil, err
			}
		}
	}

	result := make([]string, 0)
	for i := 1; i <= totalQueues; i++ {
		result = append(result, queues[i].Dequeue(1).Value())
	}

	return result, nil
}
