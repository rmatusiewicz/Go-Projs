package datastructures

import (
	"testing"
)

func TestTwoStackQueue(t *testing.T) {
	q := NewTwoStackQueue()
	q.enqueue(&Node{1})

	if q.getSize() != 1 {
		t.Error("Enqueue did not work")
	}

	if q.nqStack.nodes[0].Value != 1 {
		t.Errorf("Enqueue value incorrect %d", q.nqStack.nodes[0].Value)
	}
}
