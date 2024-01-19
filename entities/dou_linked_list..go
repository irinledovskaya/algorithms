package entities

import (
	"fmt"
	"math/rand"
)

type DLinkedList struct {
	Value int
	Prev  *DLinkedList
	Next  *DLinkedList
}

func (d *DLinkedList) String() string {
	if d == nil {
		return ""
	}
	if d.Next == nil {
		return fmt.Sprintf("{%d}", d.Value)
	}

	s := "{"
	s += fmt.Sprintf("%d", d.Value)
	d = d.Next
	for {
		s += fmt.Sprintf(",%d", d.Value)
		if d.Next == nil {
			break
		}
		d = d.Next
	}
	s += "}"
	return s
}

func FillDLinkedList(len int) *DLinkedList {
	var first, next *DLinkedList
	first = &DLinkedList{
		Value: rand.Intn(10),
	}

	next = first
	for i := 0; i < len-1; i++ {
		next.Next = &DLinkedList{
			Value: rand.Intn(10),
			Prev:  next,
		}
		next = next.Next
	}
	return first
}
