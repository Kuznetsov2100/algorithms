package priorityqueue

// Key is an interface
type Key interface {
	CompareTo(k Key) int
}
