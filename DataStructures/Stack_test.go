package datastructures

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack()
	s.push(&Node{1})
	if s.nodes[0].Value != 1 {
		t.Error("Push operation failed to add 1")
	}
}

func TestPop(t *testing.T) {
	s := NewStack()
	s.push(&Node{1})
	s.push(&Node{2})
	s.push(&Node{3})
	s.pop()

	if len(s.nodes) != 2 {
		t.Error("Pop failed to pop 3")
	}
	if s.nodes[0].Value != 1 ||
		s.nodes[1].Value != 2 {
		t.Error("Incorrect elements in stack after pop")
	}
}

func TestPopNil(t *testing.T) {
	s := NewStack()
	x := s.pop()
	if x != nil {
		t.Error("Nil should be returned when pop an empty stack")
	}
}
