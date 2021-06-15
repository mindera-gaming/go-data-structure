package list

import "container/list"

// PreviousCircularElement returns the previous list element in a circular way.
func PreviousCircularElement(list list.List, current list.Element) list.Element {
	element := current.Prev()
	if element == nil {
		element = list.Back()
	}
	return *element
}

// NextCircularElement returns the next list element in a circular way.
func NextCircularElement(list list.List, current list.Element) list.Element {
	element := current.Next()
	if element == nil {
		element = list.Front()
	}
	return *element
}
