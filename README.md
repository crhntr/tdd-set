# TDD Set in Go (golang)
A set implemented with different underlying data structures. The testing is the intresting part of this repo. It is an example of a factory type thing to test multiple implementations of the same interface.

MapSet, LLSet, and SliceSet all implement the Set interface (defined in set_test.go). As they are all sets, they all should behave alike.
The single test suite in set_test wraps anything that implements the Set interfaces and tests it for proper behavior.

Here is what it looks like to test the is empty method:

```go

type Set interface {
	IsEmpty() bool
    Insert(elem int) bool
}

func TestBehavior(t *testing.T) {
	t.Run("MapSet", injectSetToTests(t, func() Set { return &gorpi.MapSet{} }))
	t.Run("SliceSet", injectSetToTests(t, func() Set { return &gorpi.SliceSet{} }))
	t.Run("LLSet", injectSetToTests(t, func() Set { return &gorpi.LLSet{} }))
}

func injectSetToTests(t *testing.T, newSet func() Set) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("IsEmpty", func(t *testing.T) { testSet_IsEmpty(t, newSet) })
        // wrapping other methods
	}
}

func testSet_IsEmpty(t *testing.T, newSet func() Set) {
	t.Run("when a new is created", func(t *testing.T) {
		set := newSet()
		if !set.IsEmpty() {
			t.Error("it should be empty")
		}
	})

	t.Run("after an element has been inserted inserted", func(t *testing.T) {
		set := newSet()
		set.Insert(420)

		if set.IsEmpty() {
			t.Error("it should not be empty")
		}
	})
}
```

