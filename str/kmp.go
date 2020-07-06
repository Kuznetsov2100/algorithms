package str

// KMP struct finds the first occurrence of a pattern string in a text string.
// This implementation uses a version of the Knuth-Morris-Pratt substring search algorithm.
// The version takes time proportional to n + m*R in the worst case, where n is the length
// of the text string, m is the length of the pattern, and R is the alphabet size. It uses extra
// space proportional to m*R.
type KMP struct {
	radixR int     // the radix
	m      int     // length of pattern
	dfa    [][]int // the KMP automoton
}

// NewKMP preprocesses the pattern string.
func NewKMP(pat string) *KMP {
	radixR, m := 256, len(pat)
	// build DFA from pattern
	dfa := make([][]int, radixR)
	for i := range dfa {
		dfa[i] = make([]int, m)
	}

	dfa[pat[0]][0] = 1 // base case
	for x, j := 0, 1; j < m; j++ {
		// Mismatch transition: If in state j and next char c != pat[j],
		// then the last j-1 characters of input are par[1..j-1], followed by c.
		// To compute dfa[c][j]: Simulate pat[1..j-1] (state x) on DFA and take transition c.
		for c := 0; c < radixR; c++ {
			dfa[c][j] = dfa[c][x] // Copy mismatch cases.
		}
		// Match transition: If in state j and next char c == pat[j],go to j+1
		dfa[pat[j]][j] = j + 1 // Set match case.
		// we have state x after input pat[1..j-1],we can update state x by got input pat[js]
		x = dfa[pat[j]][x] // Update restart state.
	}
	return &KMP{radixR: radixR, m: m, dfa: dfa}
}

// NewKMPR preprocesses the pattern []byte with custom R.
func NewKMPR(pat []byte, R int) *KMP {
	radixR, m := R, len(pat)
	// build DFA from pattern
	dfa := make([][]int, radixR)
	for i := range dfa {
		dfa[i] = make([]int, m)
	}
	dfa[pat[0]][0] = 1
	for x, j := 0, 1; j < m; j++ {
		for c := 0; c < radixR; c++ {
			dfa[c][j] = dfa[c][x] // Copy mismatch cases.
		}
		dfa[pat[j]][j] = j + 1 // Set match case.
		x = dfa[pat[j]][x]     // Update restart state.
	}
	return &KMP{radixR: radixR, m: m, dfa: dfa}
}

// Search returns the index of the first occurrrence of the pattern string in the text string.
func (kmp *KMP) Search(txt string) int {
	// simulate operation of DFA on text
	n, i, j := len(txt), 0, 0
	for ; i < n && j < kmp.m; i++ {
		j = kmp.dfa[txt[i]][j]
	}
	if j == kmp.m {
		return i - kmp.m // found(hit end of pattern)
	}
	return n // not found(hit end of txt)
}

// SearchByte returns the index of the first occurrrence of the pattern []byte in the text []byte.
func (kmp *KMP) SearchByte(txt []byte) int {
	// simulate operation of DFA on text
	n, i, j := len(txt), 0, 0
	for ; i < n && j < kmp.m; i++ {
		j = kmp.dfa[txt[i]][j]
	}
	if j == kmp.m {
		return i - kmp.m // found
	}
	return n // not found
}
