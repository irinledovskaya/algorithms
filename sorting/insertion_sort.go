package sorting

import (
	e "algorithms/entities"
)

// сортировка вставкой для однонаправленного списка
func LLInsertionSort(input *e.LinkedList) *e.LinkedList {
	if input == nil || input.Next == nil {
		return input
	}

	sentinel := &e.LinkedList{}
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

func SliceInsertionSort(s []int) []int {
	for i := range s {
		for j := 0; j < i; j++ {
			if s[j] > s[i] {
				s[j], s[i] = s[i], s[j]
			}
		}
	}
	return s
}
