package navigableset

import (
	"container/list"
	"errors"

	"github.com/mindera-gaming/go-data-structure/comparator"
)

// NavigableSet represents a data structure that stores ordered elements with navigation methods.
type NavigableSet struct {
	list       *list.List
	comparator comparator.Compare
}

// New returns an initialized navigable set.
func New(comparator comparator.Compare) (NavigableSet, error) {
	if comparator == nil {
		return NavigableSet{}, errors.New(ErrNilComparator)
	}

	return NavigableSet{
		list:       list.New(),
		comparator: comparator,
	}, nil
}

// Add Adds the specified element to this set if it is not already present.
// Returns false if the operation is not successful.
func (n *NavigableSet) Add(element interface{}) bool {
	if n.comparator == nil {
		return false
	}

	for e := n.list.Front(); e != nil; e = e.Next() {
		result := n.comparator(e.Value, element)

		switch result {
		case comparator.Equal:
			return false
		case comparator.Less:
			continue
		case comparator.Greater:
			n.list.InsertBefore(element, e)
			return true
		}
	}
	n.list.PushBack(element)
	return true
}

// Remove removes the specified element from this set if it is present.
// Returns false if the operation is not successful.
func (n *NavigableSet) Remove(element interface{}) bool {
	if n.comparator == nil {
		return false
	}

	for e := n.list.Front(); e != nil; e = e.Next() {
		result := n.comparator(e.Value, element)

		switch result {
		case comparator.Equal:
			n.list.Remove(e)
			return true
		case comparator.Less:
			continue
		case comparator.Greater:
			return false
		}
	}
	return false
}

// Contains returns true if this set contains the specified element.
func (n NavigableSet) Contains(element interface{}) bool {
	if n.comparator == nil {
		return false
	}

	for e := n.list.Front(); e != nil; e = e.Next() {
		result := n.comparator(e.Value, element)

		switch result {
		case comparator.Equal:
			return true
		case comparator.Less:
			continue
		case comparator.Greater:
			return false
		}
	}
	return false
}

// First returns the first (lowest) element currently in this set or nil if there are no elements.
func (n NavigableSet) First() interface{} {
	if first := n.list.Front(); first != nil {
		return first.Value
	}
	return nil
}

// Last returns the last (highest) element currently in this set or nil if there are no elements.
func (n NavigableSet) Last() interface{} {
	if last := n.list.Back(); last != nil {
		return last.Value
	}
	return nil
}

// Higher returns the least element in this set strictly greater than the given element, or nil if there is no such element.
func (n NavigableSet) Higher(element interface{}) interface{} {
	if n.comparator == nil {
		return nil
	}

	for e := n.list.Front(); e != nil; e = e.Next() {
		result := n.comparator(e.Value, element)

		switch result {
		case comparator.Equal:
			return e.Next().Value
		case comparator.Less:
			continue
		case comparator.Greater:
			return e.Value
		}
	}
	return nil
}

// Ceiling returns the least element in this set greater than or equal to the given element, or nil if there is no such element.
func (n NavigableSet) Ceiling(element interface{}) interface{} {
	if n.comparator == nil {
		return nil
	}

	for e := n.list.Front(); e != nil; e = e.Next() {
		result := n.comparator(e.Value, element)

		switch result {
		case comparator.Equal:
			return e.Value
		case comparator.Less:
			continue
		case comparator.Greater:
			return e.Value
		}
	}
	return nil
}

// Lower returns the greatest element in this set strictly less than the given element, or nil if there is no such element.
func (n NavigableSet) Lower(element interface{}) interface{} {
	if n.comparator == nil {
		return nil
	}

	for e := n.list.Back(); e != nil; e = e.Prev() {
		result := n.comparator(e.Value, element)

		switch result {
		case comparator.Equal:
			return e.Prev().Value
		case comparator.Less:
			return e.Value
		case comparator.Greater:
			continue
		}
	}
	return nil
}

// Floor returns the greatest element in this set less than or equal to the given element, or null if there is no such element.
func (n NavigableSet) Floor(element interface{}) interface{} {
	if n.comparator == nil {
		return nil
	}

	for e := n.list.Back(); e != nil; e = e.Prev() {
		result := n.comparator(e.Value, element)

		switch result {
		case comparator.Equal:
			return e.Value
		case comparator.Less:
			return e.Value
		case comparator.Greater:
			continue
		}
	}
	return nil
}

// TailSet returns a view of the portion of this set whose elements are greater than (or equal to, if inclusive is true) fromElement.
func (n NavigableSet) TailSet(fromElement interface{}, inclusive bool) (set []interface{}) {
	if n.comparator == nil {
		return nil
	}

	for e := n.list.Front(); e != nil; e = e.Next() {
		result := n.comparator(e.Value, fromElement)

		switch result {
		case comparator.Equal:
			if inclusive {
				set = append(set, e.Value)
			}
		case comparator.Less:
			continue
		case comparator.Greater:
			set = append(set, e.Value)
		}
	}
	return
}

// Headset returns a view of the portion of this set whose elements are less than (or equal to, if inclusive is true) toElement.
func (n NavigableSet) Headset(fromElement interface{}, inclusive bool) (set []interface{}) {
	if n.comparator == nil {
		return nil
	}

	for e := n.list.Back(); e != nil; e = e.Prev() {
		result := n.comparator(e.Value, fromElement)

		switch result {
		case comparator.Equal:
			if inclusive {
				set = append(set, e.Value)
			}
		case comparator.Less:
			set = append(set, e.Value)
		case comparator.Greater:
			continue
		}
	}
	return
}

// DescendingSet returns a reverse order view of the elements contained in this set.
func (n NavigableSet) DescendingSet() (set []interface{}) {
	for e := n.list.Back(); e != nil; e = e.Prev() {
		set = append(set, e.Value)
	}
	return
}

// Len returns the number of elements in the set.
func (n NavigableSet) Len() int {
	return n.list.Len()
}
