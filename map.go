package fuego

// A Map is an interface for all Map kinds.
// type Map interface {
// func (m HamtMap) BiStream() BiStream
// EntrySet() Set
// KeySet() Set

// Ideas of interfaces:
// Walker (FirstRest), Mutator (Insert, Delete, Merge), Sizer, Streamer
// Seeker (Has*), ...
// Should only create them when they are required - this is not Java

// A Set is a composition of interfaces for all Set kinds.
type Map interface {
	MapMutator
	MapWalker
	Retriever
}

type MapMutator interface {
	Insert(k Entry, v interface{}) MapMutator // take in a Tuple as parms instead? This might enable to further abstract the interface and make it common to Set's Mutator
	Delete(k Entry) MapMutator
}

type MapWalker interface {
	FirstRest() (k Entry, v interface{}, rest Map)
	Sizer
}

type Retriever interface {
	Get(k Entry) interface{}         // TODO return Maybe instead of Entry
	Has(k Entry, v interface{}) bool // TODO return EntryBool????
	HasKey(k Entry) bool             // TODO return EntryBool????
	HasValue(v interface{}) bool     // TODO return EntryBool????
}
