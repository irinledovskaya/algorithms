package sorting

// сортировка вставкой для однонаправленного списка
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
