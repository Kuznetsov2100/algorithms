package searching

// Key interface describes the requirements for a type using the routines in the package.
type Key interface {
	CompareTo(k Key) int
}
