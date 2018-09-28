package gorpi

import (
	"errors"
	"fmt"
)

type SliceSet []int

func (set SliceSet) IsEmpty() bool {
	return len(set) == 0
}

func (set SliceSet) Contains(elem int) bool {
	for _, current := range set {
		if elem == current {
			return true
		}
	}
	return false
}

// Insert returns true if the elem was inserted
func (set *SliceSet) Insert(elem int) bool {
	if set.Contains(elem) {
		return false
	}
	(*set) = append((*set), elem)
	return true
}

func (set *SliceSet) SelectOne() (int, error) {
	if set.IsEmpty() {
		return 0, errors.New("set is empty")
	}
	elem := (*set)[0]
	set.Remove(elem)
	return elem, nil
}

// Remove returns true if elem was removed
func (set *SliceSet) Remove(elem int) bool {
	for index, current := range *set {
		if current == elem {
			(*set)[index] = (*set)[len((*set))-1]
			(*set) = (*set)[:len((*set))-1]
			return true
		}
	}
	return false
}

func (set SliceSet) String() string {
	return fmt.Sprintf("{ %v }", []int(set))
}
