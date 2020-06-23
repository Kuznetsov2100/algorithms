package unionfind

import "fmt"

type UF struct {
	parent []int
	rank   []int8
	count  int
}

func NewUF(n int) *UF {
	if n < 0 {
		panic("n should be non negative")
	}
	uf := &UF{
		count:  n,
		parent: make([]int, n),
		rank:   make([]int8, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 0
	}
	return uf
}

func (uf *UF) Find(p int) int {
	uf.validate(p)
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}
	return p
}

func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP == rootQ {
		return
	}
	if uf.rank[rootP] < uf.rank[rootQ] {
		uf.parent[rootP] = rootQ
	} else if uf.rank[rootP] > uf.rank[rootQ] {
		uf.parent[rootQ] = rootP
	} else {
		uf.parent[rootQ] = rootP
		uf.rank[rootP]++
	}
	uf.count--
}

func (uf *UF) validate(p int) {
	n := len(uf.parent)
	if p < 0 || p >= n {
		panic(fmt.Sprintln("index ", p, " is not between 0 and ", n-1))
	}
}
