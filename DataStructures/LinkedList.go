package main

import "fmt"

type LinkedNode struct {
	Key  int
	Prev *LinkedNode
	Next *LinkedNode
}
type LinkedList struct {
	head  *LinkedNode
	nodes []*LinkedNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) find(x int) (int, *LinkedNode) {
	for i, n := range list.nodes {
		if x == n.Key {
			return i, n
		}
	}
	return 0, nil
}

func (list *LinkedList) insert(x int) {
	n := &LinkedNode{}
	n.Key = x
	n.Prev = nil

	if list.head == nil {
		n.Next = nil
	} else {
		n.Next = list.head
	}
	list.head = n
	list.nodes = append(list.nodes, n)
}

func (list *LinkedList) delete(x int) {
	i, n := list.find(x)
	if n != nil {
		if n.Prev != nil {
			n.Prev.Next = n.Next
		} else {
			list.head = n.Next
		}

		if n.Next != nil {
			n.Next.Prev = n.Prev
		}
		z := len(list.nodes)
		list.nodes = append(list.nodes[0:i], list.nodes[i:z-1]...)
	}
}

func (list *LinkedList) reverse() {
	if list.head != nil {
		current := list.head
		next := current.Next
		current.Next = nil

		for next != nil {
			current.Prev = next
			currentCopy := current

			current = next
			next = current.Next
			current.Next = currentCopy
		}
		current.Prev = nil
		list.head = current
	}
	//nil->1->2->3->4->nil

	//nil->4->3->2->1->nil
}

func (list *LinkedList) print() {
	for i, n := range list.nodes {
		fmt.Printf("%d : %d", i, n.Key)
		fmt.Println()
	}
}
