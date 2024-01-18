package entities

import "errors"

type Stack struct {
	sentinel *LinkedList
}

func NewStack() *Stack {
	return &Stack{
		sentinel: &LinkedList{},
	}
}

func (s *Stack) String() string {
	if s.sentinel == nil || s.sentinel.Next == nil {
		return ""
	}
	return s.sentinel.Next.String()
}

func (s *Stack) Push(new int) {
	if s.sentinel == nil {
		s.sentinel = &LinkedList{}
	}
	newElem := &LinkedList{
		Value: new,
	}

	newElem.Next = s.sentinel.Next
	s.sentinel.Next = newElem
}

func (s *Stack) Pop() (int, error) {
	if s.sentinel == nil || s.sentinel.Next == nil {
		return 0, errors.New("no element to delete")
	}
	res := s.sentinel.Next.Value
	s.sentinel.Next = s.sentinel.Next.Next
	return res, nil
}
