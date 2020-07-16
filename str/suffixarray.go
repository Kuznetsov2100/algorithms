package str

import (
	"sort"
	"strings"
)

type SuffixArray struct {
	suffixes sort.StringSlice
	n        int
}

func NewSuffixArray(text string) *SuffixArray {
	n := len(text)
	sa := &SuffixArray{suffixes: make(sort.StringSlice, n), n: n}
	for i := 0; i < n; i++ {
		sa.suffixes[i] = text[i:]
	}
	sa.suffixes.Sort()
	return sa
}

func (sa *SuffixArray) Length() int {
	return sa.n
}

func (sa *SuffixArray) Index(i int) int {
	if i < 0 || i >= sa.n {
		panic("invalid i")
	}
	return sa.n - len(sa.suffixes[i])
}

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

func (sa *SuffixArray) Select(i int) string {
	if i < 0 || i >= sa.n {
		panic("invalid i")
	}
	return sa.suffixes[i]
}

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
