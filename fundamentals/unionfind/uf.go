package unionfind

import "fmt"

// UF struct represents a union–find data type (also known as the disjoint-sets data type).
// It supports the classic union and find operations, along with a count operation that returns the
// total number of sets.
// This implementation uses weighted quick union by rank with path compression by halving.
// The constructor takes O(n) time, where n is the number of elements. The union and find operations
// take O(log n) time in the worst case. The count operation takes O(1) time. Moreover, starting from
// an empty data structure with n sites, any intermixed sequence of m union and find operations takes
// O(m α(n)) time, where α(n) is the inverse of Ackermann's function.
type UF struct {
	parent []int  // parent[i] = parent of i
	rank   []int8 // rank[i] = rank of subtree rooted at i (never more than 31)
	count  int    // number of components
}

// NewUF initializes an empty union-find data structure with n elements 0 through n-1.
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

// Find returns the canonical element of the set containing element p.
func (uf *UF) Find(p int) int {
	uf.validate(p)
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]] // path compression by halving
		p = uf.parent[p]
	}
	return p
}

// Count returns the number of sets.
func (uf *UF) Count() int {
	return uf.count
}

// Union merges the set containing element p with the the set containing element q.
func (uf *UF) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP == rootQ {
		return
	}
	// make root of smaller rank point to root of larger rank
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
