package gorpi

import (
	"errors"
	"fmt"
)

type LLSet struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

func (set LLSet) IsEmpty() bool {
	return set.head == nil
}

func (set LLSet) Contains(elem int) bool {
	current := set.head
	for current != nil {
		if current.data == elem {
			return true
		}
		current = current.next
	}
	return false
}

// Insert returns true if the element was inserted
func (set *LLSet) Insert(elem int) bool {
	if set.Contains(elem) {
		return false
	}

	if set.head == nil {
		set.head = &Node{data: elem}
		return true
	}

	current := set.head
	for current.next != nil {
		current = current.next
	}
	current.next = &Node{data: elem}

	return true
}

// Remove returns true if element was removed
func (set *LLSet) Remove(elem int) bool {
	if set.head == nil {
		return false
	}

	if set.head.data == elem {
		set.head = nil
		return true
	}

	current := set.head
	for current != nil && current.next != nil {
		if current.next.data == elem {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (set *LLSet) SelectOne() (int, error) {
	if set.IsEmpty() {
		return 0, errors.New("set is empty")
	}
	elem := set.head.data
	set.head = set.head.next
	return elem, nil
}

func (set LLSet) String() string {
	current := set.head
	str := "{ "
	for current != nil {
		str += fmt.Sprintf("[%d] -> ", current.data)
		current = current.next
	}
	return str + "nil }"
}
