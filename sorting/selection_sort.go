package sorting

import (
	e "algorithms/entities"
)

// сортировка выбором для однонапрвленного списка
func SelectionSort(input *e.LinkedList) *e.LinkedList {
	if input == nil || input.Next == nil {
		return input
	}

	ferst := &e.LinkedList{}
	for input.Next != nil {
		maxAfterMe := input
		maxVal := maxAfterMe.Next.Value
		afterMe := input.Next
		for afterMe.Next != nil {
			if afterMe.Next.Value > maxVal {
				maxAfterMe = afterMe
				maxVal = afterMe.Next.Value
			}
			afterMe = afterMe.Next
		}
		maxElem := maxAfterMe.Next
		maxAfterMe.Next = maxElem.Next
		maxElem.Next = ferst.Next
		ferst.Next = maxElem
	}

	return ferst
}
