package classic

import "errors"

type Stack struct {
	top    *Node
	length int
}

func CreateStack() *Stack {
	return &Stack{
		top:    nil,
		length: 0,
	}
}

func (s *Stack) IsEmpty() int {
	if s.length == 0 {
		return 0
	} else {
		return 1
	}
}

func (s *Stack) Size() int {
	return s.length
}

func (s *Stack) Top() *Node {
	return s.top
}

func (s *Stack) Push(value int) {
	node := &Node{Key: value, next: nil}

	if s.IsEmpty() == 0 || s.top == nil {
		s.top = node
	} else {
		node.next = s.top
		s.top = node
	}

	s.length += 1
}

func (s *Stack) Pop() (*Node, error) {
	if s.IsEmpty() == 0 || s.top == nil {
		return nil, errors.New("Stack is empty!")
	} else {
		temp := s.top
		s.top = s.top.next
		s.length = s.length - 1
		return temp, nil
	}
}
