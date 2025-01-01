package gorpi_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/crhntr/tdd-set"
)

type Set interface {
	IsEmpty() bool
	Contains(elem int) bool
	Insert(elem int) bool
	Remove(elem int) bool
	SelectOne() (int, error)
}

func TestBehavior(t *testing.T) {
	t.Run("MapSet", injectSetToTests(t, func() Set { return &gorpi.MapSet{} }))
	t.Run("SliceSet", injectSetToTests(t, func() Set { return &gorpi.SliceSet{} }))
	t.Run("LLSet", injectSetToTests(t, func() Set { return &gorpi.LLSet{} }))
}

func injectSetToTests(t *testing.T, emptySetConstructor func() Set) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("IsEmpty", func(t *testing.T) { testSet_IsEmpty(t, emptySetConstructor) })
		t.Run("Contains", func(t *testing.T) { testSet_Contains(t, emptySetConstructor) })
		t.Run("Insert", func(t *testing.T) { testSet_Insert(t, emptySetConstructor) })
		t.Run("Remove", func(t *testing.T) { testSet_Remove(t, emptySetConstructor) })
		t.Run("SelectOne", func(t *testing.T) { testSet_SelectOne(t, emptySetConstructor) })
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

func testSet_SelectOne(t *testing.T, newEmptySet func() Set) {
	t.Run("when a new set is empty", func(t *testing.T) {
		set := newEmptySet()

		_, err := set.SelectOne()
		if err == nil {
			t.Error("it should return an error")
		}
	})

	t.Run("when a set has one number", func(t *testing.T) {
		set := newEmptySet()

		num := rand.Int()
		set.Insert(num)

		elem, err := set.SelectOne()
		if err != nil {
			t.Error("it should not return an error")
			t.Log(err)
		}
		if elem != num {
			t.Error("it should return the only number it has")
			t.Logf("i got a %d", elem)
		}
		if !set.IsEmpty() {
			t.Error("it ACTUALLY remove an element")
		}
	})

	t.Run("when a set has multiple elements", func(t *testing.T) {
		set := newEmptySet()

		set.Insert(1)
		set.Insert(2)
		set.Insert(3)

		_, err := set.SelectOne()
		if err != nil {
			t.Error("it should not return an error")
			t.Log(err)
		}

		if set.IsEmpty() {
			t.Error("it should not be empty")
		}
	})
}

func testSet_String(t *testing.T, newEmptySet func() Set) {
	set := newEmptySet()
	set.Insert(420)
	fmt.Sprint(set)
}
