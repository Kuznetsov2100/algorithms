package str

// RabinKarp struct finds the first occurrence of a pattern string in a text string.
// This implementation uses the Rabin-Karp algorithm.
type RabinKarp struct {
	patHash int64 // pattern hash value
	m       int   // pattern length
	q       int64 // a large prime
	radixR  int   // radix
	RM      int64 // radixR^(m-1) % Q
}

// NewRabinKarp preprocesses the pattern string.
func NewRabinKarp(pat string) *RabinKarp {
	rk := &RabinKarp{radixR: 256, m: len(pat), q: 3768087649, RM: 1}
	// precompute R^(m-1) % q for use in removing leading digit,
	// use mod to avoid overflow
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
	txthash := rk.hash(txt, rk.m) // Use Horner's rule.
	// check for match at offset 0
	if rk.patHash == txthash {
		return 0
	}
	// check for hash match
	// Use rolling hash (and % to avoid overflow).
	for i := rk.m; i < n; i++ {
		// Remove leading digit, add trailing digit, check for match.
		txthash = (txthash + rk.q - rk.RM*int64(txt[i-rk.m])%rk.q) % rk.q
		txthash = (txthash*int64(rk.radixR) + int64(txt[i])) % rk.q
		if rk.patHash == txthash { // match
			return i - rk.m + 1
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
