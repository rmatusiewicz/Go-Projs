package datastructures

import (
	"fmt"
	"strconv"
)

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
	}
	return nil

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
