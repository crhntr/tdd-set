package gorpi

import "fmt"

type MapSet struct{ m map[int]struct{} }

// IsEmpty does what you'd expect it to do...
func (set MapSet) IsEmpty() bool {
	return len(set.m) == 0
}

func (set MapSet) Contains(elem int) bool {
	_, ok := set.m[elem]
	return ok
}

// Insert returns true if the elem was inserted
func (set *MapSet) Insert(elem int) bool {
	if set.m == nil {
		set.m = make(map[int]struct{})
	}
	if set.Contains(elem) {
		return false
	}
	set.m[elem] = struct{}{}
	return true
}

// Remove returns true if elem was removed
func (set MapSet) Remove(elem int) bool {
	if set.IsEmpty() || !set.Contains(elem) {
		return false
	}
	delete(set.m, elem)
	return true
}

func (set MapSet) String() string {
	str := "{ "
	for elem, _ := range set.m {
		str += fmt.Sprintf("%d ", elem)
	}
	return str + "}"
}
