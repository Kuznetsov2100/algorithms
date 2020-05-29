package sort

// Comparable is an interface which should be satified by a type, typically a collection,
// then the collection can be sorted by the routines in this package. The methods require that the
// elements of the collection can be enumerated by an integer index.
type Comparable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
