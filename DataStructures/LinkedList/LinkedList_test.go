package datastructures

import (
	"testing"
)

func TestInsert(t *testing.T) {
	l := NewLinkedList()
	l.insert(2)

	if len(l.nodes) != 1 || l.head.Key != 2 {
		t.Error("Should have 2 at head")
	}
}

func TestDelete(t *testing.T) {
	l := NewLinkedList()
	l.insert(2)
	l.insert(3)
	l.insert(4)

	l.delete(3)
	if len(l.nodes) != 2 {
		t.Error("Failed to delete element 3")
	}

	l.delete(4)

	if l.head.Key != 2 || l.head.Next != nil {
		t.Error("Failed to delete 3 or 4")
	}
}

func TestReverse(t *testing.T) {
	l := NewLinkedList()
	l.insert(1)
	l.insert(2)
	l.insert(3)
	l.insert(4)
	if l.head.Key != 4 {
		t.Error("Head should be at 4")
	}

	l.reverse()

	if l.head.Key != 1 {
		t.Error("Head should be at 1")
	}

}
