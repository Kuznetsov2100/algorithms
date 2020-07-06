package str

type RabinKarp struct {
	pat     string
	patHash int64
	m       int
	q       int64
	R       int
	RM      int64
}

func NewRabinKarp(pat string) *RabinKarp {
	rk := &RabinKarp{pat: pat, R: 256, m: len(pat), q: 16777619, RM: 1}
	for i := 1; i <= rk.m-1; i++ {
		rk.RM = (int64(rk.R) * rk.RM) % rk.q
	}
	rk.patHash = rk.hash(pat, rk.m)
	return rk
}

func (rk *RabinKarp) hash(key string, m int) (h int64) {
	for j := 0; j < m; j++ {
		h = (int64(rk.R)*h + int64(key[j])) % rk.q
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

func (rk *RabinKarp) Search(txt string) int {
	n := len(txt)
	if n < rk.m {
		return n
	}
	txthash := rk.hash(txt, rk.m)
	if rk.patHash == txthash && rk.check(txt, 0) {
		return 0
	}
	for i := rk.m; i < n; i++ {
		txthash = (txthash + rk.q - rk.RM*int64(txt[i-rk.m])%rk.q) % rk.q
		txthash = (txthash*int64(rk.R) + int64(txt[i])) % rk.q
		offset := i - rk.m + 1
		if rk.patHash == txthash && rk.check(txt, offset) {
			return offset
		}
	}
	return n
}
