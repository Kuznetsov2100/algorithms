package sort

// sorting Comparable type using insertion sort.
// In the worst case, this implementation makes ~ 1/2*n^2 compares
// and ~ 1/2*n^2 exchanges to sort an array of length n.
// So, it is not suitable for sorting large arbitrary arrays.
// This sorting algorithm is stable. It uses O(1) extra memory (not including the input array).
func InsertionSort(data Comparable) {
	n := data.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// sorting Comparable type using selection sort.
// This implementation makes ~ 1/2*n^2 compares to sort any array of length n,
// So it is not suitable for sorting large arrays. It performs exactly n exchanges.
// This sorting algorithm is not stable. It uses Θ(1) extra memory (not including the input array).
func SelectionSort(data Comparable) {
	n := data.Len()
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}

// sorting Comparable type using shell sort.
// In the worst case, this implementation makes O(n^3/2) compares and exchanges to sort an array of length n.
// This sorting algorithm is not stable. It uses Θ(1) extra memory (not including the input array).
func ShellSort(data Comparable) {
	n := data.Len()
	for gap := n / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < n; i++ {
			j := i
			for j-gap >= 0 && data.Less(j, j-gap) {
				data.Swap(j, j-gap)
				j = j - gap
			}
		}
	}
}

// IsSorted reports whether data is sorted.
func IsSorted(data Comparable) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
