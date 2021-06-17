package list

import "container/list"

// PreviousCircularElement returns the previous list element in a circular way, or nil if there is no such element.
func PreviousCircularElement(list *list.List, current *list.Element) *list.Element {
	if list == nil || current == nil {
		return nil
	}

	element := current.Prev()
	if element == nil {
		element = list.Back()
	}
	return element
}

// NextCircularElement returns the next list element in a circular way, or nil if there is no such element.
func NextCircularElement(list *list.List, current *list.Element) *list.Element {
	if list == nil || current == nil {
		return nil
	}

	element := current.Next()
	if element == nil {
		element = list.Front()
	}
	return element
}
