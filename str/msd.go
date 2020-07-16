package str

const asciiR = 256 // extended ASCII alphabet size
const cutOFF = 15  // cutoff to insertion sort

// MsdSort rearranges the array of extended ASCII strings in ascending order using MSD radix sort.
func MsdSort(a []string) {
	n := len(a)
	aux := make([]string, n)
	msdsort(a, 0, n-1, 0, aux)
}

// msdsort from a[lo] to a[hi], starting at the dth character
func msdsort(a []string, lo, hi, d int, aux []string) {
	// cutoff to insertion sort for small subarrays
	if hi <= lo+cutOFF {
		insertion(a, lo, hi, d)
		return
	}

	// compute frequency counts
	count := make([]int, asciiR+2)
	for i := lo; i <= hi; i++ {
		c := charAt(a[i], d)
		count[c+2]++
	}

	// transform counts to indicies
	for r := 0; r < asciiR+1; r++ {
		count[r+1] += count[r]
	}

	// distribute
	for i := lo; i <= hi; i++ {
		c := charAt(a[i], d)
		aux[count[c+1]] = a[i]
		count[c+1]++
	}

	// copy back
	for i := lo; i <= hi; i++ {
		a[i] = aux[i-lo]
	}

	// recursively sort for each character (excludes sentinel -1)
	for r := 0; r < asciiR; r++ {
		msdsort(a, lo+count[r], lo+count[r+1]-1, d+1, aux)
	}
}

// insertion sort a[lo..hi], starting at dth character
func insertion(a []string, lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && less(a[j], a[j-1], d); j-- {
			exch(a, j, j-1)
		}
	}
}

// return dth character of s, -1 if d = length of string
func charAt(s string, d int) int {
	if d < 0 || d > len(s) {
		panic("string index outbound")
	}
	if d == len(s) {
		return -1
	}
	return int(s[d])
}

// exchange a[i] and a[j]
func exch(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// is v less than w, starting at character d
func less(v, w string, d int) bool {
	for i := d; i < min(len(v), len(w)); i++ {
		if v[i] < w[i] {
			return true
		} else if v[i] > w[i] {
			return false
		}
	}
	return len(v) < len(w)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
