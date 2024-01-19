package entities

import "errors"

type Queue struct {
	first *DLinkedList
	last  *DLinkedList
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) String() string {
	if q.first == nil {
		return ""
	}
	return q.first.String()
}

func (q *Queue) Enqueue(new int) {
	if q.first == nil {
		q.first = &DLinkedList{
			Value: new,
		}
		q.last = q.first
		return
	}
	newElem := &DLinkedList{
		Value: new,
	}
	q.last.Next = newElem
	newElem.Prev = q.last
	q.last = newElem
}

func (q *Queue) Dequeue() (int, error) {
	if q.first == nil {
		return 0, errors.New("no element to delete")
	}
	res := q.first.Value
	if q.first.Next == nil {
		q.first = nil
		q.last = nil
		return res, nil
	}
	q.first.Next.Prev = nil
	q.first = q.first.Next

	return res, nil
}
