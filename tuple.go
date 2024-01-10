package generichelper

// A tuple is a data structure that has a specific number and sequence of elements.

// Tuple2 represents a 2-tuple, or pair.
type Tuple2[T1, T2 any] struct {
	Item1 T1
	Item2 T2
}

// NewTuple2 creates [Tuple2].
func NewTuple2[T1, T2 any](item1 T1, item2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{Item1: item1, Item2: item2}
}
