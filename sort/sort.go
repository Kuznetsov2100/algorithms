package sort

// BubbleSort sorts Comparable type using bubble sort.
// This implementation has a worst-case and average complexity of О(n^2),
// where n is the number of items being sorted.
// This sorting algorithm is stable.
func BubbleSort(data Comparable) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		isChanged := false
		for j := 0; j < n-1-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}

// InsertionSort sorts Comparable type using insertion sort.
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

func insertionSortA2B(data Comparable, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// SelectionSort sorts Comparable type using selection sort.
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

// ShellSort sorts Comparable type using shell sort.
// In the worst case, this implementation makes O(n^3/2) compares and exchanges to sort an array of length n.
// This sorting algorithm is not stable. It uses Θ(1) extra memory (not including the input array).
func ShellSort(data Comparable) {
	n := data.Len()
	for gap := n / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < n; i++ {
			j := i
			for ; j-gap >= 0 && data.Less(j, j-gap); j = j - gap {
				data.Swap(j, j-gap)
			}
		}
	}
}

// MergeSort sorts Comparable type using a top-down, recursive version of mergesort.
// This implementation takes O(n*logn) time to sort any array of length n (assuming comparisons take constant time).
// It makes between ~ 1/2*N*log2N and ~ 1*N*log2N compares.
// This sorting algorithm is stable. It uses O(N) extra memory (not including the input array).
func MergeSort(data Comparable) {
	aux := make([]int, data.Len())
	var sort func(data Comparable, low, high int)
	merge := func(data Comparable, low, mid, high int) {
		li, ls := low, mid
		ri, rs := mid, high
		cursor := 0
		for ; li < ls && ri < rs; cursor++ {
			if data.Less(li, ri) {
				aux[li-low] = cursor
				li++
			} else {
				aux[ri-low] = cursor
				ri++
			}
		}
		for ; li < ls; li++ {
			aux[li-low] = cursor
			cursor++
		}
		for i := range aux[:cursor] {
			for j := aux[i]; j != i; {
				data.Swap(low+i, low+j)
				aux[i], aux[j], j = aux[j], aux[i], aux[j]
			}
		}
	}
	sort = func(data Comparable, low, high int) {
		size := high - low
		if size < 2 {
			return
		}
		mid := (low + high) / 2
		sort(data, low, mid)
		sort(data, mid, high)
		merge(data, low, mid, high)
	}
	sort(data, 0, data.Len())
}

// QuickSort sorts Comparable type using quicksort.
// This sorting algorithm is not stable.
func QuickSort(data Comparable) {
	var sort func(data Comparable, lo, hi int)
	// partition the subarray a[lo..hi] so that a[lo..j-1] <= a[j] <= a[j+1..hi]
	// and return the index j.
	partition := func(data Comparable, lo, hi int) int {
		pivot, i, j := lo, lo+1, hi
		for {
			// stop the for loop if a[i] >= pivot
			for ; data.Less(i, pivot); i++ {
				if i == hi { // scan from left to right finished
					break
				}
			}
			// stop the for loop if a[j] <= pivot
			for data.Less(pivot, j) {
				j--
			}
			// scan finished
			if i >= j {
				break
			}
			// i < j, a[i] >= pivot, a[j] <= pivot, call swap(i, j) to let a[i] <= a[j]
			data.Swap(i, j)
		}
		// put pivot at a[j]
		data.Swap(lo, j)
		// now, a[lo..j-1] <= a[j] <= a[j+1..hi]
		return j
	}

	sort = func(data Comparable, lo, hi int) {
		if hi <= lo {
			return
		}
		j := partition(data, lo, hi)
		sort(data, lo, j-1)
		sort(data, j+1, hi)
	}
	data.Shuffle()
	sort(data, 0, data.Len()-1)
}

func median3(data Comparable, i, j, k int) int {
	if data.Less(i, j) {
		if data.Less(j, k) {
			return j
		}
		if data.Less(i, k) {
			return k
		}
		return i
	}
	if data.Less(k, j) {
		return j
	}
	if data.Less(k, i) {
		return k
	}
	return i
}

// QuickSort2Way sorts Comparable type uses the Hoare's 2-way partitioning scheme, chooses the partitioning
// element using median-of-3, and cuts off to insertion sort.
// This sorting algorithm is not stable.
func QuickSort2Way(data Comparable) {
	const INSERTION_SORT_CUTOFF = 8
	var sort func(data Comparable, lo, hi int)
	partition := func(data Comparable, lo, hi int) int {
		n := hi - lo + 1
		m := median3(data, lo, lo+n/2, hi)
		data.Swap(m, lo)
		pivot, i, j := lo, lo+1, hi

		for ; data.Less(i, pivot); i++ {
			if i == hi {
				data.Swap(lo, hi)
				return hi
			}
		}

		for ; data.Less(pivot, j); j-- {
			if j == lo+1 {
				return lo
			}
		}

		for i < j {
			data.Swap(i, j)
			for data.Less(i, pivot) {
				i++
			}
			for data.Less(pivot, j) {
				j--
			}
		}
		data.Swap(lo, j)
		return j
	}

	sort = func(data Comparable, lo, hi int) {

		if hi <= lo {
			return
		}
		n := hi - lo + 1
		if n <= INSERTION_SORT_CUTOFF {
			insertionSortA2B(data, lo, hi+1)
			return
		}

		j := partition(data, lo, hi)
		sort(data, lo, j-1)
		sort(data, j+1, hi)
	}

	data.Shuffle()
	sort(data, 0, data.Len()-1)

}

// QuickSort3Way sorts Comparable type uses the 3-way partitioning scheme.
// This sorting algorithm is not stable.
func QuickSort3Way(data Comparable) {
	var sort func(data Comparable, lo, hi int)
	sort = func(data Comparable, lo, hi int) {
		if hi <= lo {
			return
		}

		i, lt, gt := lo+1, lo, hi
		for i <= gt {
			if data.Less(i, lt) {
				data.Swap(lt, i)
				lt++
				i++
			} else if data.Less(lt, i) {
				data.Swap(i, gt)
				gt--
			} else {
				i++
			}
		}
		sort(data, lo, lt-1)
		sort(data, gt+1, hi)
	}
	data.Shuffle()
	sort(data, 0, data.Len()-1)
}

// HeapSort sorts comparable type using heapsort.HeapSort
// This implementation takes O(n*logN) time to sort any array of length n (assuming comparisons take constant time).
// It makes at most 2*n*log2N compares.
// This sorting algorithm is not stable. It uses O(1) extra memory (not including the input array).
func HeapSort(data Comparable) {
	sink := func(data Comparable, k, N int) {
		for 2*k <= N {
			j := 2 * k
			if j < N && data.Less(j-1, j) {
				j++
			}
			if !data.Less(k-1, j-1) {
				break
			}
			data.Swap(k-1, j-1)
			k = j
		}
	}
	N := data.Len()
	for k := N / 2; k >= 1; k-- {
		sink(data, k, N)
	}
	for i := N; i > 1; {
		data.Swap(0, i-1)
		i--
		sink(data, 1, i)
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
