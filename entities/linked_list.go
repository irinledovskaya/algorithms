package entities

import (
	"fmt"
	"math/rand"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (l *LinkedList) String() string {
	if l == nil {
		return ""
	}
	if l.Next == nil {
		return fmt.Sprintf("%d", l.Value)
	}

	s := "{"
	for {
		s += fmt.Sprintf("%d,", l.Value)
		if l.Next == nil {
			break
		}
		l = l.Next
	}
	s += "}"
	return s
}

func FillLinkedList(len int) *LinkedList {
	var first, next *LinkedList
	first = &LinkedList{
		Value: rand.Intn(10),
	}

	next = first
	for i := 0; i < len-1; i++ {
		next.Next = &LinkedList{
			Value: rand.Intn(10),
		}
		next = next.Next
	}
	return first
}
