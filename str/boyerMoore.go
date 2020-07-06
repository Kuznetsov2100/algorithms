package str

// BoyerMoore struct finds the first occurrence of a pattern string in a text string.
// This implementation uses the Boyer-Moore algorithm (with the bad-character rule,
// but not the strong good suffix rule).
type BoyerMoore struct {
	radixR  int    // the radix
	right   []int  // the bad-character skip array
	pattern []byte // store the pattern as a byte array
	pat     string // or as a string
}

// NewBoyerMoore preprocesses the pattern string.
func NewBoyerMoore(pat string) *BoyerMoore {
	bm := &BoyerMoore{radixR: 256, pat: pat, right: make([]int, 256)}
	// position of rightmost occurrence of c in the pattern
	for c := 0; c < bm.radixR; c++ {
		bm.right[c] = -1
	}
	for j := 0; j < len(pat); j++ {
		bm.right[pat[j]] = j
	}
	return bm
}

// NewBoyerMooreR preprocesses the pattern []byte with custom R.
func NewBoyerMooreR(pattern []byte, R int) *BoyerMoore {
	bm := &BoyerMoore{radixR: R, pattern: []byte(pattern), right: make([]int, R)}
	// position of rightmost occurrence of c in the pattern
	for c := 0; c < bm.radixR; c++ {
		bm.right[c] = -1
	}
	for j := 0; j < len(pattern); j++ {
		bm.right[pattern[j]] = j
	}
	return bm
}

// Search returns the index of the first occurrrence of the pattern string in the text string.
func (bm *BoyerMoore) Search(txt string) int {
	m := len(bm.pat)
	n := len(txt)
	skip := 0
	for i := 0; i <= n-m; i += skip {
		skip = 0
		for j := m - 1; j >= 0; j-- {
			if bm.pat[j] != txt[i+j] {
				skip = max(1, j-bm.right[txt[i+j]])
				break
			}
		}
		if skip == 0 {
			return i // found
		}
	}
	return n // not found
}

// SearchByte returns the index of the first occurrrence of the pattern string in the text []byte.
func (bm *BoyerMoore) SearchByte(text []byte) int {
	m := len(bm.pattern)
	n := len(text)
	for i, skip := 0, 0; i <= n-m; i += skip {
		skip = 0
		for j := m - 1; j >= 0; j-- {
			if bm.pattern[j] != text[i+j] {
				skip = max(1, j-bm.right[text[i+j]])
				break
			}
		}
		if skip == 0 {
			return i // found
		}
	}
	return n // not found
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
