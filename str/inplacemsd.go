package str

// InplaceMsdSort rearranges the array of extended ASCII strings in ascending order using in-place MSD radix sort.
func InplaceMsdSort(a []string) {
	inplaceSort(a, 0, len(a)-1, 0)
}

// sort from a[lo] to a[hi], starting at the dth character
func inplaceSort(a []string, lo, hi, d int) {
	// cutoff to insertion sort for small subarrays
	if hi <= lo+cutOFF {
		insertion(a, lo, hi, d)
		return
	}

	// compute frequency counts
	heads := make([]int, asciiR+2)
	tails := make([]int, asciiR+1)
	for i := lo; i <= hi; i++ {
		c := charAt(a[i], d)
		heads[c+2]++
	}

	// transform counts to indices
	heads[0] = lo
	for r := 0; r < asciiR+1; r++ {
		heads[r+1] += heads[r]
		tails[r] = heads[r+1]
	}

	// sort by d-th character in-place
	for r := 0; r < asciiR+1; r++ {
		for heads[r] < tails[r] {
			c := charAt(a[heads[r]], d)
			for c+1 != r {
				exch(a, heads[r], heads[c+1])
				heads[c+1]++
				c = charAt(a[heads[r]], d)
			}
			heads[r]++
		}
	}

	// recursively sort for each character (excludes sentinel -1)
	for r := 0; r < asciiR; r++ {
		inplaceSort(a, tails[r], tails[r+1]-1, d+1)
	}
}
