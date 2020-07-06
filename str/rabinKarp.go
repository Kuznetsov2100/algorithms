package str

// RabinKarp struct finds the first occurrence of a pattern string in a text string.
// This implementation uses the Rabin-Karp algorithm.
type RabinKarp struct {
	pat     string // the pattern, needed only for Las Vegas
	patHash int64  // pattern hash value
	m       int    // pattern length
	q       int64  // a large prime
	radixR  int    // radix
	RM      int64  // radixR^(m-1) % Q
}

// NewRabinKarp preprocesses the pattern string.
func NewRabinKarp(pat string) *RabinKarp {
	rk := &RabinKarp{pat: pat, radixR: 256, m: len(pat), q: 16777619, RM: 1}
	// precompute R^(m-1) % q for use in removing leading digit
	for i := 1; i <= rk.m-1; i++ {
		rk.RM = (int64(rk.radixR) * rk.RM) % rk.q
	}
	rk.patHash = rk.hash(pat, rk.m)
	return rk
}

// Search returns the index of the first occurrrence of the pattern string in the text string.
func (rk *RabinKarp) Search(txt string) int {
	n := len(txt)
	if n < rk.m {
		return n
	}
	txthash := rk.hash(txt, rk.m)
	// check for match at offset 0
	if rk.patHash == txthash && rk.check(txt, 0) {
		return 0
	}
	// check for hash match; if hash match, check for exact match
	for i := rk.m; i < n; i++ {
		// Remove leading digit, add trailing digit, check for match.
		txthash = (txthash + rk.q - rk.RM*int64(txt[i-rk.m])%rk.q) % rk.q
		txthash = (txthash*int64(rk.radixR) + int64(txt[i])) % rk.q
		// match
		offset := i - rk.m + 1
		if rk.patHash == txthash && rk.check(txt, offset) {
			return offset
		}
	}
	return n // no match
}

// compute hash for key[0..m-1].
func (rk *RabinKarp) hash(key string, m int) (h int64) {
	for j := 0; j < m; j++ {
		h = (int64(rk.radixR)*h + int64(key[j])) % rk.q
	}
	return h
}

func (rk *RabinKarp) check(txt string, i int) bool {
	for j := 0; j < rk.m; j++ {
		if rk.pat[j] != txt[i+j] {
			return false
		}
	}
	return true
}
