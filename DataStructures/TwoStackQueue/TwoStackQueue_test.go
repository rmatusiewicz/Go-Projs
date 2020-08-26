package datastructures

import (
	"testing"
)

func TestTwoStackEnqueue(t *testing.T) {
	q := NewTwoStackQueue()
	q.enqueue(&Node{1})

	if q.getSize() != 1 {
		t.Error("Enqueue did not work")
	}

	if q.peek().Value != 1 {
		t.Errorf("Enqueue value incorrect %d", q.peek().Value)
	}
}

func TestTwoStackComplexQueue(t *testing.T) {
	q := NewTwoStackQueue()
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

func TestTwoStackComplexQueue2(t *testing.T) {
	q := NewTwoStackQueue()
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

func TestTwoStackNoElements(t *testing.T) {
	q := NewTwoStackQueue()
	if q.dequeue() != nil {
		t.Error("No elements in queue to dequeue")
	}
}
