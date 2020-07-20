package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value int
}

type Stack struct {
	nodes []*Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(x *Node) {
	s.nodes = append(s.nodes, x)
}

func (s *Stack) pop() *Node {
	if len(s.nodes) > 0 {
		// fmt.Println()
		// fmt.Println("Pop: " + strconv.Itoa(s.nodes[len(s.nodes)-1].Value))
		n := s.nodes[len(s.nodes)-1]

		s.nodes = s.nodes[0 : len(s.nodes)-1]
		return n
	} else {
		return nil
	}

}

func (s Stack) disp() {
	fmt.Println("Stack:")
	result := ""
	if len(s.nodes) > 0 {
		for _, x := range s.nodes {
			result = result + strconv.Itoa(x.Value) + ","
		}
		result = result[0 : len(result)-1]
	}

	fmt.Println(result)
}

func main() {
	fmt.Println("Starting")
	s := NewStack()
	s.disp()

	s.push(&Node{3})
	s.push(&Node{5})
	s.push(&Node{8})
	s.push(&Node{1})
	s.push(&Node{0})
	s.disp()

	//3,5,8,1,0

	s.pop()
	s.disp()

	s.pop()
	s.disp()

	fmt.Println("Complete")
}
