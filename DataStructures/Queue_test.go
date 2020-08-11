package datastructures

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	q.enqueue(&Node{1})

	if len(q.nodes) != 1 {
		t.Error("Enqueue did not work")
	}

	if q.nodes[0].Value != 1 {
		t.Errorf("Enqueue value incorrect %d", q.nodes[0].Value)
	}
}

func TestComplexQueue(t *testing.T) {
	q := NewQueue()
	q.enqueue(&Node{2})
	q.enqueue(&Node{3})
	q.enqueue(&Node{4})
	if q.dequeue().Value != 2 {
		t.Errorf("Should be 2, returned %d", q.dequeue().Value)
	}

	if q.dequeue().Value != 3 {
		t.Errorf("Should be 3, returned %d", q.dequeue().Value)
	}

	if q.dequeue().Value != 4 {
		t.Errorf("Should be 4, returned %d", q.dequeue().Value)
	}

	if q.dequeue() != nil {
		t.Error("Empty q.Nodes should return nil")
	}
}

func TestComplexQueue2(t *testing.T) {
	q := NewQueue()
	q.enqueue(&Node{2})
	if q.dequeue().Value != 2 {
		t.Errorf("Should be 2, returned %d", q.dequeue().Value)
	}
	q.enqueue(&Node{3})
	q.enqueue(&Node{4})

	if q.dequeue().Value != 3 {
		t.Errorf("Should be 3, returned %d", q.dequeue().Value)
	}
	if q.dequeue().Value != 4 {
		t.Errorf("Should be 4, returned %d", q.dequeue().Value)
	}
	q.enqueue(&Node{10})
	q.dequeue()
	if q.dequeue() != nil {
		t.Error("Empty q.Nodes should return nil")
	}
}
