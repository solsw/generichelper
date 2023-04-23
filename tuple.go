package generichelper

// A tuple is a data structure that has a specific number and sequence of elements.

// Tuple2 represents a 2-tuple, or pair.
type Tuple2[T1, T2 any] struct {
	Item1 T1
	Item2 T2
}
