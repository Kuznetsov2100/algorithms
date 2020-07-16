package str

import (
	"sort"
)

type SuffixArray struct {
	suffixes []*suffix
}

func NewSuffixArray(text string) *SuffixArray {
	n := len(text)
	sa := &SuffixArray{suffixes: make([]*suffix, n)}
	for i := 0; i < n; i++ {
		sa.suffixes[i] = newsuffix(text, i)
	}

	sort.Slice(sa.suffixes, func(x, y int) bool {
		n := min(sa.suffixes[x].length(), sa.suffixes[y].length())
		for i := 0; i < n; i++ {
			if sa.suffixes[x].charAt(i) < sa.suffixes[y].charAt(i) {
				return true
			}
			if sa.suffixes[x].charAt(i) > sa.suffixes[y].charAt(i) {
				return false
			}
		}
		return sa.suffixes[x].length() < sa.suffixes[y].length()
	})
	return sa
}

func (sa *SuffixArray) Length() int {
	return len(sa.suffixes)
}

func (sa *SuffixArray) Index(i int) int {
	if i < 0 || i >= len(sa.suffixes) {
		panic("invalid i")
	}
	return sa.suffixes[i].index
}

func (sa *SuffixArray) Lcp(i int) int {
	if i < 1 || i >= len(sa.suffixes) {
		panic("invalid i")
	}
	return sa.lcpSuffix(sa.suffixes[i], sa.suffixes[i-1])
}

func (sa *SuffixArray) lcpSuffix(s, t *suffix) int {
	n := min(s.length(), t.length())
	for i := 0; i < n; i++ {
		if s.charAt(i) != t.charAt(i) {
			return i
		}
	}
	return n
}

func (sa *SuffixArray) Select(i int) string {
	if i < 0 || i >= len(sa.suffixes) {
		panic("invalid i")
	}
	return sa.suffixes[i].string()
}

func (sa *SuffixArray) Rank(query string) int {
	lo, hi := 0, len(sa.suffixes)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if cmp := sa.compare(query, sa.suffixes[mid]); cmp < 0 {
			hi = mid - 1
		} else if cmp > 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

func (sa *SuffixArray) compare(query string, s *suffix) int {
	n := min(len(query), s.length())
	for i := 0; i < n; i++ {
		if query[i] < s.charAt(i) {
			return -1
		}
		if query[i] > s.charAt(i) {
			return 1
		}
	}
	return len(query) - s.length()
}

type suffix struct {
	text  string
	index int
}

func newsuffix(text string, index int) *suffix {
	return &suffix{text: text, index: index}
}

func (s *suffix) length() int {
	return len(s.text) - s.index
}

func (s *suffix) charAt(i int) byte {
	return s.text[s.index+i]
}

func (s *suffix) string() string {
	return s.text[s.index:]
}
