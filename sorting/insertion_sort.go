package sorting

import (
	"fmt"
)

// сортировка вставкой для однонаправленного списка

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

func InsertionSort(input *LinkedList) *LinkedList {
	if input == nil || input.Next == nil {
		return input
	}

	sentinel := &LinkedList{}
	for input != nil {
		next := input
		input = input.Next
		afterMe := sentinel
		for afterMe.Next != nil && afterMe.Next.Value < next.Value {
			afterMe = afterMe.Next
		}
		next.Next = afterMe.Next
		afterMe.Next = next
	}
	return sentinel.Next
}
