package str

// Lcs returns the longest common string of the two specified strings.
func Lcs(s, t string) string {
	suffix1 := NewSuffixArray(s)
	suffix2 := NewSuffixArray(t)

	lcs := ""
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		p := suffix1.Index(i)
		q := suffix2.Index(j)
		x := lcp(s, p, t, q)
		if len(x) > len(lcs) {
			lcs = x
		}
		if compare(s, p, t, q) < 0 {
			i++
		} else {
			j++
		}
	}
	return lcs
}

func lcp(s string, p int, t string, q int) string {
	n := min(len(s)-p, len(t)-q)
	for i := 0; i < n; i++ {
		if s[p+i] != t[q+i] {
			return s[p:(p + i)]
		}
	}
	return s[p:(p + n)]
}

func compare(s string, p int, t string, q int) int {
	n := min(len(s)-p, len(t)-q)
	for i := 0; i < n; i++ {
		if s[p+i] != t[q+i] {
			return int(s[p+i] - t[q+i])
		}
	}
	if len(s)-p < len(t)-q {
		return -1
	} else if len(s)-p > len(t)-q {
		return 1
	} else {
		return 0
	}
}
