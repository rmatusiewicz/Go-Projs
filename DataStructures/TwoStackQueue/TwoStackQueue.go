package datastructures

import "fmt"

func main() {
	fmt.Println("Implementing 2 stack queue")

}

type TwoStackQueue struct {
	nqStack, dqStack *Stack
}

func NewTwoStackQueue() *TwoStackQueue {
	stack1 := NewStack()
	stack2 := NewStack()
	s := &TwoStackQueue{}
	s.nqStack = stack1
	s.dqStack = stack2
	return s
}

func (q *TwoStackQueue) peek() (x *Node) {
	if len(q.dqStack.nodes) > 0 {
		return q.dqStack.nodes[len(q.dqStack.nodes)]
	}
	return q.nqStack.nodes[0]
}

func (q *TwoStackQueue) enqueue(x *Node) {
	q.nqStack.nodes = append(q.nqStack.nodes, x)
}

func (q *TwoStackQueue) dequeue() (x *Node) {
	if len(q.dqStack.nodes) > 0 { //dequeue operation just performed so dqStack still populated
		return q.dqStack.pop()
	}
	q.moveAllNodesOver()
	return q.dqStack.pop()
}

func (q *TwoStackQueue) moveAllNodesOver() {
	if len(q.nqStack.nodes) > 0 { //Meaning move nodes to dqStack
		for range q.nqStack.nodes {
			q.dqStack.nodes = append(q.dqStack.nodes, q.nqStack.pop())
		}
	} else { //Meaning move nodes to nqStack
		for range q.dqStack.nodes {
			q.nqStack.nodes = append(q.nqStack.nodes, q.dqStack.pop())
		}
	}
}

func (q *TwoStackQueue) getSize() int {
	return len(q.nqStack.nodes) + len(q.dqStack.nodes)
}
