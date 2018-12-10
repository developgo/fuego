package fuego

import (
	"github.com/raviqqe/hamt"
)

// A HamtSet is an unnaturally ordered set
type HamtSet struct {
	set hamt.Set
}

// NewHamtSet creates a new HamtSet
func NewHamtSet() HamtSet {
	return HamtSet{
		set: hamt.NewSet(),
	}
}

// Stream returns a sequential Stream with this collection as its source.
func (s HamtSet) Stream() Stream {
	return NewStream(NewSetIterator(s))
}

// Insert inserts a value into a set.
func (s HamtSet) Insert(e Entry) Mutator {
	return HamtSet{
		set: s.set.Insert(e.(hamt.Entry)),
	}
}

// Delete a value from a set.
func (s HamtSet) Delete(e Entry) Mutator {
	if e == nil {
		panic(PanicNoSuchElement)
	}
	return HamtSet{
		set: s.set.Delete(e),
	}
}

// Size of the HamtSet.
func (s HamtSet) Size() int {
	return s.set.Size()
}

// FirstRest returns a value in a set and a rest of the set.
// This method is useful for iteration.
func (s HamtSet) FirstRest() (Entry, Walker) {
	e, s2 := s.set.FirstRest()
	if e == nil {
		return nil, NewHamtSet()
	}
	return e.(Entry), HamtSet{set: s2}
}

// Merge merges 2 sets into one.
// TODO: this is wrong - it shouldn be part of the interface, instead it should be a flat method, without receiver.
// func (s HamtSet) Merge(t HamtSet) Mutator {
// 	return HamtSet{
// 		set: s.set.Merge((t.set)),
// 	}
// }
