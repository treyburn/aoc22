package day_five

// Node is a node in a linked list
type Node struct {
	value rune
	prev  *Node
	next  *Node
}

// Find traverses the linked list to find the Node at the requested 0-based position
func (n *Node) Find(x int) *Node {
	if x == 0 {
		return n
	}
	if n.next != nil {
		return n.next.Find(x - 1)
	}
	return nil
}

// PopNext pops the next node off the current node
func (n *Node) PopNext() *Node {
	next := n.next
	if next != nil {
		n.next = nil
		next.prev = nil
	}
	return next
}

// Last gets the last Node in the linked list
func (n *Node) Last() *Node {
	if n.next != nil {
		return n.next.Last()
	}
	return n
}

// Pop removed the Node from the linked list
//
//	Note that this is not a traditional linked list pop method and would fragment a list if used in the middle
func (n *Node) Pop() *Node {
	if n.prev != nil {
		n.prev.next = nil
		n.prev = nil
	}

	if n.next != nil {
		n.next.prev = nil
		n.next = nil
	}

	return n
}

// Value is a helper to convert rune to string char for convenience
func (n *Node) Value() string {
	return string(n.value)
}

// Queue is a FIFO queue utilizing a linked list
type Queue struct {
	head *Node
	tail *Node
}

// NewQueue creates a FIFO queue with a pre-made head and tail for an empty linked list
func NewQueue() *Queue {
	h := &Node{}
	t := &Node{}
	h.next = t
	t.prev = h
	return &Queue{
		head: h,
		tail: t,
	}
}

// Enqueue adds the linked Node(s) to the front of the queue
//
//	Example: add A-B-C to queue of head-D-E-F-tail
//	result is: head-A-B-C-D-E-F-tail
func (q *Queue) Enqueue(n *Node) {
	last := n.Last()
	next := q.head.next
	last.next = next
	next.prev = last
	n.prev = q.head
	q.head.next = n
}

// ReverseEnqueue reverses the linked list Node(s) and adds to the front of the queue
//
//	Example: add A-B-C to queue of head-D-E-F-tail
//	result is: head-C-B-A-D-E-F-tail
func (q *Queue) ReverseEnqueue(n *Node) {
	curr := n
	for curr.next != nil {
		curr.prev = nil
		next := curr.next
		curr.next = nil
		q.Enqueue(curr)
		curr = next
	}
	curr.prev = nil
	q.Enqueue(curr)
}

// Dequeue will return up to the requested number of linked Node(s) from the front of the queue
// or less if the requested count cannot be satisfied due to insufficient queue length
func (q *Queue) Dequeue(count int) *Node {
	if count <= 0 {
		return nil
	}

	next := q.head.next
	if next == q.tail {
		return nil
	}

	last := next.Find(count - 1)
	if last != nil && last != q.tail {
		newNext := last.PopNext()
		q.head.next = newNext
		newNext.prev = q.head
	} else { // this should ensure that head and tail are not lost be reattaching the two
		tail := q.tail.prev.PopNext()
		tail.prev = q.head
		q.head.next = q.tail
	}

	// remove ref to head
	next.prev = nil

	return next
}

// AddToBack is essentially a LIFO mechanism to make it easier to construct our queue
// from the starting diagram
//
//	Example: add A-B-C to back of queue with head-D-E-F-tail
//	Result: head-D-E-F-A-B-C-tail
func (q *Queue) AddToBack(n *Node) {
	last := n.Last()
	oldLast := q.tail.prev
	oldLast.next = n
	n.prev = oldLast
	last.next = q.tail
	q.tail.prev = last
}
