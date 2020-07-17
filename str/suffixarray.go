package str

import (
	"sort"
	"strings"
)

// SuffixArray struct represents a suffix array of a string of length n.
// It supports the selecting the ith smallest suffix, getting the index of the ith smallest suffix,
// computing the length of the longest common prefix between the ith smallest suffix and the i-1st
// smallest suffix, and determining the rank of a query string (which is the number of suffixes
// strictly less than the query string).
type SuffixArray struct {
	suffixes sort.StringSlice
	n        int
}

// NewSuffixArray initializes a suffix array for the given text string.
func NewSuffixArray(text string) *SuffixArray {
	n := len(text)
	sa := &SuffixArray{suffixes: make(sort.StringSlice, n), n: n}
	for i := 0; i < n; i++ {
		sa.suffixes[i] = text[i:]
	}
	sa.suffixes.Sort()
	return sa
}

// Length returns the length of the input string.
func (sa *SuffixArray) Length() int {
	return sa.n
}

// Index returns the index into the original string of the ith smallest suffix.
func (sa *SuffixArray) Index(i int) int {
	if i < 0 || i >= sa.n {
		panic("invalid i")
	}
	return sa.n - len(sa.suffixes[i])
}

// Lcp returns the length of the longest common prefix of
// the ith smallest suffix and the i-1st smallest suffix.
func (sa *SuffixArray) Lcp(i int) int {
	if i < 1 || i >= sa.n {
		panic("invalid i")
	}
	return sa.lcpSuffix(sa.suffixes[i], sa.suffixes[i-1])
}

func (sa *SuffixArray) lcpSuffix(s, t string) int {
	n := min(len(s), len(t))
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			return i
		}
	}
	return n
}

// Select returns the ith smallest suffix as a string.
func (sa *SuffixArray) Select(i int) string {
	if i < 0 || i >= sa.n {
		panic("invalid i")
	}
	return sa.suffixes[i]
}

// Rank returns the number of suffixes strictly less than the query string.
func (sa *SuffixArray) Rank(query string) int {
	lo, hi := 0, sa.n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if cmp := strings.Compare(query, sa.suffixes[mid]); cmp < 0 {
			hi = mid - 1
		} else if cmp > 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}
