package queue

import (
	"container/list"

	"github.com/mindera-gaming/go-data-structure/comparator"
)

// Queue represents a data structure that stores a collection of elements held in a sequence.
// Always try to use the functions provided in this structure.
//
// This is a wrapper from the built-in list in Golang.
// To avoid having problems with this structure,
// we suggest that you do not manually modify `List` (exposed for iteration purposes only).
type Queue struct {
	List       *list.List
	comparator comparator.Compare
}

// New returns an initialized queue.
func New(comparator comparator.Compare) Queue {
	return Queue{
		List:       list.New(),
		comparator: comparator,
	}
}

// Add inserts the specified element into this queue.
func (q *Queue) Add(element interface{}) {
	if q.comparator == nil {
		q.List.PushBack(element)
		return
	}

	for e := q.List.Front(); e != nil; e = e.Next() {
		result := q.comparator(e.Value, element)

		switch result {
		case comparator.Equal:
			q.List.InsertAfter(element, e)
			return
		case comparator.Less:
			continue
		case comparator.Greater:
			q.List.InsertBefore(element, e)
			return
		}
	}
	q.List.PushBack(element)
}

// Peek retrieves, but does not remove, the head of this queue, or returns nil if this queue is empty.
func (q Queue) Peek() interface{} {
	if head := q.List.Front(); head != nil {
		return head.Value
	}
	return nil
}

// Poll retrieves and removes the head of this queue, or returns nil if this queue is empty.
func (q *Queue) Poll() interface{} {
	if head := q.List.Front(); head != nil {
		q.List.Remove(head)
		return head.Value
	}
	return nil
}

// Remove removes the specified (list.)element from this queue.
// Returns false if the operation is not successful.
func (q *Queue) Remove(element *list.Element) bool {
	if element != nil {
		return false
	}
	q.List.Remove(element)
	return true
}

// Len returns the number of elements in the queue.
func (q Queue) Len() int {
	return q.List.Len()
}
