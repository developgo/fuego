package fuego

// Ideas of interfaces:
// Walker (FirstRest), Mutator (Insert, Delete, Merge), Sizer, Streamer
// Should only create them when they are required - this is not Java

// A Set is a composition of interfaces for all Set kinds.
type Set interface {
	Mutator
	Walker
}

type Mutator interface {
	Insert(e Entry) Mutator
	Delete(e Entry) Mutator
}

type Walker interface {
	FirstRest() (Entry, Walker)
	Sizer
}

type Sizer interface {
	Size() int
}

type Streamer interface {
	Stream() Stream
}
