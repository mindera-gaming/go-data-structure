package queue

import (
	"container/list"

	"github.com/mindera-gaming/go-data-structure/comparator"
)

// Queue represents a data structure that stores a collection of elements held in a sequence.
type Queue struct {
	list       *list.List
	comparator comparator.Compare
}

// New returns an initialized queue.
func New(comparator comparator.Compare) Queue {
	return Queue{
		list:       list.New(),
		comparator: comparator,
	}
}

// Add inserts the specified element into this queue.
func (q *Queue) Add(element interface{}) {
	if q.comparator == nil {
		q.list.PushBack(element)
		return
	}

	for e := q.list.Front(); e != nil; e = e.Next() {
		result := q.comparator(e.Value, element)

		switch result {
		case comparator.Equal:
			q.list.InsertAfter(element, e)
			return
		case comparator.Less:
			continue
		case comparator.Greater:
			q.list.InsertBefore(element, e)
			return
		}
	}
	q.list.PushBack(element)
}

// Peek retrieves, but does not remove, the head of this queue, or returns nil if this queue is empty.
func (q Queue) Peek() interface{} {
	if head := q.list.Front(); head != nil {
		return head.Value
	}
	return nil
}

// Poll retrieves and removes the head of this queue, or returns nil if this queue is empty.
func (q *Queue) Poll() interface{} {
	if head := q.list.Front(); head != nil {
		q.list.Remove(head)
		return head
	}
	return nil
}

// Len returns the number of elements in the queue.
func (q Queue) Len() int {
	return q.list.Len()
}
