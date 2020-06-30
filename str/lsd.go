package str

// LsdSort rearranges the array of w-character strings in ascending order.
func LsdSort(a []string, w int) {
	n := len(a)
	R := 256 // extend ASCII alphabet size
	aux := make([]string, n)

	// sort by key-indexed counting on dth character
	for d := w - 1; d >= 0; d-- {

		// compute frequency counts
		count := make([]int, R+1)
		for i := 0; i < n; i++ {
			count[a[i][d]+1]++
		}
		// compute cumulates
		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		// move data
		for i := 0; i < n; i++ {
			aux[count[a[i][d]]] = a[i]
			count[a[i][d]]++
		}
		// copy back
		for i := 0; i < n; i++ {
			a[i] = aux[i]
		}
	}
}
