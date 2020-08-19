package goutils

type StackNode struct {
	Val  interface{}
	Next *StackNode
}

type Stack struct {
	Head *StackNode
}

func (s *Stack) Peek() *StackNode {
	return s.Head
}

func (s *Stack) Pop() *StackNode {
	if s.Head != nil {
		temp := s.Head
		s.Head = s.Head.Next
		temp.Next = nil
		return temp
	}

	return nil
}

func (s *Stack) Push(n *StackNode) {
	n.Next = s.Head
	s.Head = n
}