## Full Test Output
```
=== RUN   TestBehavior
=== RUN   TestBehavior/MapSet
=== RUN   TestBehavior/MapSet/IsEmpty
=== RUN   TestBehavior/MapSet/IsEmpty/when_a_new_set_is_empty
=== RUN   TestBehavior/MapSet/IsEmpty/when_a_set_has_had_an_element_inserted
=== RUN   TestBehavior/MapSet/Contains
=== RUN   TestBehavior/MapSet/Contains/when_a_new_set_is_empty
=== RUN   TestBehavior/MapSet/Contains/when_a_set_has_an_element_inserted
=== RUN   TestBehavior/MapSet/Insert
=== RUN   TestBehavior/MapSet/Insert/when_a_new_set_is_empty
=== RUN   TestBehavior/MapSet/Insert/when_a_set_already_has_an_element
=== RUN   TestBehavior/MapSet/Remove
=== RUN   TestBehavior/MapSet/Remove/when_a_new_set_is_empty
=== RUN   TestBehavior/MapSet/Remove/when_a_set_is_contains_one_element
=== RUN   TestBehavior/MapSet/Remove/when_an_existing_element_is_removed
=== RUN   TestBehavior/MapSet/Remove/when_an_existing_element_is_removed_from_a_larger_set
=== RUN   TestBehavior/MapSet/SelectOne
=== RUN   TestBehavior/MapSet/SelectOne/when_a_new_set_is_empty
=== RUN   TestBehavior/MapSet/SelectOne/when_a_set_has_one_number
=== RUN   TestBehavior/MapSet/SelectOne/when_a_set_has_multiple_elements
=== RUN   TestBehavior/MapSet/String
=== RUN   TestBehavior/SliceSet
=== RUN   TestBehavior/SliceSet/IsEmpty
=== RUN   TestBehavior/SliceSet/IsEmpty/when_a_new_set_is_empty
=== RUN   TestBehavior/SliceSet/IsEmpty/when_a_set_has_had_an_element_inserted
=== RUN   TestBehavior/SliceSet/Contains
=== RUN   TestBehavior/SliceSet/Contains/when_a_new_set_is_empty
=== RUN   TestBehavior/SliceSet/Contains/when_a_set_has_an_element_inserted
=== RUN   TestBehavior/SliceSet/Insert
=== RUN   TestBehavior/SliceSet/Insert/when_a_new_set_is_empty
=== RUN   TestBehavior/SliceSet/Insert/when_a_set_already_has_an_element
=== RUN   TestBehavior/SliceSet/Remove
=== RUN   TestBehavior/SliceSet/Remove/when_a_new_set_is_empty
=== RUN   TestBehavior/SliceSet/Remove/when_a_set_is_contains_one_element
=== RUN   TestBehavior/SliceSet/Remove/when_an_existing_element_is_removed
=== RUN   TestBehavior/SliceSet/Remove/when_an_existing_element_is_removed_from_a_larger_set
=== RUN   TestBehavior/SliceSet/SelectOne
=== RUN   TestBehavior/SliceSet/SelectOne/when_a_new_set_is_empty
=== RUN   TestBehavior/SliceSet/SelectOne/when_a_set_has_one_number
=== RUN   TestBehavior/SliceSet/SelectOne/when_a_set_has_multiple_elements
=== RUN   TestBehavior/SliceSet/String
=== RUN   TestBehavior/LLSet
=== RUN   TestBehavior/LLSet/IsEmpty
=== RUN   TestBehavior/LLSet/IsEmpty/when_a_new_set_is_empty
=== RUN   TestBehavior/LLSet/IsEmpty/when_a_set_has_had_an_element_inserted
=== RUN   TestBehavior/LLSet/Contains
=== RUN   TestBehavior/LLSet/Contains/when_a_new_set_is_empty
=== RUN   TestBehavior/LLSet/Contains/when_a_set_has_an_element_inserted
=== RUN   TestBehavior/LLSet/Insert
=== RUN   TestBehavior/LLSet/Insert/when_a_new_set_is_empty
=== RUN   TestBehavior/LLSet/Insert/when_a_set_already_has_an_element
=== RUN   TestBehavior/LLSet/Remove
=== RUN   TestBehavior/LLSet/Remove/when_a_new_set_is_empty
=== RUN   TestBehavior/LLSet/Remove/when_a_set_is_contains_one_element
=== RUN   TestBehavior/LLSet/Remove/when_an_existing_element_is_removed
=== RUN   TestBehavior/LLSet/Remove/when_an_existing_element_is_removed_from_a_larger_set
=== RUN   TestBehavior/LLSet/SelectOne
=== RUN   TestBehavior/LLSet/SelectOne/when_a_new_set_is_empty
=== RUN   TestBehavior/LLSet/SelectOne/when_a_set_has_one_number
=== RUN   TestBehavior/LLSet/SelectOne/when_a_set_has_multiple_elements
=== RUN   TestBehavior/LLSet/String
--- PASS: TestBehavior (0.00s)
    --- PASS: TestBehavior/MapSet (0.00s)
        --- PASS: TestBehavior/MapSet/IsEmpty (0.00s)
            --- PASS: TestBehavior/MapSet/IsEmpty/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/MapSet/IsEmpty/when_a_set_has_had_an_element_inserted (0.00s)
        --- PASS: TestBehavior/MapSet/Contains (0.00s)
            --- PASS: TestBehavior/MapSet/Contains/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/MapSet/Contains/when_a_set_has_an_element_inserted (0.00s)
        --- PASS: TestBehavior/MapSet/Insert (0.00s)
            --- PASS: TestBehavior/MapSet/Insert/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/MapSet/Insert/when_a_set_already_has_an_element (0.00s)
        --- PASS: TestBehavior/MapSet/Remove (0.00s)
            --- PASS: TestBehavior/MapSet/Remove/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/MapSet/Remove/when_a_set_is_contains_one_element (0.00s)
            --- PASS: TestBehavior/MapSet/Remove/when_an_existing_element_is_removed (0.00s)
            --- PASS: TestBehavior/MapSet/Remove/when_an_existing_element_is_removed_from_a_larger_set (0.00s)
        --- PASS: TestBehavior/MapSet/SelectOne (0.00s)
            --- PASS: TestBehavior/MapSet/SelectOne/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/MapSet/SelectOne/when_a_set_has_one_number (0.00s)
            --- PASS: TestBehavior/MapSet/SelectOne/when_a_set_has_multiple_elements (0.00s)
        --- PASS: TestBehavior/MapSet/String (0.00s)
    --- PASS: TestBehavior/SliceSet (0.00s)
        --- PASS: TestBehavior/SliceSet/IsEmpty (0.00s)
            --- PASS: TestBehavior/SliceSet/IsEmpty/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/SliceSet/IsEmpty/when_a_set_has_had_an_element_inserted (0.00s)
        --- PASS: TestBehavior/SliceSet/Contains (0.00s)
            --- PASS: TestBehavior/SliceSet/Contains/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/SliceSet/Contains/when_a_set_has_an_element_inserted (0.00s)
        --- PASS: TestBehavior/SliceSet/Insert (0.00s)
            --- PASS: TestBehavior/SliceSet/Insert/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/SliceSet/Insert/when_a_set_already_has_an_element (0.00s)
        --- PASS: TestBehavior/SliceSet/Remove (0.00s)
            --- PASS: TestBehavior/SliceSet/Remove/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/SliceSet/Remove/when_a_set_is_contains_one_element (0.00s)
            --- PASS: TestBehavior/SliceSet/Remove/when_an_existing_element_is_removed (0.00s)
            --- PASS: TestBehavior/SliceSet/Remove/when_an_existing_element_is_removed_from_a_larger_set (0.00s)
        --- PASS: TestBehavior/SliceSet/SelectOne (0.00s)
            --- PASS: TestBehavior/SliceSet/SelectOne/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/SliceSet/SelectOne/when_a_set_has_one_number (0.00s)
            --- PASS: TestBehavior/SliceSet/SelectOne/when_a_set_has_multiple_elements (0.00s)
        --- PASS: TestBehavior/SliceSet/String (0.00s)
    --- PASS: TestBehavior/LLSet (0.00s)
        --- PASS: TestBehavior/LLSet/IsEmpty (0.00s)
            --- PASS: TestBehavior/LLSet/IsEmpty/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/LLSet/IsEmpty/when_a_set_has_had_an_element_inserted (0.00s)
        --- PASS: TestBehavior/LLSet/Contains (0.00s)
            --- PASS: TestBehavior/LLSet/Contains/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/LLSet/Contains/when_a_set_has_an_element_inserted (0.00s)
        --- PASS: TestBehavior/LLSet/Insert (0.00s)
            --- PASS: TestBehavior/LLSet/Insert/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/LLSet/Insert/when_a_set_already_has_an_element (0.00s)
        --- PASS: TestBehavior/LLSet/Remove (0.00s)
            --- PASS: TestBehavior/LLSet/Remove/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/LLSet/Remove/when_a_set_is_contains_one_element (0.00s)
            --- PASS: TestBehavior/LLSet/Remove/when_an_existing_element_is_removed (0.00s)
            --- PASS: TestBehavior/LLSet/Remove/when_an_existing_element_is_removed_from_a_larger_set (0.00s)
        --- PASS: TestBehavior/LLSet/SelectOne (0.00s)
            --- PASS: TestBehavior/LLSet/SelectOne/when_a_new_set_is_empty (0.00s)
            --- PASS: TestBehavior/LLSet/SelectOne/when_a_set_has_one_number (0.00s)
            --- PASS: TestBehavior/LLSet/SelectOne/when_a_set_has_multiple_elements (0.00s)
        --- PASS: TestBehavior/LLSet/String (0.00s)
PASS
coverage: 100.0% of statements
ok  	github.com/crhntr/tdd-set	0.013s
```

## Initial Context

A friend of mine wanted to be introduced to practice pair programming, Go, and
[RPIs (Rob’s Pairing Interview)](https://builttoadapt.io/the-developer-hiring-process-is-broken-672bf273c183)
in preparation for an interview. So the initial version of this codebase
was the result of that pairing session. We paired for a little under two hours.
Subsequently, I finishing the linked list and the map implementations.
I also refactored the test suite to follow a more BDD style.

