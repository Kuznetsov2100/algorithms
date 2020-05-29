package sort

type Comparable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
