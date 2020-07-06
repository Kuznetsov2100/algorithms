package str

type BoyerMoore struct {
	radixR  int
	right   []int
	pattern []byte
	pat     string
}

func NewBoyerMoore(pat string) *BoyerMoore {
	bm := &BoyerMoore{radixR: 256, pat: pat, right: make([]int, 256)}

	for c := 0; c < bm.radixR; c++ {
		bm.right[c] = -1
	}
	for j := 0; j < len(pat); j++ {
		bm.right[pat[j]] = j
	}
	return bm
}

func NewBoyerMooreR(pattern []byte, R int) *BoyerMoore {
	bm := &BoyerMoore{radixR: R, pattern: make([]byte, len(pattern)), right: make([]int, R)}
	for j := 0; j < len(pattern); j++ {
		bm.pattern[j] = pattern[j]
	}
	for c := 0; c < bm.radixR; c++ {
		bm.right[c] = -1
	}
	for j := 0; j < len(pattern); j++ {
		bm.right[pattern[j]] = j
	}
	return bm
}

func (bm *BoyerMoore) Search(txt string) int {
	m := len(bm.pat)
	n := len(txt)
	for i, skip := 0, 0; i < n-m; i += skip {
		skip = 0
		for j := m - 1; j >= 0; j-- {
			if bm.pat[j] != txt[i+j] {
				skip = max(1, j-bm.right[txt[i+j]])
				break
			}
		}
		if skip == 0 {
			return i
		}
	}
	return n
}

func (bm *BoyerMoore) SearchByte(text []byte) int {
	m := len(bm.pattern)
	n := len(text)
	for i, skip := 0, 0; i < n-m; i += skip {
		skip = 0
		for j := m - 1; j >= 0; j-- {
			if bm.pattern[j] != text[i+j] {
				skip = max(1, j-bm.right[text[i+j]])
				break
			}
		}
		if skip == 0 {
			return i
		}
	}
	return n
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
