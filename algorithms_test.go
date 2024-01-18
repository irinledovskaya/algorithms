package algorithms_test

import (
	algorithms "algorithms/numbers"
	sorting "algorithms/sorting"
	"fmt"
	"testing"
)

func TestFindFactors(t *testing.T) {
	factors := algorithms.FindFactors(568)
	fmt.Println(factors)
}

func TestInsertionSort(t *testing.T) {
	lenList := 10

	input := sorting.FillLinkedList(lenList)
	fmt.Println(input)

	input = sorting.InsertionSort(input)
	fmt.Println(input)
}

func TestSelectionSort(t *testing.T) {
	lenList := 10

	input := sorting.FillLinkedList(lenList)
	fmt.Println(input)

	input = sorting.SelectionSort(input)
	fmt.Println(input)
}
