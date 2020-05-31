package priorityqueue

type Key interface {
	CompareTo(k Key) int
}
