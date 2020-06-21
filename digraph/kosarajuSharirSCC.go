package digraph

import "fmt"

type KosarajuSharirSCC struct {
	marked []bool
	id     []int
	count  int
}

func NewKosarajuSharirSCC(G *Digraph) *KosarajuSharirSCC {
	scc := &KosarajuSharirSCC{
		marked: make([]bool, G.V()),
		id:     make([]int, G.V()),
	}

	dfo := NewDepthFirstOrder(G.Reverse())
	for _, v := range dfo.ReversePost() {
		if !scc.marked[v] {
			scc.dfs(G, v)
			scc.count++
		}
	}
	return scc
}

func (scc *KosarajuSharirSCC) dfs(G *Digraph, v int) {
	scc.marked[v] = true
	scc.id[v] = scc.count
	for _, w := range G.Adj(v) {
		if !scc.marked[w] {
			scc.dfs(G, w)
		}
	}
}

func (scc *KosarajuSharirSCC) Count() int {
	return scc.count
}

func (scc *KosarajuSharirSCC) StronglyConnected(v, w int) bool {
	scc.validateVertex(v)
	scc.validateVertex(w)
	return scc.id[v] == scc.id[w]
}

func (scc *KosarajuSharirSCC) Id(v int) int {
	scc.validateVertex(v)
	return scc.id[v]
}

func (scc *KosarajuSharirSCC) validateVertex(v int) {
	V := len(scc.marked)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
