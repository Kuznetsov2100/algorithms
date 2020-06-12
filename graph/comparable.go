package graph

// Comparable interface describes the requirements for a type using the routines in the package.
type Comparable interface {
	CompareTo(c Comparable) int
}
