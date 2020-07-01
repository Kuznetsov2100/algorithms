package str

import (
	"math/rand"
	"time"
)

// Quick3string rearranges the array of strings in ascending order using 3-way radix quicksort.
func Quick3string(a []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	quick3string(a, 0, len(a)-1, 0)
}

// 3-way string quicksort a[lo..hi] starting at dth character
func quick3string(a []string, lo, hi, d int) {
	// cutoff to insertion sort for small subarrays
	if hi <= lo+cutOFF {
		insertion(a, lo, hi, d)
		return
	}

	lt, gt := lo, hi
	v := charAt(a[lo], d)
	i := lo + 1

	for i <= gt {
		t := charAt(a[i], d)
		if t < v {
			exch(a, lt, i)
			lt++
			i++
		} else if t > v {
			exch(a, i, gt)
			gt--
		} else {
			i++
		}
	}

	// a[lo..lt-1] < v = a[lt..gt] < a[gt+1..hi].
	quick3string(a, lo, lt-1, d)
	if v >= 0 {
		quick3string(a, lt, gt, d+1)
	}
	quick3string(a, gt+1, hi, d)
}
