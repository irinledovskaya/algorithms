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

func TestStackPush(t *testing.T) {
	stack := e.NewStack()
	stack.Push(5)
	stack.Push(10)
	stack.Push(0)
	fmt.Println(stack)
}

func TestStackPop(t *testing.T) {
	stack := e.NewStack()
	stack.Push(5)
	stack.Push(10)
	stack.Push(0)
	fmt.Println(stack)
	var err error
	var val int
	for err == nil {
		val, err = stack.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(val)
		}
	}

}
