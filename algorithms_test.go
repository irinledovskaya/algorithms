package algorithms_test

import (
	e "algorithms/entities"
	numbers "algorithms/numbers"
	sorting "algorithms/sorting"
	"fmt"
	"testing"
)

func TestFindFactors(t *testing.T) {
	factors := numbers.FindFactors(568)
	fmt.Println(factors)
}

func TestInsertionSort(t *testing.T) {
	lenList := 10

	input := e.FillLinkedList(lenList)
	fmt.Println(input)

	input = sorting.InsertionSort(input)
	fmt.Println(input)
}

func TestSelectionSort(t *testing.T) {
	lenList := 10

	input := e.FillLinkedList(lenList)
	fmt.Println(input)

	input = sorting.SelectionSort(input)
	fmt.Println(input)
}
