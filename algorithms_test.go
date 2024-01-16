package algorithms_test

import (
	algorithms "algorithms/numbers"
	sorting "algorithms/sorting"
	"fmt"
	"math/rand"
	"testing"
)

func TestFindFactors(t *testing.T) {
	factors := algorithms.FindFactors(568)
	fmt.Println(factors)
}

func TestInsertionSort(t *testing.T) {
	lenList := 10

	var input, next *sorting.LinkedList
	input = &sorting.LinkedList{
		Value: rand.Intn(10),
	}

	next = input
	for i := 0; i < lenList-1; i++ {
		next.Next = &sorting.LinkedList{
			Value: rand.Intn(10),
		}
		next = next.Next
	}
	fmt.Println(input)

	input = sorting.InsertionSort(input)
	fmt.Println(input)
}
