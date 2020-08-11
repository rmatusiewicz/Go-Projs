package datastructures

type Queue struct {
	nodes []*Node
	next  int
}

func NewQueue() *Queue {
	q := new(Queue)
	q.next = 0
	return q
}

func (q *Queue) enqueue(x *Node) {
	q.nodes = append(q.nodes, x)
}

func (q *Queue) dequeue() *Node {

	if len(q.nodes) <= q.next {
		return nil
	}
	q.next = q.next + 1
	return q.nodes[q.next-1]
}
