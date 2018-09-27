package gorpi

import (
	"fmt"
	"testing"
)

type Set interface {
	IsEmpty() bool
	Contains(elem int) bool
	Insert(elem int) bool
	Remove(elem int) bool
}

func TestSets(t *testing.T) {
	t.Run("MapSet", testSet(t, func() Set { return &MapSet{} }))
	t.Run("SliceSet", testSet(t, func() Set { return &SliceSet{} }))
	t.Run("LLSet", testSet(t, func() Set { return &LLSet{} }))
}

func testSet(t *testing.T, getSet func() Set) func(*testing.T) {
	return func(t *testing.T) {
		testSet_IsEmpty(t, getSet)
		testSet_Contains(t, getSet)
		testSet_Insert(t, getSet)
		testSet_Remove(t, getSet)
		testSet_String(t, getSet)
	}
}

func testSet_IsEmpty(t *testing.T, getSet func() Set) {
	t.Run("when an empty set is created", func(t *testing.T) {
		set := getSet()
		if !set.IsEmpty() {
			t.Error("it should be empty")
		}
	})

	t.Run("when a set has had an element inserted", func(t *testing.T) {
		set := getSet()
		set.Insert(420)

		if set.IsEmpty() {
			t.Error("it should NOT be empty")
		}
	})
}

func testSet_Contains(t *testing.T, getSet func() Set) {
	t.Run("when an empty set is created", func(t *testing.T) {
		set := getSet()
		if set.Contains(420) {
			t.Error("it should not contain an element")
		}
	})

	t.Run("when a set has an element inserted", func(t *testing.T) {
		set := getSet()
		set.Insert(420)

		if !set.Contains(420) {
			t.Error("it should contain that element")
		}
	})
}

func testSet_Insert(t *testing.T, getSet func() Set) {
	t.Run("when an empty set is created", func(t *testing.T) {
		set := getSet()
		if set.Insert(420) != true {
			t.Error("it should allow adding a value")
		}
	})

	t.Run("when a set already has an element", func(t *testing.T) {
		set := getSet()
		set.Insert(420)
		if set.Insert(420) {
			t.Error("it should not allow adding an element twice")
		}
	})
}

func testSet_Remove(t *testing.T, getSet func() Set) {
	t.Run("when an empty set is created", func(t *testing.T) {
		set := getSet()
		if set.Remove(420) == true {
			t.Error("it should not allow an element to be removed")
		}
	})

	t.Run("when an empty set is created", func(t *testing.T) {
		set := getSet()
		set.Insert(420)
		if set.Remove(420) == false {
			t.Error("it should not allow an element to be removed")
		}
	})

	t.Run("when an existing element is removed", func(t *testing.T) {
		set := getSet()
		set.Insert(420)
		set.Remove(420)
		if set.Remove(420) == true {
			t.Error("it should not allow you to remove it again")
		}
	})

	t.Run("when an existing element is removed from a len(set) > 1", func(t *testing.T) {
		set := getSet()

		set.Insert(0)
		set.Insert(420)
		set.Insert(9000)

		set.Remove(420)
		if set.Remove(420) == true {
			t.Error("it should not allow you to remove it again")
		}
	})
}

func testSet_String(t *testing.T, getSet func() Set) {
	set := getSet()
	set.Insert(420)
	fmt.Sprint(set)
}
