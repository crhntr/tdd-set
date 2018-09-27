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

func Test(t *testing.T) {
	t.Run("MapSet", testFanOut(t, func() Set { return &MapSet{} }))
	t.Run("SliceSet", testFanOut(t, func() Set { return &SliceSet{} }))
	t.Run("LLSet", testFanOut(t, func() Set { return &LLSet{} }))
}

func testFanOut(t *testing.T, emptySetConstructor func() Set) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("IsEmpty", func(t *testing.T) { testSet_IsEmpty(t, emptySetConstructor) })
		t.Run("Contains", func(t *testing.T) { testSet_Contains(t, emptySetConstructor) })
		t.Run("Insert", func(t *testing.T) { testSet_Insert(t, emptySetConstructor) })
		t.Run("Remove", func(t *testing.T) { testSet_Remove(t, emptySetConstructor) })
		t.Run("String", func(t *testing.T) { testSet_String(t, emptySetConstructor) })
	}
}

func testSet_IsEmpty(t *testing.T, newEmptySet func() Set) {
	t.Run("when a new set is empty", func(t *testing.T) {
		set := newEmptySet()
		if !set.IsEmpty() {
			t.Error("it should be empty")
		}
	})

	t.Run("when a set has had an element inserted", func(t *testing.T) {
		set := newEmptySet()
		set.Insert(420)

		if set.IsEmpty() {
			t.Error("it should NOT be empty")
		}
	})
}

func testSet_Contains(t *testing.T, newEmptySet func() Set) {
	t.Run("when a new set is empty", func(t *testing.T) {
		set := newEmptySet()
		if set.Contains(420) {
			t.Error("it should not contain an element")
		}
	})

	t.Run("when a set has an element inserted", func(t *testing.T) {
		set := newEmptySet()
		set.Insert(420)

		if !set.Contains(420) {
			t.Error("it should contain that element")
		}
	})
}

func testSet_Insert(t *testing.T, newEmptySet func() Set) {
	t.Run("when a new set is empty", func(t *testing.T) {
		set := newEmptySet()
		if set.Insert(420) != true {
			t.Error("it should allow adding a value")
		}
	})

	t.Run("when a set already has an element", func(t *testing.T) {
		set := newEmptySet()
		set.Insert(420)
		if set.Insert(420) {
			t.Error("it should not allow adding an element twice")
		}
	})
}

func testSet_Remove(t *testing.T, newEmptySet func() Set) {
	t.Run("when a new set is empty", func(t *testing.T) {
		set := newEmptySet()
		if set.Remove(420) == true {
			t.Error("it should not allow an element to be removed")
		}
	})

	t.Run("when a set is contains one element", func(t *testing.T) {
		set := newEmptySet()
		set.Insert(420)
		if set.Remove(420) == false {
			t.Error("it should not allow an element to be removed")
		}
	})

	t.Run("when an existing element is removed", func(t *testing.T) {
		set := newEmptySet()
		set.Insert(420)
		set.Remove(420)
		if set.Remove(420) == true {
			t.Error("it should not allow you to remove it again")
		}
	})

	t.Run("when an existing element is removed from a larger set", func(t *testing.T) {
		set := newEmptySet()

		set.Insert(0)
		set.Insert(420)
		set.Insert(9000)

		set.Remove(420)
		if set.Remove(420) == true {
			t.Error("it should not allow you to remove it again")
		}
	})
}

func testSet_String(t *testing.T, newEmptySet func() Set) {
	set := newEmptySet()
	set.Insert(420)
	fmt.Sprint(set)
}
